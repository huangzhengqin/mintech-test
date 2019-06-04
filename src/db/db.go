package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"model"
	"strings"
)


var SqlDB *gorm.DB

const (
	ZERO				= 0					//常量0,无意义
	DESC 				= "DESC"			//排序
	DEFAULT_KEY			= "user_name"		//模糊查询默认查询字段
)

func InsertOrder (order *model.Order)(int64, error ){
	if err := checkParam(order); err != nil {
		return ZERO, err
	}

	b := SqlDB.NewRecord(order)
	if !b {
		return ZERO, fmt.Errorf("first sqlDB.NewRecord return false. ")
	}

	db := SqlDB.Create(&order)
	if db.Error != nil {
		return ZERO, db.Error
	}

	b = SqlDB.NewRecord(order)
	if b {
		return ZERO, fmt.Errorf("insert order fail. ")
	}

	if order.Id < 1 {
		return ZERO, fmt.Errorf("return id is:%d", order.Id)
	}

	return order.Id, nil
}

func UpdateOrder (order *model.Order) error {
	if err := checkParam(order); err != nil {
		return err
	}

	db := SqlDB.Model(&order).Where("order_id=?", order.OrderId).Update(order)
	if db.Error != nil {
		return db.Error
	}

	return nil
}


func QueryByCondition (condition *model.QueryCondition) ([]*model.Order, error) {
	if SqlDB == nil {
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
		 db = SqlDB.Where(whereKey, whereValue).Order("amount " + desc).Order("create_time " + desc).Find(&orders)

	} else {
		db = SqlDB.Order("amount " + desc).Order("create_time " + desc).Find(&orders)
	}

	if db.Error != nil {
		return nil, db.Error
	}

	return orders, nil
}

func QueryByOrderId (orderId string) (*model.Order, error) {
	if SqlDB == nil {
		return nil, fmt.Errorf("sqlDB is nil. ")
	}

	if len(strings.TrimSpace(orderId)) == 0 {
		return nil, fmt.Errorf("orderId is nil. ")
	}

	order := &model.Order{}
	db := SqlDB.Where("order_id=?", orderId).Find(&order)
	if db.Error != nil {
		return nil, db.Error
	}

	return order, nil
}

func checkParam (order *model.Order) error {
	if SqlDB == nil {
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

func OpenDB () {
	db, err := gorm.Open("mysql", "root:luyun123@/test?charset=utf8")
	if err != nil {
		log.Fatalf("open the DB fail, err:%s", err)
	}

	SqlDB = db

	fmt.Println("open db success. ")
}