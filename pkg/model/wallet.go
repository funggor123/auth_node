package model

import (
	db "../db"
	common "../common"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type Wallet struct {
	ID 			bson.ObjectId  `bson:"id" `
	Money		int  `bson:"money" `

}

func CreateWalletnInDB() (Wallet, error){
	walletCollection, session, err := db.GetCollection(walletTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return  Wallet{}, err
	}
	defer session.Close()
	wallet := Wallet{ID: bson.NewObjectId()}
	err = walletCollection.Insert(wallet)
	if err != nil {
		return Wallet{}, err
	}
	return wallet, nil
}

func ClearWallet() error {
	err:= db.ClearCollections(walletTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return err
	}
	return nil
}

func GetWalletByIDFromDB(walletID bson.ObjectId) ([]Wallet, error){
	var wallets []Wallet
	walletCollection, session, err := db.GetCollection(walletTable, common.GetConfiger().Configs.MongodbName)
	defer session.Close()
	if err != nil {
		return nil, err
	}
	walletCollection.Find(bson.M{"id": walletID}).All(&wallets)
	return wallets, nil
}

func UpdateWalletMoneyInDB(walletID bson.ObjectId, money int) (int, error){
	var wallets []Wallet
	walletCollection, session, err := db.GetCollection(walletTable, common.GetConfiger().Configs.MongodbName)
	defer session.Close()
	if err != nil {
		return wallets[0].Money, nil
	}
	walletCollection.Find(bson.M{"id": walletID}).All(&wallets)

	if wallets[0].Money + money < 0 {
		return wallets[0].Money, errors.New("The amount of money in the wallet cannot be negative after withdrawal")
	}

	selector := bson.M{"id": walletID}
	err = walletCollection.Update(selector, bson.M{"$set": bson.M{"money": wallets[0].Money + money}})
	return wallets[0].Money + money, nil
}

