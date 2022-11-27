package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"

	"lib_gorm/utils/db_client"
)

type GormBasic struct {
	gorm.Model
	Name string
}

type BasicType struct {
	Id    int        `gorm:"primary_key;"`
	Str   string     `gorm:"not null;"`
	PStr  *string    `gorm:""`
	Int   int        `gorm:"not null;"`
	PInt  *int       `gorm:""`
	Time  time.Time  `gorm:"not null;"`
	PTime *time.Time `gorm:""`
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

	if err := dbClient.BasicSenario(); err != nil {
		log.Fatalf("Failed BasicSenario: %v", err)
	}

	if err := dbClient.BasicSenarioWithGorm(); err != nil {
		log.Fatalf("Failed BasicSenarioWithGorm: %v", err)
	}
}

func (self *DbClient) BasicSenario() (err error) {
	// CREATE TABLE `basic_types` (`id` int AUTO_INCREMENT,`str` varchar(255) NOT NULL,`p_str` varchar(255),`int` int NOT NULL,`p_int` int,`time` timestamp NOT NULL,`p_time` timestamp NULL , PRIMARY KEY (`id`))
	if err = self.DB.AutoMigrate(&BasicType{}); err != nil {
		return
	}

	// Create
	// INSERT  INTO `basic_types` (`str`,`p_str`,`int`,`p_int`,`time`,`p_time`) VALUES ('Hoge',NULL,0,NULL,'2019-10-21 22:03:18',NULL)
	hoge := BasicType{Str: "Hoge", Int: 0, Time: time.Now()}
	if err = self.DB.Create(&hoge).Error; err != nil {
		return
	}

	pstr := "piyo"
	pint := 0
	ptime := time.Now()
	piyo := BasicType{Str: "Piyo", PStr: &pstr, Int: 2, PInt: &pint, Time: time.Now(), PTime: &ptime}
	// INSERT  INTO `basic_types` (`str`,`p_str`,`int`,`p_int`,`time`,`p_time`) VALUES ('Piyo','piyo',2,0,'2019-10-21 21:43:10','2019-10-21 21:43:10')
	if err = self.DB.Create(&piyo).Error; err != nil {
		return
	}

	// WhereでStructを利用すると初期値以外の値でWHEREが実行される
	// UPDATE `basic_types` SET `str` = 'piyo1'  WHERE `basic_types`.`id` = 2 AND ((`basic_types`.`str` = 'piyo'))
	if err = self.DB.Model(&piyo).Where(&BasicType{Str: "piyo"}).Updates(&BasicType{
		Str: "piyo1",
	}).Error; err != nil {
		return
	}

	// Whereで Structに無駄にデータを入れてると、primary_keyでないものも対象となるので注意
	// UPDATE `basic_types` SET `time` = '2019-10-21 22:26:21'  WHERE `basic_types`.`id` = 2 AND ((`basic_types`.`id` = 2) AND (`basic_types`.`str` = 'piyo2') AND (`basic_types`.`p_str` = 'piyo') AND (`basic_types`.`int` = 2) AND (`basic_types`.`p_int` = 0) AND (`basic_types`.`time` = '2019-10-21 22:26:21') AND (`basic_types`.`p_time` = '2019-10-21 22:26:21'))
	if err = self.DB.Model(&piyo).Where(&piyo).Updates(&BasicType{
		Str: "piyo2",
	}).Error; err != nil {
		return
	}

	// UpdatesでStructを利用すると、初期値以外のデータが更新されないので注意
	// UPDATE `basic_types` SET `time` = '2019-10-21 22:03:18'  WHERE `basic_types`.`id` = 2 AND ((`basic_types`.`id` = 2) AND (`basic_types`.`str` = 'Piyo') AND (`basic_types`.`p_str` = 'piyo') AND (`basic_types`.`int` = 2) AND (`basic_types`.`p_int` = 0) AND (`basic_types`.`time` = '2019-10-21 22:03:18') AND (`basic_types`.`p_time` = '2019-10-21 22:03:18'))
	if err = self.DB.Model(&piyo).Where(&piyo).Updates(&BasicType{
		Str:  "",
		Int:  0,
		PInt: nil,
		Time: time.Now(),
	}).Error; err != nil {
		return
	}

	// Saveを利用すると、更新する必要のないカラムもすべて更新されるので注意
	// UPDATE `basic_types` SET `str` = '', `p_str` = NULL, `int` = 0, `p_int` = NULL, `time` = '2019-10-21 22:03:18', `p_time` = '2019-10-21 22:03:18'  WHERE `basic_types`.`id` = 2
	piyo.Str = ""
	piyo.Int = 0
	piyo.PStr = nil
	piyo.PInt = nil
	if err = self.DB.Save(&piyo).Error; err != nil {
		return
	}

	// 更新処理は以下のように行うとよい
	// UPDATE `basic_types` SET `int` = 0, `p_int` = NULL, `p_str` = NULL, `str` = ''  WHERE `basic_types`.`id` = 2 AND ((id = 2))
	if err = self.DB.Model(&piyo).Where("id = ?", piyo.Id).
		Updates(map[string]interface{}{
			"str":   "",
			"p_str": nil,
			"int":   0,
			"p_int": nil,
		}).Error; err != nil {
		return
	}

	// 更新処理時に更新行数とerrorを取得する
	// 更新時や削除時は、対象が存在しない場合も成功となりバグの原因となるので注意する(必要に応じて明示的に行数も確認すべき)
	// UPDATE `basic_types` SET `str` = 'piyo2'  WHERE `basic_types`.`id` = 2 AND ((id = 2))
	// [1 rows affected or returned ]
	// DEBUG update query: count=1, err=<nil>
	query := self.DB.Model(&piyo).Where("id = ?", piyo.Id).Updates(map[string]interface{}{"str": "piyo2"})
	count, tmpErr := query.RowsAffected, query.Error
	fmt.Printf("DEBUG update query: count=%d, err=%v\n", count, tmpErr)

	var types []BasicType
	// SELECT * FROM `basic_types`
	if err = self.DB.Find(&types).Error; err != nil {
		return
	}

	// Findを単体で利用すると*で取得するため、Selectで必要なものだけに絞ったほうがよい
	// SELECT str, p_str FROM `basic_types`
	if err = self.DB.Select("str, p_str").Find(&types).Error; err != nil {
		return
	}

	var tmpType BasicType
	// Firstは無駄にORDER BYするので利用しないほうがよい
	// SELECT * FROM `basic_types`  WHERE (id = 2) ORDER BY `basic_types`.`id` ASC LIMIT 1
	if err = self.DB.Where("id = ?", piyo.Id).First(&tmpType).Error; err != nil {
		return
	}

	// DELETEでは、primary_keyでWhereしてDeleteされる
	// DELETE FROM `basic_types`  WHERE `basic_types`.`id` = 2
	if err = self.DB.Delete(&tmpType).Error; err != nil {
		return
	}

	return
}

// gorm.Modelを利用した場合の挙動
// deleted_atを利用した論理削除前提の挙動になる
func (self *DbClient) BasicSenarioWithGorm() (err error) {
	// CREATE TABLE `gorm_basics` (`id` int unsigned AUTO_INCREMENT,`created_at` timestamp NULL,`updated_at` timestamp NULL,`deleted_at` timestamp NULL,`name` varchar(255) , PRIMARY KEY (`id`))
	// CREATE INDEX idx_gorm_basics_deleted_at ON `gorm_basics`(deleted_at)
	if err = self.DB.AutoMigrate(&GormBasic{}); err != nil {
		return
	}

	// INSERT  INTO `gorm_basics` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES ('2019-10-21 23:23:04','2019-10-21 23:23:04',NULL,'hoge')
	if err = self.DB.Create(&GormBasic{Name: "hoge"}).Error; err != nil {
		return
	}

	var datum []GormBasic
	// SELECT id FROM `gorm_basics`  WHERE `gorm_basics`.`deleted_at` IS NULL AND ((name = 'hoge'))
	if err = self.DB.Model(&GormBasic{}).Select("id").Where("name = ?", "hoge").Find(&datum).Error; err != nil {
		return
	}

	// UPDATE `gorm_basics` SET `name` = 'piyo', `updated_at` = '2019-10-21 23:20:55'  WHERE `gorm_basics`.`deleted_at` IS NULL AND ((id = 1))
	if err = self.DB.Model(&GormBasic{}).Where("id = ?", datum[0].ID).Updates(map[string]interface{}{
		"name": "piyo",
	}).Error; err != nil {
		return
	}

	// UPDATE `gorm_basics` SET `deleted_at`='2019-10-21 23:22:11'  WHERE `gorm_basics`.`deleted_at` IS NULL AND `gorm_basics`.`id` = 1
	if err = self.DB.Delete(&datum[0]).Error; err != nil {
		return
	}

	// DELETE FROM `gorm_basics`  WHERE `gorm_basics`.`id` = 1
	if err = self.DB.Unscoped().Delete(&datum[0]).Error; err != nil {
		return
	}

	return
}
