package models

import (
	"fmt"
	"gitee.com/ax/todo/common"
	"gitee.com/ax/todo/databases"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Content string `gorm:"type:varchar(100);not null" json:"content" binding:"required,min=1,max=100" label:"内容"`
	Status  int    `gorm:"type:int;default:1" json:"status" binding:"oneof=0 1" label:"状态"`
	Remarks string `gorm:"type:varchar(50)" json:"remarks" lable:"备注"`
}

func GetTodoList(params ValidGetTodoList) (err error, data []Todo, total int64) {
	page := params.Page
	size := params.Size
	kw := params.Kw

	tx := databases.DB.Model(&Todo{})
	if kw != "" {
		tx.Where(fmt.Sprintf("content like '%%%s%%'", kw))
	}
	err = tx.Limit(size).Offset((page - 1) * size).Find(&data).Count(&total).Error
	if err != nil {
		return err, nil, 0
	}
	return nil, data, total
}
func CreateTodo(data Todo) (err error) {
	err = databases.DB.Model(&Todo{}).Create(&data).Error
	return err
}
func DeleteTodo(tid int) (err error){
	tx := databases.DB.Delete(&Todo{}, tid)
	err = common.HandleError(tx, "所指定的记录不存在！")
	return
}
func UpdateTodo(tid int, data Todo) (err error)  {
	var maps = make(map[string]interface{})
	maps["content"] = data.Content
	maps["status"] = data.Status
	maps["remarks"] = data.Remarks
	tx := databases.DB.Model(&Todo{}).Where("id = ?", tid).Updates(&maps)
	err = common.HandleError(tx, "所指定的记录不存在！")
	return
}
func GetTodoDetail(tid int) (err error, data Todo)  {
	tx := databases.DB.Where("id = ?", tid).First(&data)
	err = common.HandleError(tx, "所指定的记录不存在！")
	return err, data
}
