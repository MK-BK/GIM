package models

type Environment struct {
	Comments
	Groups
	Moments
	Sessions
	Requests
	Users
	RoleAssignments
}

var GlobalEnvironment Environment

type Configure struct {
	MysqlUser  string `json:"mysql_username"`
	MysqlPwd   string `json:"mysql_password"`
	MysqlDB    string `json:"mysql_dbname"`
	MysqlHost  string `json:"mysql_host"`
	RedisHost  string `json:"redis_address"`
	RedisPwd   string `json:"redis_password"`
	ListenPort string `json:"listen_port"`
	SecretKey  string `json:"secretkey"`
	DataPath   string `json:"data_path"`
}

var Conf = &Configure{}
