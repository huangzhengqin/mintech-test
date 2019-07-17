github:https://github.com/hzq-qiyuan


create table sql:
create table orders (
	id int(11) NOT NULL AUTO_INCREMENT,
	order_id varchar(30) DEFAULT NULL,
	user_name varchar(30) DEFAULT NULL,
	status varchar(30) DEFAULT NULL,
	file_url varchar(200) DEFAULT NULL,
	amount float DEFAULT NULL,
	PRIMARY KEY (id)
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;



根据实际情况修改数据库密码,默认使用:luyun123