package credentials

import (
	"encoding/base64"
	"os"

	"github.com/nobuenhombre/suikat/pkg/futi"
)

// у нас есть некий файл credentials.json (содержащий например токены авторизации)
// Цель: не хранить credentials файлы в docker
//--------------------------------------------
// файл кодируется в base64
// закодированное сообщение хранится в OS.ENV (в секретах GKE)
// Нам нужно присвоить этому credentials имя
// по этому имени прочитать OS.ENV
// декодировать содержимое
// 1) предоставить приложению оригинальный файл
// 2) предоставить приложению []byte содержимое оригинального файла
// 3) предоставить приложению string содержимое оригинального файла

type Data struct {
	Key string
}

type EnvCred interface {
	GetBytes() ([]byte, error)
	GetString() (string, error)
	GetFile() (string, error)
}

func New(key string) EnvCred {
	return &Data{
		Key: key,
	}
}

func (d *Data) readEnv() ([]byte, error) {
	base64Str := ""
	if value, exists := os.LookupEnv(d.Key); exists {
		base64Str = value
	}

	dataBytes, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return []byte{}, err
	}

	return dataBytes, nil
}

func (d *Data) GetBytes() ([]byte, error) {
	return d.readEnv()
}

func (d *Data) GetString() (string, error) {
	data, err := d.readEnv()
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (d *Data) GetFile() (string, error) {
	data, err := d.readEnv()
	if err != nil {
		return "", err
	}

	file, err := futi.CreateTempFile("", "*", &data)
	if err != nil {
		return "", err
	}

	return file, nil
}
