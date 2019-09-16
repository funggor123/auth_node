package model

import (
	"errors"

	common "../common"
	db "../db"
	"gopkg.in/mgo.v2/bson"
)

type Player struct {
	AgentID        bson.ObjectId `bson:"agentId" `
	ActiveSessions string        `bson:"active_sessions" `
	FirstName      string        `bson:"firstName" `
	LastName       string        `bson:"lastName" `

	Name     string `bson:"name" `
	UserName string `bson:"username" `
	Email    string `bson:"email" `
	Gender   string `bson:"gender" `

	DayPhone     string `bson:"dayPhone" `
	EveningPhone string `bson:"eveningPhone" `
	CellPhone    string `bson:"cellPhone" `

	BirthDate     string `bson:"birthdate" `
	PlayerClassID string `bson:"playerClassId" `
	CountryID     string `bson:"countryId" `

	Currency string        `bson:"currency" `
	WalletID bson.ObjectId `bson:"walletId" `
	ID       bson.ObjectId `bson:"id"`
}

func CreatePlayerInDb(player Player) (Player, error) {
	playerCollection, session, err := db.GetCollection(playerTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return Player{}, err
	}
	defer session.Close()
	player.ID = bson.NewObjectId()
	err = playerCollection.Insert(player)
	if err != nil {
		return Player{}, err
	}
	return player, nil
}

func GetPlayerByIDFromDB(playerID string) ([]Player, error) {
	var players []Player
	playerCollection, session, err := db.GetCollection(playerTable, common.GetConfiger().Configs.MongodbName)
	defer session.Close()
	if err != nil {
		return nil, err
	}

	if !bson.IsObjectIdHex(playerID) {
		return nil, errors.New("Invalid player id ")
	}
	playerCollection.Find(bson.M{"id": bson.ObjectIdHex(playerID)}).All(&players)
	return players, nil
}

func ClearPlayer() error {
	err := db.ClearCollections(playerTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return err
	}
	return nil
}
