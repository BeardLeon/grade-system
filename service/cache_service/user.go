package cache_service

import (
	"github.com/EDDYCJY/go-gin-example/pkg/e"
	"strconv"
	"strings"
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

func (u *User) GetUsersKey() string {
	keys := []string{
		e.CACHE_USER,
		"LIST",
	}

	if u.Name != "" {
		keys = append(keys, u.Name)
	}

	if u.NickName != "" {
		keys = append(keys, u.NickName)
	}

	if u.Password != "" {
		keys = append(keys, u.Password)
	}

	if u.Age >= 0 {
		keys = append(keys, strconv.Itoa(u.Age))
	}

	if u.Sex >= 0 {
		keys = append(keys, strconv.Itoa(u.Sex))
	}

	if u.Major != "" {
		keys = append(keys, u.Major)
	}

	if u.Phone != "" {
		keys = append(keys, u.Phone)
	}

	if u.Status >= 0 {
		keys = append(keys, strconv.Itoa(u.Status))
	}
	if u.PageNum > 0 {
		keys = append(keys, strconv.Itoa(u.PageNum))
	}
	if u.PageSize > 0 {
		keys = append(keys, strconv.Itoa(u.PageSize))
	}

	return strings.Join(keys, "_")
}
