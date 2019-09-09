package db

import (
	common "../common"
	mgo "gopkg.in/mgo.v2"
)

var globalSession *mgo.Session

func getSession() (*mgo.Session, error) {
	var err error
	if globalSession != nil {
		return globalSession, nil
	}
	globalSession, err = mgo.Dial(MONGO_PRE_FIX + common.GetConfiger().Configs.MongodbURL)
	if err != nil {
		return nil, err
	}
	return globalSession, nil
}

func CheckMongoDBConnection() {
	_, err := getSession()
	if err != nil {
		common.Terminate(err)
	}
}

func GetCollection(CollectionName string, dbName string) (*mgo.Collection, *mgo.Session, error) {
	session, err := getSession()
	if err != nil {
		return nil, nil, err
	}
	newSession := session.Copy()
	return newSession.DB(dbName).C(CollectionName), newSession, nil
}

func ClearCollections(CollectionName string, dbName string) error {
	session, err := getSession()
	if err != nil {
		return err
	}
	newSession := session.Copy()
	defer newSession.Close()
	if _, err = newSession.DB(dbName).C(CollectionName).RemoveAll(nil); err != nil {
		return err
	}
	return nil
}

