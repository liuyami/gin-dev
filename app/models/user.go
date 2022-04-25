package models

import "strconv"

type User struct {
	ID
	Name     string `json:"name" gorm:"not null;comment:用户名称"`
	Mobile   string `json:"mobile" gorm:"not null;index;comment:手机号"`
	Openid   string `json:"openid" gorm:"index;comment:OpenID"`
	Password string `json:"password" gorm:"not null;default:'';comment:密码"`
	TimeStamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
