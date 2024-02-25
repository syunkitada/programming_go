package repository

import (
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Config struct {
	mysql.Config
	IsDebug bool
}

func GetDefaultConfig() Config {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		print("Failed to time.LoadLocation: err=%s", err.Error())
		os.Exit(1)
	}

	return Config{
		Config: mysql.Config{
			User:      "admin",
			Passwd:    "adminpass",
			Addr:      "localhost:3306",
			DBName:    "gorm_test",
			Net:       "tcp",
			ParseTime: true,
			Collation: "utf8mb4_unicode_ci",
			Loc:       jst,
		},
		IsDebug: true,
	}
}
