package model

import (
	auth "../auth"
	db "../db"
	common "../common"
	"gopkg.in/mgo.v2/bson"
)

type Token struct {
	TokenString 	string  `bson:"token_string" `
	ExpireDate		string  `bson:"expire_date" `
	AgentID			bson.ObjectId  `bson:"agent_id" `
}

func CreateTokenInDB(agentID bson.ObjectId) (string,error) {
	tokenCollection, session, err := db.GetCollection(tokenTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return "", err
	}
	defer session.Close()
	
	tokenString := auth.GenerateToken()
	err = tokenCollection.Insert(Token{TokenString: tokenString, AgentID: agentID})
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ClearToken() error {
	err:= db.ClearCollections(tokenTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return err
	}
	return nil
}
