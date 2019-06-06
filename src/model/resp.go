/*
 *@author Novice.Huang
 *@date 19-6-5
 */

package model

const (
	STATUS_FAIL			= 1
	STATUS_SUCCESS		= 0
)

type Result struct {
	STATUS	 	int		`gorm:"status" json:"status"`		//状态：0表示成功， 1表示失败
	Data		string	`gorm:"data" json:"data"`			//返回数据，状态为0时有值
	ErrorStr	string	`jorm:"error_str" json:"error_str"`	//异常数据，状态为1时有值
	Message		string	`gorm:"message json:"message""`		//备注
}