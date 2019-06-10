package service

import (
	"db"
	"fmt"
	"model"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	d := db.OpenDB()

	id, err := NewServiceManager(db.NewDbManager(d)).CreateOrder(&model.Order{OrderId:"12343", Amount:11, FileUrl:"dfd", UserName:"zhangsan", Status:"good"})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("id:",id)
}


func TestUpdateOrderById(t *testing.T) {
	d := db.OpenDB()

	err := NewServiceManager(db.NewDbManager(d)).UpdateOrderById(&model.Order{OrderId:"11111", Amount:1, FileUrl:"1111334", UserName:"11114", Status:"11114"})
	if err != nil {
		t.Fatal(err)
	}
}


func TestGetOrderById(t *testing.T) {
	d := db.OpenDB()

	ret, err := NewServiceManager(db.NewDbManager(d)).GetOrderById(&model.Order{OrderId:"124"})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fmt.Printf("%+v",ret))

}


func TestGetOrderByCondition(t *testing.T) {
	d := db.OpenDB()

	ret, err := NewServiceManager(db.NewDbManager(d)).GetOrderByCondition(&model.QueryCondition{Key:"", LikeStr:"", Desc:true})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(ret); i++ {
		fmt.Println(fmt.Printf("%+v",ret[i]))
	}
}