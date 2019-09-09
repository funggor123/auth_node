package model

import (
	"time"
	mgo "gopkg.in/mgo.v2"
	db "../db"
	common "../common"
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

	expireDateIndex := mgo.Index{
        Key:    []string{"expire_date"},
		Background:  true,
		ExpireAfter: time.Duration(expireSecond) * time.Second,
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
        Key:    []string{"expire_date"},
		Background:  true,
		ExpireAfter: time.Duration(expireSecond) * time.Second,
	}
	playCollection := session.DB(common.GetConfiger().Configs.MongodbName).C(playTable)
	playCollection.EnsureIndex(expireDateIndex)

}

func Clear(){
	ClearToken()
	ClearPlayer()
	ClearGame()
	ClearWallet()
	ClearAgent()
}
