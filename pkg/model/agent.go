package model

import (
	"fmt"

	common "../common"
	db "../db"
	"gopkg.in/mgo.v2/bson"
)

type Agent struct {
	ID        bson.ObjectId `bson:"agent_id" `
	AgentName string        `bson:"agent_name" `
	Password  string        `bson:"password" `
	UserName  string        `bson:"username" `
	Currenies []string      `bson:"currenies" `
	Locales   []string      `bson:"locales" `
}

func GetAgentFromDB(username string, password string) ([]Agent, error) {
	var agents []Agent
	agentCollection, session, err := db.GetCollection(agentTable, common.GetConfiger().Configs.MongodbName)
	defer session.Close()
	if err != nil {
		return nil, err
	}
	agentCollection.Find(bson.M{"username": username, "password": password}).All(&agents)
	return agents, nil
}

func GetAgentByIDFromDB(agentId string) ([]Agent, error) {
	var agents []Agent
	agentCollection, session, err := db.GetCollection(agentTable, common.GetConfiger().Configs.MongodbName)
	defer session.Close()
	if err != nil {
		return nil, err
	}
	agentCollection.Find(bson.M{"agent_id": bson.ObjectIdHex(agentId)}).All(&agents)
	return agents, nil
}

func CreateAgentInDB(agent Agent) error {
	agentCollection, session, err := db.GetCollection(agentTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return err
	}
	defer session.Close()
	agent.ID = bson.ObjectIdHex("5d8dacfb56c00a034cc0635d")
	err = agentCollection.Insert(agent)
	if err != nil {
		return err
	}
	return nil
}

func CreateDefaultAgent() {
	CreateAgentInDB(Agent{UserName: default_agent_username, Password: default_agent_password, AgentName: "testing"})
	agent, _ := GetAgentFromDB(default_agent_username, default_agent_password)
	fmt.Println("test", agent[0].ID)
	return
}

func ClearAgent() error {
	err := db.ClearCollections(agentTable, common.GetConfiger().Configs.MongodbName)
	if err != nil {
		return err
	}
	return nil
}
