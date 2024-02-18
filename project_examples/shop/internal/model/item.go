package model

type User struct {
	ID      uint   `gorm:"colomn:id;not null;primaryKey;autoIncrement;"`
	Name    string `gorm:"colomn:name;not null;index:idx_name_deleted,unique;"`
	Deleted uint   `gorm:"colomn:deleted;not null;index:idx_name_deleted,unique;"`
}

type Item struct {
	ID      uint   `gorm:"colomn:id;not null;primaryKey;autoIncrement;"`
	Name    string `gorm:"colomn:name;not null;index:idx_name_deleted,unique;"`
	Deleted uint   `gorm:"colomn:deleted;not null;index:idx_name_deleted,unique;"`
}
