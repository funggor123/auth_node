package model

import (
	common "../common"
	db "../db"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Transaction struct {
	ID        			bson.ObjectId `bson:"id" `
	PlayerID			bson.ObjectId  `bson:"player_id" `	
	MoneyRemain			int  `bson:"money_remain" `	
	MoneyExchange		int  `bson:"money_exchange" `	
	TransactionDate     time.Time `bson:"trans_date"`
}

func GetTransactionByIDFromDB(PlayerID string) ([]Transaction, error){
	var trans []Transaction
	transCollection, session, err := db.GetCollection(transactionTable, common.GetConfiger().Configs.MongodbName)
	defer session.Close()
	if err != nil {
		return nil, err
	}
	transCollection.Find(bson.M{"player_id": bson.ObjectIdHex(PlayerID)}).All(&trans)
	return trans, nil
}


func CreateTransactionInDB(trans Transaction) (Transaction, error) {
	transCollection, session, err := db.GetCollection(transactionTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return Transaction{}, err
	}
	defer session.Close()
	trans.ID = bson.NewObjectId()
	trans.TransactionDate = time.Now()
	err = transCollection.Insert(trans)
	if err != nil {
		return Transaction{}, err
	}
	return trans, nil
}


func ClearTransaction() error {
	err:= db.ClearCollections(transactionTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return err
	}
	return nil
}
