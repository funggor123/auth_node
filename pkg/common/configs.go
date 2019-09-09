package common

type Configs struct {
	Port        string `json:"port" binding:"required"`
	LogDir      string `json:"log_dir" binding:"required"`
	MongodbURL  string `json:"mongodb_url" binding:"required"`
	MongodbName string `json:"mongodb_db_name" binding:"required"`
	GameNodeAddress string `json:"game_node_address" binding:"required"`
}
