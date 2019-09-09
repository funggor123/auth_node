package model


import (
	"gopkg.in/mgo.v2/bson"
	db "../db"
	common "../common"
	auth "../auth"
)


type Play struct {
	TokenString 	string  `bson:"token_string" `
	ExpireDate		string  `bson:"expire_date" `
	PlayerID		bson.ObjectId 	`bson:"player_id" `	
	ID				bson.ObjectId  `bson:"id" `	
	GameID 			bson.ObjectId 	`bson:"game_id" `	
	ReturnURL       string  `bson:"return_url" `	
}

func CreatePlayInDB(play Play) (Play, error) {
	playCollection, session, err := db.GetCollection(playTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return Play{}, err
	}
	defer session.Close()
	tokenString := auth.GenerateToken()
	play.ID = bson.NewObjectId()
	play.TokenString = tokenString
	err = playCollection.Insert(play)
	if err != nil {
		return Play{}, err
	}
	return play, nil
}

