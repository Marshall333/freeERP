package config_util

import (
	//log "ad-service/alog"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"os"
)

//load config
func LoadConfig(config interface{}, configFile string) error {
	file, err := os.OpenFile(configFile, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Printf("open file error: %+v", err)
		//log.Error("open file error: ", err)
		return err
	}
	defer file.Close()

	str, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("read all error: %+v", err)
		//log.Error("read all error: ", err)
		return err
	}

	if err = json.Unmarshal(str, config); err != nil {
		fmt.Printf("Unmarshal error: %+v", err)
		//log.Error("Unmarshal error: ", err)
		return err
	}
	return nil
}

func ParsePbFromTextFile(filePath string, pb proto.Message) error {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("ReadFile error. error: %+v", err)
		//log.Error("ReadFile error. error: ", err)
		return err
	}
	if err := proto.UnmarshalText(string(fileBytes), pb); err != nil {
		fmt.Printf("UnmarshalText error. error: %+v", err)
		//log.Error("UnmarshalText error. error: ", err)
		return err
	}
	return nil
}
