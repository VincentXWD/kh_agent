package main

import (
	"os"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"strings"
)

type New struct {
	Prefix string
	NewId string
	Title string
	Time string
	Content string
	Subject string
}

type Subject struct {
	Name string
	Url string
}

func CreateDir(PathName string) error {
	err := os.Mkdir(PathName, 0777)
	if err != nil && !os.IsExist(err) {
		return err
	}
	return nil
}

func AppendFile(SavePath string, FileName string, buf string) {
	out, err := os.OpenFile(SavePath+FileName, os.O_WRONLY, 0644)
	defer out.Close()
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	offset, err := out.Seek(0, os.SEEK_END)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	_, err = out.WriteAt([]byte(buf), offset)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	log.Warnln("Save file finished. Locate in ", SavePath + FileName)
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func SaveFile(SavePath string, FileName string, buf string) {
	out, err := os.Create(SavePath + FileName)
	defer out.Close()
	fmt.Fprintf(out, "%s", buf)
	if err != nil {
		log.Errorln(err.Error())
		return
	}
	log.Warnln("Save file finished. Locate in ", SavePath + FileName)
}

func ReadAll(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func ReadFile(path string) []string {
	var fp interface{}
	fp, err := ReadAll(path)
	if err != nil {
		log.Errorln(err.Error())
		return nil
	}
	fp = string(fp.([]byte))
	return strings.Split(fp.(string), "\n")
}