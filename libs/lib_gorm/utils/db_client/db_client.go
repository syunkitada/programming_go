package db_client

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-sql-driver/mysql"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	mysql.Config
	IsDebug bool
}

var DefaultConfig = Config{
	Config: mysql.Config{
		User:      "admin",
		Passwd:    "adminpass",
		Addr:      "localhost:3306",
		DBName:    "gorm_test",
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
	},
	IsDebug: true,
}

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatalf("Failed LoadLocation: err=%s", err.Error())
	}
	DefaultConfig.Config.Loc = jst
}

type DbClient struct {
	conf *Config
	DB   *gorm.DB
}

func New(conf *Config) *DbClient {
	return &DbClient{
		conf: conf,
	}
}

func (self *DbClient) MustOpen() {
	if err := self.Open(); err != nil {
		log.Fatalf("Failed Open")
	}
}

func (self *DbClient) Open() (err error) {
	self.DB, err = gorm.Open(gorm_mysql.Open(self.conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	if self.conf.IsDebug {
		self.DB.Logger.LogMode(logger.Info)
		self.DB = self.DB.Debug()
	}
	return
}

func (self *DbClient) MustClose() {
	if err := self.Open(); err != nil {
		log.Fatalf("Failed Close")
	}
}

func (self *DbClient) Close() (err error) {
	if db, err := self.DB.DB(); err != nil {
		return err
	} else {
		if err := db.Close(); err != nil {
			return err
		}
	}
	return
}

func (self *DbClient) MustDropDatabase() {
	if err := self.DropDatabase(); err != nil {
		log.Fatalf("Failed DropDatabase")
	}
}

func (self *DbClient) DropDatabase() (err error) {
	dbName := self.conf.DBName
	self.conf.DBName = ""
	defer func() {
		self.conf.DBName = dbName
	}()

	db, err := gorm.Open(gorm_mysql.Open(self.conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)).Error; err != nil {
		return err
	}
	return nil
}

func (self *DbClient) MustCreateDatabase() {
	if err := self.CreateDatabase(); err != nil {
		log.Fatalf("Failed CreateDatabase")
	}
}

func (self *DbClient) CreateDatabase() (err error) {
	dbName := self.conf.DBName
	self.conf.DBName = ""
	defer func() {
		self.conf.DBName = dbName
	}()

	db, err := gorm.Open(gorm_mysql.Open(self.conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)).Error; err != nil {
		return err
	}
	return nil
}

func (self *DbClient) Transact(txFunc func(tx *gorm.DB) (err error)) (err error) {
	tx := self.DB.Begin()
	if err = tx.Error; err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				log.Printf("Failed rollback on recover: %s", tmpErr.Error())
			}
			err = fmt.Errorf("Recovered: %v", p)
		} else if err != nil {
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				log.Printf("Failed rollback on err: %s", tmpErr.Error())
			} else {
				log.Printf("Rollbacked because of err: %s", err.Error())
			}
		} else {
			if err = tx.Commit().Error; err != nil {
				log.Printf("Failed commit: %s", err.Error())
				if tmpErr := tx.Rollback().Error; tmpErr != nil {
					log.Printf("Failed rollback on commit: %s", tmpErr.Error())
				}
			}
		}
	}()
	err = txFunc(tx)
	return
}

type RetryError struct {
	Ttl int
	Msg string
}

func (e *RetryError) Error() string {
	return e.Msg
}

func (self *DbClient) TransactWithRetry(txFunc func(tx *gorm.DB) (err error)) (err error) {
	err = transact(self.DB, txFunc)
	if err != nil {
		switch err.(type) {
		case *RetryError:
			ttl := err.(*RetryError).Ttl
			n := rand.Intn(ttl)
			time.Sleep(time.Duration(n) * time.Second)
			for i := 0; i < ttl; i++ {
				fmt.Printf("Retry count=%d, %s\n", i, err.Error())
				err = transact(self.DB, txFunc)
				switch err.(type) {
				case *RetryError:
					continue
				default:
					return
				}
			}
		default:
			return
		}
	}
	return
}

func transact(db *gorm.DB, txFunc func(tx *gorm.DB) (err error)) (err error) {
	tx := db.Begin()
	if err = tx.Error; err != nil {
		return
	}
	defer func() {
		if p := recover(); p != nil {
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				log.Printf("Failed rollback on recover: %s", tmpErr.Error())
			}
			err = fmt.Errorf("Recovered: %v", p)
		} else if err != nil {
			if tmpErr := tx.Rollback().Error; tmpErr != nil {
				log.Printf("Failed rollback on err: %s", tmpErr.Error())
			} else {
				log.Printf("Rollbacked because of err: %s", err.Error())
			}
		} else {
			if err = tx.Commit().Error; err != nil {
				log.Printf("Failed commit: %s", err.Error())
				if tmpErr := tx.Rollback().Error; tmpErr != nil {
					log.Printf("Failed rollback on commit: %s", tmpErr.Error())
				}
			}
		}
	}()
	err = txFunc(tx)
	return
}
