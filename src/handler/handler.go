package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"model"
	"net/http"
	"os"
	"service"
	"strings"
)

func Upload(c *gin.Context) {
	name := c.PostForm("upload")
	fmt.Println("文件名：", name)

	file, header, err := c.Request.FormFile("upload")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}


	filename := header.Filename
	fmt.Println(file, err, filename)

	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()


	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusCreated, "upload successful.")
}

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
	orderId := c.Param("order_id")

	if len(strings.TrimSpace(orderId)) == 0 {
		c.String(http.StatusOK,"orderId is nil. ", orderId)
		return
	}

	order, err := service.QueryByOrderId(orderId)
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


