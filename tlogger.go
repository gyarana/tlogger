package main

import (
	"fmt"
	"github.com/bmuller/arrow"
	"github.com/gookit/color"
	"log"
	"os"
)

func infoFileLogger(msg string, data interface{}){
	file:= getLogFile()
	infoLog := log.New(file,"INFO:\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog.Println(msg,data)
	file.Close()
}

func errorFileLogger(msg string, data interface{}){
	file:= getLogFile()
	errorLog := log.New(file,"ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	errorLog.Println(msg,data)
	file.Close()
}

func warningFileLogger(msg string, data interface{}){
	file:= getLogFile()
	warningLog := log.New(file,"WARNING:\t", log.Ldate|log.Ltime|log.Lshortfile)
	warningLog.Println(msg,data)
	file.Close()
}

func debugFileLogger(msg string, data interface{}){
	file:= getLogFile()
	debugLog := log.New(file,"DEBUG:\t", log.Ldate|log.Ltime|log.Lshortfile)
	debugLog.Println(msg,data)
	file.Close()
}


func infoLogger(msg string, data interface{}) {
	color.Info.Println( "INFO: ", arrow.Now().CFormat("%H:%M:%S"), msg, data)
}

func debugLogger(msg string, data interface{}) {
	color.Debug.Println( "DEBUG: ", arrow.Now().CFormat("%H:%M:%S"), msg, data)
}

func errorLogger(msg string, data interface{}) {
	color.Error.Println( "ERROR: ", arrow.Now().CFormat("%H:%M:%S"), msg, data)
}

func warningLogger(msg string, data interface{}) {
	color.Warn.Println( "WARNING: ", arrow.Now().CFormat("%H:%M:%S"), msg, data)
}

func getLogFile() *os.File{

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs",0765)
	}

	fileName := fmt.Sprintf("logs/"+arrow.Now().CFormat("02_Jan_2006")+".txt")
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0765)
	if err!= nil{
		return nil
	}
	return file
}
