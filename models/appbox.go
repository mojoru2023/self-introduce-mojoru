package models

import (
	"appbox_go_v/utils"
	"database/sql/driver"
	"fmt"
	"time"
)

// App_info  Model
// type App_info struct {
// 	ID          int    `json:"id"`
// 	App_name    string `json:"app_name"`
// 	App_type    string `json:"app_type"`
// 	Domain      string `json:"domain"`
// 	Stack       string `json:"stack"`
// 	Fontport    string `json:"fontport"`
// 	Backendport string `json:"backendport"`
// 	Delete_flg  string `gorm:"default:'0'"` //gorm 设置默认值
// 	Others      string `gorm:"default: ''"` //gorm 设置默认值
// }

type LocalTime time.Time

func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// gorm框架表名自动加s问题
type OPENWORK_KAKAKU_JOBS struct {
	Id             int64      `json:"id"`
	OpenworkPython int64      `json:"openwork_python"`
	OpenworkGolang int64      `json:"openwork_golang"`
	OpenworkVue    int64      `json:"openwork_vue"`
	KakakuPython   int64      `json:"kakaku_python"`
	KakakuGolang   int64      `json:"kakaku_golang"`
	KakakuVue      int64      `json:"kakaku_vue"`
	CreateTime     *LocalTime `json:"create_time" gorm:"autoCreateTime"`
}

func (OPENWORK_KAKAKU_JOBS) TableName() string {
	return "openwork_kakaku_jobs"
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// Appbox 创建todo
// 添加

func CreateDt(dt *OPENWORK_KAKAKU_JOBS) (err error) {
	err = utils.DB.Create(&dt).Error
	return
}

// 查看所有的待办事项

func GetAllDt() (dtList []*OPENWORK_KAKAKU_JOBS, err error) {
	if err = utils.DB.Debug().Find(&dtList).Error; err != nil {
		return nil, err
	}
	return
}

// 查一条todo

func GetOneDt(id string) (dt *OPENWORK_KAKAKU_JOBS, err error) {
	dt = new(OPENWORK_KAKAKU_JOBS)
	if err = utils.DB.Debug().Where("id=?", id).Find(dt).Error; err != nil {
		return nil, err
	}
	return
}

// 修改某一个待办事项

func UpdateOneDt(dt *OPENWORK_KAKAKU_JOBS) (err error) {
	err = utils.DB.Save(dt).Error
	return
}

// 删除某一个待办事项

func DeleteOneDt(id string) (dt *OPENWORK_KAKAKU_JOBS, err error) {
	dt = new(OPENWORK_KAKAKU_JOBS)
	if err = utils.DB.Debug().Where("id=?", id).Find(dt).Error; err != nil {
		return nil, err
	}

	return
}
