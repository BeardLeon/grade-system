package models

import (
	"fmt"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	ID           int       `gorm:"primary_key", json:"id"`
	Name         string    `json:"name"`
	NickName     string    `json:"nick_name"`
	Password     string    `json:"password"`
	Age          int       `json:"age"`
	Sex          int       `json:"sex"`
	Major        string    `json:"major"`
	Phone        string    `json:"phone"`
	Status       int       `json:"status"`
	CreatedTime  time.Time `json:"created_time"`
	ModifiedTime time.Time `json:"modified_time"`
	DeletedOn    int       `json:"deleted_on"`
}

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}

func ExistUserByID(id int) (bool, error) {
	var user User
	// 未被软删除
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetUserTotal(maps interface{}) (int, error) {
	var count int
	err := db.Model(&User{}).Where((maps)).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return count, nil
}

func GetUsers(pageNum int, pageSize int, maps interface{}) ([]User, error) {
	var (
		users []User
		err   error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&users).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&users).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

func GetUser(id int) (*User, error) {
	var user User
	err := db.Where("id = ? AND deleted_on = ?", id, 0).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

func EditUser(id int, data interface{}) error {
	err := db.Model(&User{}).Where("id = ?", id).Updates(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	return nil
}

func AddUser(data map[string]interface{}) error {
	//ID           int       `gorm:"primary_key" json:"id"`
	//Name         string    `json:"name"`
	//NickName     string    `json:"nick_name"`
	//Password     string    `json:"password"`
	//Age          int       `json:"age"`
	//Sex          int       `json:"sex"`
	//Major        string    `json:"major"`
	//Phone        string    `json:"phone"`
	//Status       int       `json:"status"`
	//CreatedTime  time.Time `json:"created_by"`
	//ModifiedTime time.Time `json:"modified_by"`
	//DeletedOn    int       `json:"deleted_on"`
	err := db.Create(&User{
		ID:       data["id"].(int),
		Name:     data["name"].(string),
		NickName: data["nick_name"].(string),
		Password: data["password"].(string),
		Age:      data["age"].(int),
		Sex:      data["sex"].(int),
		Major:    data["major"].(string),
		Phone:    data["phone"].(string),
		Status:   data["status"].(int),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id int) error {
	err := db.Where("id = ?", id).Delete(User{}).Error
	if err != nil {
		return err
	}
	return nil
}

func Login(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where("name = ? AND password = ?", username, password).Find(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}
