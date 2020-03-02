package main

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/crgimenes/goconfig"
	_ "github.com/crgimenes/goconfig/ini"
)

/*
step 1: Declare your configuration struct,
it may or may not contain substructures.
*/

type systemUser struct {
	Name     string `json:"name" ini:"name" cfg:"name"`
	Password string `json:"passwd" ini:"passwd" cfg:"passwd"`
}

type systemSqs struct {
	ReceiveQueue string `json:"receive_queue" ini:"receive_queue" cfg:"receive_queue"`
	SendQueue string `json:"send_queue" ini:"send_queue" cfg:"send_queue"`
}

type configTest struct {
	DebugMode  bool       `json:"debug" ini:"debug" cfg:"debug" cfgDefault:"false"`
	ServerName string     `json:"server_name" ini:"server_name" cfg:"server_name" cfgDefault:"local"`
	User       systemUser `json:"user" ini:"user" cfg:"user"`
	Sqs        systemSqs `json:"sqs" ini:"sqs" cfg:"sqs"`
}

func getConfigFilePath(env string) string {
	//_, dirname, _, _ := runtime.Caller(1)
	path, _ := os.Getwd()
	filename := []string{path, "\\config\\config.", env, ".ini"}
	result := strings.Join(filename, "")
	//println("filename : ", filename)
	println("result : ", result)
	return result
}

func confgTestFunc(env string) *configTest {
	//env := flag.String("env", "dev1", "a string")
	config := configTest{}

	goconfig.File = getConfigFilePath(env)
	err := goconfig.Parse(&config)
	if err != nil {
		panic(err)
	}
	//println("configTestFunc MarshalIndent")
	// just print struct on screen
	j, _ := json.MarshalIndent(config, "", "\t")

	//obj := json.Unmarshal(config)
	//println("configTestFunc End")
	println(string(j))
	return &config
}
