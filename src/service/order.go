package service

import (
	"db"
	"model"
)

func CreateOrder (order *model.Order) (int64, error) {

	return db.InsertOrder(order)
}


func QueryByCondition (condition *model.QueryCondition) ([]*model.Order, error) {

	return db.QueryByCondition(condition)
}


func UpdateOrder (order *model.Order) error {

	return db.UpdateOrder(order)
}

func QueryByOrderId (orderId string) (*model.Order, error) {

	return db.QueryByOrderId(orderId)
}