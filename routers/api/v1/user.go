package v1

import (
	"github.com/EDDYCJY/go-gin-example/pkg/app"
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/pkg/util"
	"github.com/EDDYCJY/go-gin-example/service/user_service"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//ID           int       `gorm:"primary_key", json:"id"`
//Name         string    `json:"name"`
//NickName     string    `json:"nick_name"`
//Password     string    `json:"password"`
//Age          int       `json:"age"`
//Sex          int       `json:"sex"`
//Major        string    `json:"major"`
//Phone        string    `json:"phone"`
//Status       int       `json:"status"`

// @Summary Get multiple article tags
// @Produce  json
// @Param id query int false "Id"
// @Param name query string false "Name"
// @Param nick_name query string false "NickName"
// @Param password query string false "Password"
// @Param age query int false "Age"
// @Param sex query int false "Sex"
// @Param major query string false "Major"
// @Param phone query string false "Phone"
// @Param status query string false "Status"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/users [get]
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	nick_name := c.Query("nick_name")
	password := c.Query("password")
	age := com.StrTo(c.Query("age")).MustInt()
	sex := com.StrTo(c.Query("sex")).MustInt()
	major := c.Query("major")
	phone := c.Query("phone")
	status := -1
	if arg := c.Query("status"); arg != "" {
		status = com.StrTo(arg).MustInt()
	}

	userService := user_service.User{
		Name:     name,
		NickName: nick_name,
		Password: password,
		Age:      age,
		Sex:      sex,
		Major:    major,
		Phone:    phone,
		Status:   status,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	users, err := userService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_USERS_FAIL, nil)
		return
	}

	count, err := userService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_USER_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": users,
		"total": count,
	})
}
