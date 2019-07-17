package router

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mintech-test/src/db"
	"mintech-test/src/model"
	"mintech-test/src/service"
	"net/http"
	"strings"
)

type Server struct {
	ss *service.ServiceManager
}

func NewService(d *gorm.DB) *Server {
	return &Server{
		ss: service.NewServiceManager(db.NewDbManager(d)),
	}
}

func (s *Server) New() *gin.Engine {
	router := gin.Default()

	v2 := router.Group("/v2")
	v2.POST("/order", s.Create)
	v2.GET("/order/:order_id", s.QueryById)
	v2.PUT("/order", s.Update)
	v2.GET("/order", s.Query)

	return router
}

func (s *Server) Query(c *gin.Context) {
	var req model.OrderConditionReq
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: fmt.Sprintf("param error:%s", err)})
		return
	}

	ret, err := s.ss.GetOrderByCondition(&req.QueryCondition)
	if err != nil {
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: err.Error()})
		return
	}

	b, err := json.Marshal(ret)
	if err != nil {
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: fmt.Sprintf("解析返回结果失败,error:%s", err)})
		return
	}

	c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_SUCCESS, Data: string(b)})
}

func (s *Server) QueryById(c *gin.Context) {
	orderId := c.Param("order_id")

	if len(strings.TrimSpace(orderId)) == 0 {
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: "orderId is nil."})
		return
	}

	order, err := s.ss.GetOrderById(&model.Order{OrderId: orderId})
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: err.Error()})
		return
	}

	b, err := json.Marshal(order)
	if err != nil {
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: fmt.Sprintf("解析返回结果失败,error:%s", err)})
		return
	}

	c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_SUCCESS, Data: string(b)})
}

func (s *Server) Update(c *gin.Context) {
	var req model.OrderReq
	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.String(http.StatusNotFound, "param error:%s", err)
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: fmt.Sprintf("param error:%s", err)})
		return
	}

	err := s.ss.UpdateOrderById(&req.Order)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: err.Error()})
		return
	}

	c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_SUCCESS, Message: "update success."})
}

func (s *Server) Create(c *gin.Context) {
	var req model.OrderReq

	if err := c.BindJSON(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: err.Error()})
		return
	}

	_, err := s.ss.CreateOrder(&req.Order)
	if err != nil {
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: err.Error()})
		return
	}

	b, err := json.Marshal(req.Order)
	if err != nil {
		c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_FAIL, ErrorStr: fmt.Sprintf("解析返回结果失败,error:%s", err)})
		return
	}

	c.JSON(http.StatusOK, &model.Result{STATUS: model.STATUS_SUCCESS, Data: string(b)})

	fmt.Println("create order success.")
}
