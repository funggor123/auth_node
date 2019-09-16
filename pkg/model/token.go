package model

import (
	"time"

	auth "../auth"
	common "../common"
	db "../db"
	"gopkg.in/mgo.v2/bson"
)

type Token struct {
	TokenString string        `bson:"token_string" `
	CreateAt    time.Time     `bson:"create_at"`
	AgentID     bson.ObjectId `bson:"agent_id" `
}

func CreateTokenInDB(agentID bson.ObjectId) (Token, error) {
	tokenCollection, session, err := db.GetCollection(tokenTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return Token{}, err
	}
	defer session.Close()

	token := Token{TokenString: auth.GenerateToken(), AgentID: agentID, CreateAt: time.Now()}
	err = tokenCollection.Insert(token)
	if err != nil {
		return Token{}, err
	}
	return token, nil
}

func GetExpireDate(createAt time.Time, durationInSec int) time.Time {
	duration := time.Duration(durationInSec) * time.Second
	expireAt := createAt.Add(duration)
	return expireAt
}

func ClearToken() error {
	err := db.ClearCollections(tokenTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return err
	}
	return nil
}

func GetTokenFromDB(tokenString string) ([]Token, error) {
	var tokens []Token
	tokenCollection, session, err := db.GetCollection(tokenTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	tokenCollection.Find(bson.M{"token_string": tokenString}).All(&tokens)
	return tokens, nil
}
