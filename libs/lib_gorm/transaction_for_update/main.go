package main

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"lib_gorm/utils/db_client"
)

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
	defer func() {
		// connectionを明示的に閉じる（プロセス終了時に勝手に閉じてくれるがお作法として閉じておく）
		dbClient.MustClose()
	}()

	if err := dbClient.TransactTest1(); err != nil {
		return
	}
}

func (self *DbClient) TransactTest1() (err error) {
	if err = self.DB.AutoMigrate(&Vm{}); err != nil {
		return
	}
	if err = self.DB.AutoMigrate(&Ip{}); err != nil {
		return
	}

	if tmpErr := self.CreateVm("hoge"); tmpErr != nil {
		fmt.Println("Failed CreateVm: ", tmpErr.Error())
	} else {
		fmt.Println("Success CreateVm")
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		nodeId := i
		go func() {
			if node, tmpErr := self.AssignVm("hoge", nodeId); tmpErr != nil {
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

func (self *DbClient) CreateVm(name string) (err error) {
	availableIps := []string{
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.3",
		"192.168.1.4",
		"192.168.1.5",
	}
	err = self.Transact(func(tx *gorm.DB) (err error) {
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

func (self *DbClient) AssignVm(name string, node int) (resultNode int, err error) {
	err = self.Transact(func(tx *gorm.DB) (err error) {
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
