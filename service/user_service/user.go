package user_service

import (
	"encoding/json"
	"fmt"
	"github.com/EDDYCJY/go-gin-example/models"
	"github.com/EDDYCJY/go-gin-example/pkg/gredis"
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/EDDYCJY/go-gin-example/service/cache_service"
)

type User struct {
	ID       int
	Name     string
	NickName string
	Password string
	Age      int
	Sex      int
	Major    string
	Phone    string
	Status   int

	PageNum  int
	PageSize int
}

func (u *User) GetUserById() (*models.User, error) {
	return models.GetUser(u.ID)
}

func (u *User) Count() (int, error) {
	return models.GetUserTotal(u.getMaps())
}

func (u *User) GetAll() ([]models.User, error) {

	var (
		users, cacheUsers []models.User
	)

	cache := cache_service.User{
		Name:     u.Name,
		NickName: u.NickName,
		Password: u.Password,
		Age:      u.Age,
		Sex:      u.Sex,
		Major:    u.Major,
		Phone:    u.Phone,
		Status:   u.Status,

		PageNum:  u.PageNum,
		PageSize: u.PageSize,
	}

	fmt.Println("[cache]", cache)

	key := cache.GetUsersKey()

	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheUsers)
			return cacheUsers, nil
		}
	}

	//// redis not exist key
	users, err := models.GetUsers(u.PageNum, u.PageSize, u.getMaps())

	if err != nil {
		return nil, err
	}

	gredis.Set(key, users, 3600)

	return users, nil
}

func (u *User) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	maps["deleted_on"] = 0

	if u.Status >= 0 {
		maps["status"] = u.Status
	}

	return maps
}
