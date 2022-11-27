package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"

	"lib_gorm/utils/db_client"
)

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

type DbClient struct {
	*db_client.DbClient
}

func main() {
	dbClient := DbClient{
		DbClient: db_client.New(&db_client.DefaultConfig),
	}

	dbClient.MustDropDatabase()
	dbClient.MustCreateDatabase()
	dbClient.MustOpen()

	if err := dbClient.TransactTest1(); err != nil {
		return
	}
}

func (self *DbClient) TransactTest1() (err error) {
	if err = self.DB.AutoMigrate(&User{}); err != nil {
		return
	}
	if err = self.DB.AutoMigrate(&Project{}); err != nil {
		return
	}

	if tmpErr := self.CreateUser("hoge"); tmpErr != nil {
		fmt.Println("Failed CreateUser: ", tmpErr.Error())
	} else {
		fmt.Println("Success CreateUser")
	}
	if tmpErr := self.DeleteUser("hoge"); tmpErr != nil {
		fmt.Println("Failed DeleteUser: ", tmpErr.Error())
	} else {
		fmt.Println("Success DeleteUser")
	}
	if tmpErr := self.CreateUser("hoge"); tmpErr != nil {
		fmt.Println("Failed CreateUser: ", tmpErr.Error())
	} else {
		fmt.Println("Success CreateUser")
	}

	if tmpErr := self.CreateProject("hoge"); tmpErr != nil {
		fmt.Println("Failed CreateProject: ", tmpErr.Error())
	} else {
		fmt.Println("Success CreateProject")
	}
	if tmpErr := self.DeleteProject("hoge"); tmpErr != nil {
		fmt.Println("Failed DeleteProject: ", tmpErr.Error())
	} else {
		fmt.Println("Success DeleteProject")
	}
	if tmpErr := self.CreateProject("hoge"); tmpErr != nil {
		fmt.Println("Failed CreateProject: ", tmpErr.Error())
	} else {
		fmt.Println("Success CreateProject")
	}

	log.Printf("End: %v", err)
	return
}

func (self *DbClient) CreateUser(name string) (err error) {
	err = self.Transact(func(tx *gorm.DB) (err error) {
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

func (self *DbClient) DeleteUser(name string) (err error) {
	err = self.Transact(func(tx *gorm.DB) (err error) {
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

func (self *DbClient) CreateProject(name string) (err error) {
	err = self.Transact(func(tx *gorm.DB) (err error) {
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

func (self *DbClient) DeleteProject(name string) (err error) {
	err = self.Transact(func(tx *gorm.DB) (err error) {
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
