package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"

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

type Vm struct {
	Id         uint   `gorm:"not null;primary_key;"`
	Name       string `gorm:"not null;unique_index:idx_name_deleted;"`
	Deleted    uint   `gorm:"not null;unique_index:idx_name_deleted;"`
	Address    string `gorm:"not null;"`
	AssignNode int    `gorm:"not null;"`
}

type Ip struct {
	Id      uint   `gorm:"not null;primary_key;"`
	Address string `gorm:"not null;unique_index;"`
	VmId    uint   `gorm:"not null;"`
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
	if err = db.AutoMigrate(&Vm{}).Error; err != nil {
		return
	}
	if err = db.AutoMigrate(&Ip{}).Error; err != nil {
		return
	}

	if tmpErr := CreateVm(db, "hoge"); tmpErr != nil {
		fmt.Println("Failed CreateUser: ", tmpErr.Error())
	} else {
		fmt.Println("Success CreateUser")
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		nodeId := i
		go func() {
			if node, tmpErr := AssignVm(db, "hoge", nodeId); tmpErr != nil {
				fmt.Printf("Failed AssignVm: err=%s node=%d\n", tmpErr.Error(), node)
			} else {
				fmt.Printf("Success AssignVm: node=%d\n", node)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	log.Printf("End: %v", err)
	return
}

func CreateVm(db *gorm.DB, name string) (err error) {
	availableIps := []string{
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.3",
		"192.168.1.4",
		"192.168.1.5",
	}
	err = Transact(db, func(tx *gorm.DB) (err error) {
		var vms []Vm
		if err = tx.Table("vms").Select("name").Where("name = ? AND deleted = 0", name).Scan(&vms).Error; err != nil {
			return
		}
		if len(vms) > 0 {
			fmt.Println("vm is already exists")
			return
		}
		vm := Vm{Name: name}
		if err = tx.Create(&vm).Error; err != nil {
			return
		}

		var assignedIps []Ip
		if err = tx.Table("ips").Select("id, address").Scan(&assignedIps).Error; err != nil {
			return
		}
		var freeIps []string
		for _, availableIp := range availableIps {
			isAssigned := false
			for _, assignedIp := range assignedIps {
				if availableIp == assignedIp.Address {
					isAssigned = true
					break
				}
			}
			if !isAssigned {
				freeIps = append(freeIps, availableIp)
			}
		}

		if len(freeIps) == 0 {
			err = fmt.Errorf("Failed assign: freeIps is not found")
			return
		}

		ip := Ip{Address: freeIps[0], VmId: vm.Id}
		if err = tx.Create(&ip).Error; err != nil {
			return
		}

		if err = tx.Table("vms").Where("id = ?", vm.Id).Updates(map[string]interface{}{
			"address": ip.Address,
		}).Error; err != nil {
			return
		}

		return
	})
	return
}

func AssignVm(db *gorm.DB, name string, node int) (resultNode int, err error) {
	err = Transact(db, func(tx *gorm.DB) (err error) {
		var vms []Vm
		// FOR UPDATEはindexエントリに対してロックがかかる
		// 同じ行を参照しようとしたトランザクションはここで待たされる
		if err = tx.Set("gorm:query_option", "FOR UPDATE").Table("vms").
			Select("id, name, assign_node").Where("name = ? AND deleted = 0", name).Scan(&vms).Error; err != nil {
			return
		}

		if len(vms) == 0 {
			err = fmt.Errorf("vm is not found")
			return
		}
		vm := vms[0]
		if vm.AssignNode != 0 {
			resultNode = vm.AssignNode
			err = fmt.Errorf("vm is already assigned")
			return
		}

		if err = tx.Table("vms").Where("id = ?", vm.Id).Updates(map[string]interface{}{
			"assign_node": node,
		}).Error; err != nil {
			return
		}

		resultNode = node

		return
	})
	return
}
