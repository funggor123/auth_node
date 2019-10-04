package model

import (
	"fmt"
	"time"

	common "../common"
	db "../db"
	"gopkg.in/mgo.v2/bson"
)

type GameRecord struct {
	ID           bson.ObjectId `bson:"id" json:"id"`
	PlayerID     bson.ObjectId `bson:"player_id" json:"player_id"`
	GameDateTime time.Time     `bson:"game_date_time" json:"game_date_time"`
	MoneyRemain  float32       `bson:"money_remain" json:"money_remain"`
	MoneyWin     float32       `bson:"money_win" json:"money_win"`
	Decision     int           `bson:"decision" json:"decision"`
}

const game_record_table = "game_record"

func InsertGameRecord(gameRecord GameRecord) ([]GameRecord, error) {
	gameRecordCollection, session, err := db.GetCollection(game_record_table, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return nil, err
	}
	defer session.Close()
	gameRecord.GameDateTime = time.Now()
	gameRecord.ID = bson.NewObjectId()
	err = gameRecordCollection.Insert(gameRecord)
	if err != nil {
		fmt.Println("err")
		return nil, err
	}
	fmt.Println("awdaw")
	fmt.Println(gameRecord)
	return nil, nil
}

func GetGameRecord(PlayerID string) ([]GameRecord, error) {
	var gameRecords []GameRecord
	gameRecordCollection, session, err := db.GetCollection(game_record_table, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return nil, err
	}
	defer session.Close()

	gameRecordCollection.Find(bson.M{"player_id": bson.ObjectIdHex(PlayerID)}).All(&gameRecords)
	return gameRecords, nil
}
