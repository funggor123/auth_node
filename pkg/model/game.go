package model

import (
	"gopkg.in/mgo.v2/bson"
	db "../db"
	common "../common"
	"fmt"
)

type Game struct {
	ID        			bson.ObjectId `bson:"id" `
	Locale      		string `bson:"locale" `
	Name       			string `bson:"name" `
	Description 	   	string  `bson:"description" `
	LongDescription 	string  `bson:"long_description" `
}

func GetGameFromDB(locale string) ([]Game, error){
	var games []Game
	gameCollection, session, err := db.GetCollection(gameTable, common.GetConfiger().Configs.MongodbName)
	defer session.Close()
	if err != nil {
		return nil, err
	}
	gameCollection.Find(bson.M{"locale": locale}).All(&games)
	return games, nil
}

func GetGameByIDFromDB(GameID string) ([]Game, error){
	var games []Game
	gameCollection, session, err := db.GetCollection(gameTable, common.GetConfiger().Configs.MongodbName)
	defer session.Close()
	if err != nil {
		return nil, err
	}
	gameCollection.Find(bson.M{"id": bson.ObjectIdHex(GameID)}).All(&games)
	return games, nil
}


func CreateGameInDB(game Game) error {
	gameCollection, session, err := db.GetCollection(gameTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return err
	}
	defer session.Close()
	game.ID = bson.NewObjectId()
	err = gameCollection.Insert(game)
	if err != nil {
		return err
	}
	return nil
}

func CreateDefaultGame() {
	CreateGameInDB(Game{Name: "testGame", Locale: "cn-CN"})
	game, _ := GetGameFromDB("cn-CN")
	fmt.Println("game", game[0].ID)
	return 
}

func ClearGame() error {
	err:= db.ClearCollections(gameTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return err
	}
	return nil
}
