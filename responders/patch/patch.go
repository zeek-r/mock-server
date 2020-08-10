package service

import (
	"io/ioutil"
)

func Put(filePath string) ([]byte, error) {
	return ioutil.ReadFile(filePath)
}
