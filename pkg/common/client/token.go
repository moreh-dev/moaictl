package client

import (
	"errors"
	"log"
	. "moaictl/pkg/common/config"
	"os"
)

// TODO: This code came from moai-smi, need to refactor into one pkg

const (
	TokenType string = "Bearer"
)

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func openFile(filePath string, truncate bool) (*os.File, error) {
	flags := os.O_RDWR | os.O_CREATE
	if truncate {
		flags = flags | os.O_TRUNC
	}
	file, err := os.OpenFile(filePath, flags, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func SaveToken(token string) error {
	f, err := openFile(Config.TokenPath, true)
	defer f.Close()
	checkError(err)

	_, err = f.WriteString(token)
	checkError(err)
	return nil
}

func GetToken() string {
	buf := make([]byte, 8192) // default http header size
	f, err := openFile(Config.TokenPath, false)
	checkError(err)

	cnt, err := f.Read(buf)
	if cnt == 0 || err != nil {
		var str = "no such token"
		err := errors.New(str)
		log.Fatal(str)
		panic(err)
	}

	return TokenType + " " + string(buf[:cnt-1])

}
