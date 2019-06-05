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
		c.String(http.StatusBadRequest, "Bad request. ")
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
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:err.Error()})
		return
	}

	_, err := service.CreateOrder(&req.Order)
	if err != nil {
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:err.Error()})
		return
	}

	b, err := json.Marshal(req.Order)
	if err != nil {
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:fmt.Sprintf("解析返回结果失败,error:%s", err)})
		return
	}

	c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_SUCCESS, Data:string(b)})

	fmt.Println("create order success.")
}

func QueryByOrderId(c *gin.Context) {
	orderId := c.Param("order_id")

	if len(strings.TrimSpace(orderId)) == 0 {
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:"orderId is nil."})
		return
	}

	order, err := service.QueryByOrderId(orderId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:err.Error()})
		return
	}

	b, err := json.Marshal(order)
	if err != nil {
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:fmt.Sprintf("解析返回结果失败,error:%s", err)})
		return
	}

	c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_SUCCESS, Data:string(b)})
}

func UpdateOrder(c *gin.Context) {
	var req  model.OrderReq
	if err:=c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.String(http.StatusNotFound, "param error:%s", err)
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:fmt.Sprintf("param error:%s", err)})
		return
	}

	err := service.UpdateOrder(&req.Order)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:err.Error()})
		return
	}

	c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_SUCCESS, Message:"update success."})
}

func Query(c *gin.Context) {
	var req  model.OrderConditionReq
	if err:=c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:fmt.Sprintf("param error:%s", err)})
		return
	}

	ret, err := service.QueryByCondition(&req.QueryCondition)
	if err != nil {
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:err.Error()})
		return
	}

	b, err := json.Marshal(ret)
	if err != nil {
		c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_FAIL, ErrorStr:fmt.Sprintf("解析返回结果失败,error:%s", err)})
		return
	}

	c.JSON(http.StatusOK, &model.Result{STATUS:model.STATUS_SUCCESS, Data:string(b)})
}


