package model

type Order struct {
	Id				int64			`json:"id"`
	OrderId       	string			`json:"order_id"`
	UserName      	string 			`json:"user_name"`
	Amount         	float64			`json:"amount"`
	Status         	string			`json:"status"`
	FileUrl       	string			`json:"file_url"`
}


type QueryCondition struct {
	Key 			string		//查询字段
	LikeStr			string		//关键字
	Desc			bool		//是否使用降序
}

