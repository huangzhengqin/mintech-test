package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"model"
	"net/http"
	"service"
	"strings"
)

func CreateOrder(c *gin.Context) {
	var req  model.OrderReq

	if err:=c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.String(http.StatusNotFound, "param error:%s", err)
		return
	}

	_, err := service.CreateOrder(&req.Order)
	if err != nil {
		c.String(http.StatusOK, "error:%s", err)
		return
	}

	b, err := json.Marshal(req.Order)
	if err != nil {
		c.String(http.StatusOK, "error:%s", err)
		return
	}

	c.Data(http.StatusOK, "application/json", b)

	fmt.Println("create order success.")
}

func QueryByOrderId(c *gin.Context) {
	var req  model.OrderReq
	if err:=c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.String(http.StatusNotFound, "param error:%s", err)
		return
	}

	if len(strings.TrimSpace(req.OrderId)) == 0 {
		c.String(http.StatusOK,"orderId is nil. ", req.OrderId)
		return
	}

	order, err := service.QueryByOrderId(req.OrderId)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError,"service error:%s ", err)
		return
	}

	b, err := json.Marshal(order)
	if err != nil {
		c.String(http.StatusOK, "error:%s", err)
		return
	}

	c.Data(http.StatusOK, "application/json", b)
}

func UpdateOrder(c *gin.Context) {
	var req  model.OrderReq
	if err:=c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.String(http.StatusNotFound, "param error:%s", err)
		return
	}

	err := service.UpdateOrder(&req.Order)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusInternalServerError,"service error:%s ", err)
		return
	}

	c.String(http.StatusOK, "update success. ")
}

func Query(c *gin.Context) {
	var req  model.OrderConditionReq
	if err:=c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.String(http.StatusNotFound, "param error:%s", err)
		return
	}

	ret, err := service.QueryByCondition(&req.QueryCondition)
	if err != nil {
		c.String(http.StatusInternalServerError,"service error:%s ", err)
		return
	}

	b, err := json.Marshal(ret)
	if err != nil {
		c.String(http.StatusOK, "error:%s", err)
		return
	}

	c.Data(http.StatusOK, "application/json", b)
}


