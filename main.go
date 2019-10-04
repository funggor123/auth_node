package main

import (
	"math/rand"
	"time"

	common "./pkg/common"
	db "./pkg/db"
	model "./pkg/model"
	node "./pkg/node"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	common.GetConfiger().ReadConfigs()

	common.GetConfiger().CheckPathsAndFilesExists()

	common.GetLogger().SetLogFilePath(common.GetConfiger().Configs.LogDir)

	db.CheckMongoDBConnection()

	//model.Clear()

	//model.Init()

	model.CreateDefaultAgent()
	model.CreateDefaultGame()

	common.GetConfiger().PrintConfigs()

	node.StartRouter()

}
