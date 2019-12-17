package config

import (
	"encoding/json"
	"log"
	"os"
)

type ServerConfig struct {
	Address      string
	ReadTimeout  int64
	WriteTimeout int64
	Static       string
	TemplateBase string
	Version      string
}

type DbConfig struct {
	UserName string
	Password string
	IP       string
	DbName   string
}

type Config struct {
	ServerConfig ServerConfig
	DbConfig     DbConfig
}

var AppConfig Config

var logger *log.Logger

func init() {
	loadConfig()
	loadLogger()
}

func loadConfig() {
	file, err := os.Open("C:\\gaonl\\study\\golang\\goweb\\src\\config\\config.json")
	if err != nil {
		log.Fatalln("Failed to open config file", err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalln("Unable to get configuration from file", err)
	}
}

func loadLogger() {
	file, err := os.OpenFile("chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func Error(args ...interface{}) {
	logger.SetPrefix("Error ")
	logger.Println(args...)
}

func Warning(args ...interface{}) {
	logger.SetPrefix("Warning ")
	logger.Println(args...)
}
