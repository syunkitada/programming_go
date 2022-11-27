package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"

	"lib_gorm/utils/db_client"
	"lib_gorm/utils/db_model"
)

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
	if err = self.DB.AutoMigrate(&db_model.Vm{}); err != nil {
		return
	}
	if err = self.DB.AutoMigrate(&db_model.Port{}); err != nil {
		return
	}
	if err = self.DB.AutoMigrate(&db_model.VmPort{}); err != nil {
		return
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		name := fmt.Sprintf("hoge%d", i)
		go func() {
			if tmpErr := self.CreateVm(name); tmpErr != nil {
				fmt.Println("Failed CreateVm: ", tmpErr.Error())
			} else {
				fmt.Println("Success CreateVm")
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
	err = self.TransactWithRetry(func(tx *gorm.DB) (err error) {
		var vms []db_model.Vm
		if err = tx.Table("vms").Select("name").Where("name = ? AND deleted = 0", name).Scan(&vms).Error; err != nil {
			return
		}
		if len(vms) > 0 {
			fmt.Println("vm is already exists")
			return
		}
		vm := db_model.Vm{Name: name}
		if err = tx.Create(&vm).Error; err != nil {
			return
		}

		var assignedPorts []db_model.Port
		if err = tx.Table("ports").Select("*").Scan(&assignedPorts).Error; err != nil {
			return
		}
		var freeIps []string
		for _, availableIp := range availableIps {
			isAssigned := false
			for _, port := range assignedPorts {
				if port.IP.String() == availableIp {
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

		port := db_model.Port{
			IP: db_model.IP{IP: net.ParseIP(freeIps[0])},
		}
		if err = tx.Create(&port).Error; err != nil {
			if strings.Contains(err.Error(), "Duplicate entry") {
				err = &db_client.RetryError{Msg: err.Error(), Ttl: 3}
			}
			return
		}

		vmPort := db_model.VmPort{
			VmId:   vm.VmId,
			PortId: port.PortId,
		}
		if err = tx.Create(&vmPort).Error; err != nil {
			return
		}

		return
	})
	return
}
