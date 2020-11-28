package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Transact(db *gorm.DB, txFunc func(tx *gorm.DB) (err error)) (err error) {
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
				log.Printf("Failed rollback on err: %s", tmpErr.Error)
			}
		} else {
			if err = tx.Commit().Error; err != nil {
				log.Printf("Failed commit: %s", err.Error())
				if tmpErr := tx.Rollback().Error; tmpErr != nil {
					log.Printf("Failed rollback on commit: %s", tmpErr)
				}
			}
		}
	}()
	err = txFunc(tx)
	return
}

type User struct {
	gorm.Model
	Name    string `gorm:"not null;unique_index:idx_name_deleted;"`
	Deleted uint   `gorm:"not null;unique_index:idx_name_deleted;"`
}

type Project struct {
	gorm.Model
	Name  string `gorm:"not null;unique_index:idx_name_deleted;"`
	Exist bool   `gorm:"default:true;unique_index:idx_name_deleted;"`
}

func main() {
	connection := "goapp:goapppass@tcp(127.0.0.1:3306)/gorm_test?parseTime=true"
	cmds := []string{"mysql", "-ugoapp", "-pgoapppass", "-e", "drop database if exists gorm_test; create database gorm_test;"}
	out, err := exec.Command(cmds[0], cmds[1:]...).CombinedOutput()
	if err != nil {
		log.Fatalf("Failed connect: out=%s, err=%v", string(out), err)
	}

	db, err := gorm.Open("mysql", connection)
	if err != nil {
		log.Fatalf("Failed connect: %v", err)
	}
	defer db.Close()
	db.LogMode(true)

	if err := TransactTest1(db); err != nil {
		return
	}
}

func TransactTest1(db *gorm.DB) (err error) {
	if err = db.AutoMigrate(&User{}).Error; err != nil {
		return
	}
	if err = db.AutoMigrate(&Project{}).Error; err != nil {
		return
	}

	if tmpErr := CreateUser(db, "hoge"); tmpErr != nil {
		fmt.Println("Failed CreateUser: ", tmpErr.Error())
	} else {
		fmt.Println("Success CreateUser")
	}
	if tmpErr := DeleteUser(db, "hoge"); tmpErr != nil {
		fmt.Println("Failed DeleteUser: ", tmpErr.Error())
	} else {
		fmt.Println("Success DeleteUser")
	}
	if tmpErr := CreateUser(db, "hoge"); tmpErr != nil {
		fmt.Println("Failed CreateUser: ", tmpErr.Error())
	} else {
		fmt.Println("Success CreateUser")
	}

	if tmpErr := CreateProject(db, "hoge"); tmpErr != nil {
		fmt.Println("Failed CreateProject: ", tmpErr.Error())
	} else {
		fmt.Println("Success CreateProject")
	}
	if tmpErr := DeleteProject(db, "hoge"); tmpErr != nil {
		fmt.Println("Failed DeleteProject: ", tmpErr.Error())
	} else {
		fmt.Println("Success DeleteProject")
	}
	if tmpErr := CreateProject(db, "hoge"); tmpErr != nil {
		fmt.Println("Failed CreateProject: ", tmpErr.Error())
	} else {
		fmt.Println("Success CreateProject")
	}

	log.Printf("End: %v", err)
	return
}

func CreateUser(db *gorm.DB, name string) (err error) {
	err = Transact(db, func(tx *gorm.DB) (err error) {
		var users []User
		if err = tx.Table("users").Select("name").Where("name = ? AND deleted = 0", name).Scan(&users).Error; err != nil {
			return
		}
		if len(users) > 0 {
			fmt.Println("user is already exists")
			return
		}
		if err = tx.Create(&User{Name: name}).Error; err != nil {
			return
		}
		return
	})
	return
}

func DeleteUser(db *gorm.DB, name string) (err error) {
	err = Transact(db, func(tx *gorm.DB) (err error) {
		var users []User
		if err = tx.Table("users").Select("id").Where("name = ? AND deleted = 0", name).Scan(&users).Error; err != nil {
			return
		}
		if len(users) == 0 {
			fmt.Println("user is already deleted")
			return
		}
		if err = tx.Model(User{}).Where("name = ?", name).Updates(map[string]interface{}{
			"deleted":    users[0].ID,
			"deleted_at": time.Now(),
		}).Error; err != nil {
			return
		}
		return
	})
	return
}

func CreateProject(db *gorm.DB, name string) (err error) {
	err = Transact(db, func(tx *gorm.DB) (err error) {
		var users []Project
		if err = tx.Table("projects").Select("name").Where("name = ? AND exist IS NOT NULL", name).Scan(&users).Error; err != nil {
			return
		}
		if len(users) > 0 {
			fmt.Println("user is already exists")
			return
		}
		if err = tx.Create(&Project{Name: name}).Error; err != nil {
			return
		}
		return
	})
	return
}

func DeleteProject(db *gorm.DB, name string) (err error) {
	err = Transact(db, func(tx *gorm.DB) (err error) {
		var users []Project
		if err = tx.Table("projects").Select("id").Where("name = ? AND exist IS NOT NULL", name).Scan(&users).Error; err != nil {
			return
		}
		if len(users) == 0 {
			fmt.Println("user is already deleted")
			return
		}
		if err = tx.Model(Project{}).Where("name = ?", name).Updates(map[string]interface{}{
			"exist":      nil,
			"deleted_at": time.Now(),
		}).Error; err != nil {
			return
		}
		return
	})
	return
}
