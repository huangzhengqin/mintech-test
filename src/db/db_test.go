package db

import (
	"fmt"
	"model"
	"testing"
)

func TestOpenDB(t *testing.T) {
	OpenDB()
}



func TestInsertOrder(t *testing.T) {
	OpenDB()

	id, err := InsertOrder(&model.Order{OrderId:"12344", Amount:11, FileUrl:"dfd", UserName:"zhangsan", Status:"good"})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("id:", id)
}


func TestUpdateOrder(t *testing.T) {
	OpenDB()

	err := UpdateOrder(&model.Order{OrderId:"11111", Amount:11114, FileUrl:"1111334", UserName:"11114", Status:"11114"})
	if err != nil {
		t.Fatal(err)
	}
}


func TestQueryByOrderId(t *testing.T) {
	OpenDB()

	ret, err := QueryByOrderId("11111")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fmt.Printf("%+v",ret))

}


func TestQueryByCondition(t *testing.T) {
	OpenDB()

	ret, err := QueryByCondition(&model.QueryCondition{Key:"", LikeStr:"", Desc:true})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(ret); i++ {
		fmt.Println(fmt.Printf("%+v",ret[i]))
	}
}