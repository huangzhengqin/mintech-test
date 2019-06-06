package db

import (
	"fmt"
	"model"
	"testing"
)

func TestOpenDB(t *testing.T) {
	d := OpenDB()
	defer d.Close()
}



func TestCreateOrder(t *testing.T) {
	d := OpenDB()
	defer d.Close()

	order, err := NewDbManager(d).CreateOrder(&model.Order{OrderId:"144", Amount:11, FileUrl:"dfd", UserName:"zhangsan", Status:"good"})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("id:", order)
}


func TestUpdateOrderById(t *testing.T) {
	d := OpenDB()

	err := NewDbManager(d).UpdateOrderById(&model.Order{OrderId:"124", Amount:11114, FileUrl:"1111334", UserName:"11114", Status:"11114"})
	if err != nil {
		t.Fatal(err)
	}
}


func TestGetOrderById(t *testing.T) {
	d := OpenDB()
	defer d.Close()

	ret, err := NewDbManager(d).GetOrderById(&model.Order{OrderId:"124"})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fmt.Printf("%+v",ret))

}


func TestGetOrderByCondition(t *testing.T) {
	d := OpenDB()
	defer d.Close()

	ret, err := NewDbManager(d).GetOrderByCondition(&model.QueryCondition{Key:"", LikeStr:"", Desc:true})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(ret); i++ {
		fmt.Println(fmt.Printf("%+v",ret[i]))
	}
}