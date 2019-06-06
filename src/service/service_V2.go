/*
 *@author Novice.Huang
 *@date 19-6-6
 */

package service

import (
	"model"
)

type DbManager interface {
	GetOrderById(o *model.Order)							(*model.Order, error)
	GetOrderByCondition(condition *model.QueryCondition)	([]*model.Order, error)
	UpdateOrderById(order *model.Order)						error
	CreateOrder(order *model.Order)							(*model.Order, error)
}

type ServiceOrder interface {
	DbManager
	Delete()	(string, error)
}

func NewServiceManager(d DbManager) *ServiceManager {
	return &ServiceManager{
		db:d,
		Remark:"test",
	}
}

type ServiceManager struct {
	db 			DbManager
	Remark		string
}

func (s *ServiceManager) GetOrderById(order *model.Order) (*model.Order, error) {

	return s.db.GetOrderById(order)
}

func (s *ServiceManager) GetOrderByCondition(condition *model.QueryCondition) ([]*model.Order, error) {

	return s.db.GetOrderByCondition(condition)
}


func (s *ServiceManager) UpdateOrderById(order *model.Order) error {

	return s.db.UpdateOrderById(order)
}

func (s *ServiceManager) CreateOrder(order *model.Order) (*model.Order, error) {

	return s.db.CreateOrder(order)
}

func (s *ServiceManager) Delete ()(string, error) {

	return s.Remark, nil
}
