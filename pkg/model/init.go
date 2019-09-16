package model

import (
	"time"

	common "../common"
	db "../db"
	mgo "gopkg.in/mgo.v2"
)

func Init() {

	// User Table //
	userNameIndex := mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	}
	playerCollection, session, err := db.GetCollection(playerTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		common.Terminate(err)
	}
	defer session.Close()
	playerCollection.EnsureIndex(userNameIndex)

	// Token Table //
	tokenStringIndex := mgo.Index{
		Key:    []string{"token_string"},
		Unique: true,
	}
	tokenCollection := session.DB(common.GetConfiger().Configs.MongodbName).C(tokenTable)
	tokenCollection.EnsureIndex(tokenStringIndex)

	err = tokenCollection.DropIndex("create_at")
	expireDateIndex := mgo.Index{
		Key:         []string{"create_at"},
		Background:  true,
		ExpireAfter: time.Duration(ExpireSecond) * time.Second,
	}
	tokenCollection.EnsureIndex(expireDateIndex)

	// Agent Table //
	agentUserNameIndex := mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	}
	agentCollection := session.DB(common.GetConfiger().Configs.MongodbName).C(agentTable)
	agentCollection.EnsureIndex(agentUserNameIndex)

	// Play Table //
	expireDateIndex = mgo.Index{
		Key:         []string{"expire_date"},
		Background:  true,
		ExpireAfter: time.Duration(ExpireSecond) * time.Second,
	}
	playCollection := session.DB(common.GetConfiger().Configs.MongodbName).C(playTable)
	playCollection.EnsureIndex(expireDateIndex)

}

func Clear() {
	ClearToken()
	ClearPlayer()
	ClearGame()
	ClearWallet()
	ClearAgent()
	ClearTransaction()

}
