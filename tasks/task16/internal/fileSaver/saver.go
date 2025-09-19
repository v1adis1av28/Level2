package filesaver

import (
	"fmt"
	"os"
)

var dataDir string = "./data"

func CreateLocalDir() (string, error) {
	_, err := os.Stat(dataDir)
	if err == nil { // папка существует
		return dataDir, nil
	} else if os.IsNotExist(err) { // папки не существует следует ее создать
		err := os.Mkdir(dataDir, 777)
		if err != nil {
			return "", fmt.Errorf("creating directory error %v:", err.Error())
		}
	} else {
		return "", fmt.Errorf("error on creating local directory %v : ", err.Error())
	}
	return dataDir, nil
}
