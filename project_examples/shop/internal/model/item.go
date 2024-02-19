package model

type Item struct {
	ID      uint64 `gorm:"colomn:id;type:bigint;not null;primaryKey;autoIncrement;"`
	Name    string `gorm:"colomn:name;type:varchar(250);not null;index:idx_name_deleted,unique;"`
	Deleted uint64 `gorm:"colomn:deleted;type:bigint;not null;index:idx_name_deleted,unique;"`
}

type User struct {
	ID      uint64 `gorm:"colomn:id;type:bigint;not null;primaryKey;autoIncrement;"`
	Name    string `gorm:"colomn:name;type:varchar(250);not null;index:idx_name_deleted,unique;"`
	Deleted uint64 `gorm:"colomn:deleted;type:bigint;not null;index:idx_name_deleted,unique;"`
}
