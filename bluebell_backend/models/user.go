package models

import (
	"encoding/json"
	"errors"
	"time"
)

type User struct {
	UserID   string `json:"user_id" db:"user_id"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Role     int    `json:"role" db:"role"`
}

type UserInfo struct {
	*User
	CreateTime time.Time `json:"createTime" db:"create_time"`
	DelFlg     bool      `json:"delFlg" db:"del_flg"`
}

func (u *User) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName string `json:"username" db:"username"`
		Password string `json:"password" db:"password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else {
		u.UserName = required.UserName
		u.Password = required.Password
	}
	return
}

type RegisterForm struct {
	UserName        string `json:"userName"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"re_password"`
}

func (r *RegisterForm) UnmarshalJSON(data []byte) (err error) {
	required := struct {
		UserName        string `json:"userName"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"re_password"`
	}{}
	err = json.Unmarshal(data, &required)
	if err != nil {
		return
	} else if len(required.UserName) == 0 {
		err = errors.New("缺少必填字段username")
	} else if len(required.Password) == 0 {
		err = errors.New("缺少必填字段password")
	} else {
		r.UserName = required.UserName
		r.Password = required.Password
		r.ConfirmPassword = required.ConfirmPassword
	}
	return
}
