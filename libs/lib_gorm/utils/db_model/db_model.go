package db_model

import (
	"database/sql/driver"
	"fmt"
	"net"
)

type Vm struct {
	VmId    uint   `gorm:"not null;primary_key;"`
	Name    string `gorm:"not null;unique_index:idx_name_deleted;"`
	Deleted uint   `gorm:"not null;unique_index:idx_name_deleted;"`
}

type VmPort struct {
	VmId   uint `gorm:"primary_key;"`
	Vm     Vm   `gorm:"references:VmId"`
	PortId uint `gorm:"primary_key;"`
	Port   Port `gorm:"references:PortId"`
}

type Port struct {
	PortId uint `gorm:"not null;primary_key;auto_increment;"`
	IP     IP   `gorm:"not null;uniqueIndex;"`
}

type IP struct {
	net.IP
}

func (self *IP) GormDataType() string {
	return "VARBINARY(16)"
}

func (self *IP) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed Scan: value=%v", value)
	}
	*self = IP{IP: bytes}
	return nil
}

func (self IP) Value() (driver.Value, error) {
	ipv4 := self.IP.To4()
	if ipv4 == nil {
		if len(self.IP) != 16 {
			return nil, fmt.Errorf("Invalid: len=%d", len(self.IP))
		}
		var bytes []byte = self.IP
		return bytes, nil
	}
	var bytes []byte = ipv4
	return bytes, nil
}
