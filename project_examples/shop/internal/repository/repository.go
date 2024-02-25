package repository

//go:generate mockgen -source=$GOFILE -destination=mock_$GOPACKAGE/$GOFILE -package=mock_$GOPACKAGE

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/syunkitada/programming_go/project_examples/shop/internal/model"
)

type IRepositoryCommon interface {
	MustOpenMock() sqlmock.Sqlmock
	MustOpen()
	MustClose()
	MustCreateDatabase()
	MustRecreateDatabase()
	MustMigrate()
}

type IRepositoryItem interface {
	FindItems(input *FindItemsInput) ([]model.Item, error)
	AddItem(item *model.Item) (*model.Item, error)
	DeleteItem(id uint64) error
}

type IRepository interface {
	IRepositoryCommon
	IRepositoryItem
}

type Repository struct {
	conf *Config
	DB   *gorm.DB
}

func New(conf *Config) IRepository {
	return &Repository{
		conf: conf,
	}
}

func (self *Repository) MustOpen() {
	if err := self.Open(); err != nil {
		print("Failed to MustOpen")
		os.Exit(1)
	}
}

func (self *Repository) MustOpenMock() (mock sqlmock.Sqlmock) {
	mock, err := self.OpenMock()
	if err != nil {
		print("Failed to MustOpenMock")
		os.Exit(1)
	}
	return mock
}

func (self *Repository) OpenMock() (mock sqlmock.Sqlmock, err error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, err
	}

	self.DB, err = gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return mock, nil
}

func (self *Repository) Open() (err error) {
	self.DB, err = gorm.Open(mysql.Open(self.conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	if self.conf.IsDebug {
		self.DB.Logger.LogMode(logger.Info)
		self.DB = self.DB.Debug()
	}
	return
}

func (self *Repository) MustClose() {
	if err := self.Open(); err != nil {
		log.Fatalf("Failed Close")
	}
}

func (self *Repository) Close() (err error) {
	if db, err := self.DB.DB(); err != nil {
		return err
	} else {
		if err := db.Close(); err != nil {
			return err
		}
	}
	return
}

func (self *Repository) MustCreateDatabase() {
	if err := self.CreateDatabase(); err != nil {
		log.Fatalf("failed to CreateDatabase")
	}
}

func (self *Repository) MustDropDatabase() {
	if err := self.DropDatabase(); err != nil {
		log.Fatalf("Failed DropDatabase")
	}
}

func (self *Repository) DropDatabase() (err error) {
	dbName := self.conf.DBName
	self.conf.DBName = ""
	defer func() {
		self.conf.DBName = dbName
	}()

	db, err := gorm.Open(mysql.Open(self.conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)).Error; err != nil {
		return err
	}
	return nil
}

func (self *Repository) MustRecreateDatabase() {
	self.MustDropDatabase()
	self.MustCreateDatabase()
}

func (self *Repository) CreateDatabase() (err error) {
	dbName := self.conf.DBName
	self.conf.DBName = ""
	defer func() {
		self.conf.DBName = dbName
	}()

	db, err := gorm.Open(mysql.Open(self.conf.FormatDSN()), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)).Error; err != nil {
		return err
	}
	return nil
}

func (self *Repository) Transact(txFunc func(tx *gorm.DB) (err error)) (err error) {
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

func (self *Repository) TransactWithRetry(txFunc func(tx *gorm.DB) (err error)) (err error) {
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
