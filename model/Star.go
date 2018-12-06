package model

import "time"

type Star struct {
	Id          int64     `xorm:"not null pk autoincr comment('主键') INT(11)" form:"id"`
	NameZh      string    `xorm:"not null comment('球星中文名') VARCHAR(50)" form:"name_zh"`
	NameEn      string    `xorm:"not null comment('球星英文名') VARCHAR(50)" form:"name_en"`
	CreatedTime time.Time `xorm:"created 'created_time'"`
	UpdatedTime time.Time `xorm:"created 'updated_time'"`
}

func (s Star) TableName() string {
	return "t_star"
}
