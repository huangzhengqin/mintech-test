/*
 *@author Novice.Huang
 *@date 19-6-6
 */

package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"model"
	"strings"
)

const (
	ZERO				= 0					//常量0,无意义
	DESC 				= "DESC"			//排序
	DEFAULT_KEY			= "user_name"		//模糊查询默认查询字段
)

func NewDbManager(d *gorm.DB) *DbManager {
	 return &DbManager{
	 	MysqlDB:d,
	 }
}

type DbManager struct {
	//db 			*model.Order
	MysqlDB		*gorm.DB
}

func (m *DbManager) CreateOrder(order *model.Order) (o *model.Order, err error){
	if err = m.checkParamV2(order); err != nil {
		return
	}

	tx := m.MysqlDB.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// if order_id exist, delete it
	if !tx.Where("order_id=?", order.OrderId).RecordNotFound() {
		if er := tx.Table("orders").Where("order_id=?", order.OrderId).Delete(&model.Order{}).Error; er != nil {
			err = er
			return
		}

		fmt.Println("order_id exist, delete success.")
	}

	b := tx.NewRecord(order)
	if !b {
		err = fmt.Errorf("first sqlDB.NewRecord return false. ")
		return
	}

	db := tx.Create(&order)
	if db.Error != nil {
		err = db.Error
		return
	}

	b = tx.NewRecord(order)
	if b {
		err = fmt.Errorf("insert order fail. ")
		return
	}

	if order.Id < 1 {
		err = fmt.Errorf("return id is:%d", order.Id)
		return
	}

	return order, nil
}

func (m *DbManager) UpdateOrderById(order *model.Order) error {
	if err := m.checkParamV2(order); err != nil {
		return err
	}

	_, err := m.GetOrderById(order)
	if err != nil {
		return err
	}

	db := m.MysqlDB.Model(&order).Where("order_id=?", order.OrderId).Update(order)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (m *DbManager) GetOrderByCondition(condition *model.QueryCondition) ([]*model.Order, error) {
	if m.MysqlDB == nil {
		return nil, fmt.Errorf("sqlDB is nil. ")
	}

	if condition == nil {
		return nil, fmt.Errorf("condition is nil. ")
	}

	whereKey 	:= ""
	whereValue 	:= ""
	whereFlag 	:= false
	if len(strings.TrimSpace(condition.LikeStr)) > 0 {
		whereKey 	= fmt.Sprintf("%s LIKE ?", condition.Key)
		whereValue 	= "%"+ condition.LikeStr+ "%"
		whereFlag 	= true
	}

	if len(strings.TrimSpace(condition.Key)) == 0 {
		condition.Key  = DEFAULT_KEY
	}

	desc := ""
	if condition.Desc {
		desc = DESC
	}

	orders 	:= make([]*model.Order, 0)
	db		:= &gorm.DB{}
	if whereFlag {
		db = m.MysqlDB.Where(whereKey, whereValue).Order("amount " + desc).Order("create_time " + desc).Find(&orders)

	} else {
		db = m.MysqlDB.Order("amount " + desc).Order("create_time " + desc).Find(&orders)
	}

	if db.Error != nil {
		return nil, db.Error
	}

	return orders, nil
}

func (m *DbManager) GetOrderById(o *model.Order) (*model.Order, error) {
	if m.MysqlDB == nil {
		return nil, fmt.Errorf("sqlDB is nil. ")
	}

	if len(strings.TrimSpace(o.OrderId)) == 0 {
		return nil, fmt.Errorf("orderId is nil. ")
	}

	order := &model.Order{}
	db := m.MysqlDB.Where("order_id=?", o.OrderId).Find(&order)
	if db.RecordNotFound() {
		return nil, fmt.Errorf("没有查到相关信息, order_id:%s", o.OrderId)
	}
	if db.Error != nil {
		return nil, db.Error
	}

	return order, nil
}

func (m *DbManager) checkParamV2 (order *model.Order) error {
	if m.MysqlDB == nil {
		return fmt.Errorf("sqlDB is nil. ")
	}

	if order == nil {
		return fmt.Errorf("order is nil. ")
	}

	if strings.TrimSpace(order.OrderId) == "" || len(order.OrderId) > 30 {
		return fmt.Errorf("OrderId is error, OrderId:%s", order.OrderId)
	}

	if strings.TrimSpace(order.UserName) == "" || len(order.UserName) > 30 {
		return fmt.Errorf("UserName is error, UserName:%s", order.UserName)
	}

	if strings.TrimSpace(order.Status) == "" || len(order.Status) > 30 {
		return fmt.Errorf("Status is error, Status:%s ", order.Status)
	}

	if strings.TrimSpace(order.FileUrl) == "" || len(order.FileUrl) > 200 {
		return fmt.Errorf("FileUrl is error, FileUrl:%s", order.FileUrl)
	}

	if order.Amount < 0 {
		return fmt.Errorf("Amount if error, Amount:%f ", order.Amount)
	}

	return nil
}


