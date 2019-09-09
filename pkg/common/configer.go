package common

import (
	"encoding/json"
	"errors"
	"os"

	e "../e"
)

var configer *Configer

type Configer struct {
	Configs Configs
}

func GetConfiger() *Configer {
	if configer != nil {
		return configer
	}
	configer = new(Configer)
	return configer
}

func (configer *Configer) ReadConfigs() {
	file, err := os.Open(CONFIG_FILE_NAME)
	defer file.Close()
	if err != nil {
		Terminate(err)
	}

	decoder := json.NewDecoder(file)
	Configuration := Configs{}
	if err = decoder.Decode(&Configuration); err != nil {
		Terminate(err)
	}

	configer.Configs = Configuration
}

func (configer Configer) CheckPathsAndFilesExists() {

	exists := true

	_ = IsDirExists(configer.Configs.LogDir, true)

	if !exists {
		Terminate(errors.New("Some Paths or Files specified in Configuration file are not exists"))
	}
}

func (configer Configer) PrintConfigs() {
	GetLogger().Log(e.TRACK, configer.Configs)
}
