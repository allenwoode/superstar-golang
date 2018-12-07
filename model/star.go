package model

import "time"

type Star struct {
	Id          int       `xorm:"not null pk autoincr comment('主键') INT(11)" form:"id"`
	Avatar      string    `xorm:"not null comment('球星头像') VARCHAR(255)" form:"avatar"`
	NameZh      string    `xorm:"not null comment('球星中文名') VARCHAR(50)" form:"name_zh"`
	NameEn      string    `xorm:"not null comment('球星英文名') VARCHAR(50)" form:"name_en"`
	CreatedTime time.Time `xorm:"created 'created_time'"`
	UpdatedTime time.Time `xorm:"created 'updated_time'"`
}

func (s Star) TableName() string {
	return "t_star"
}
