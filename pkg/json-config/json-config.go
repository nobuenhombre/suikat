// Package jsonconfig provides function to load config files in json format
package jsonconfig

import (
	"encoding/json"

	"github.com/nobuenhombre/suikat/pkg/fico"
	"github.com/nobuenhombre/suikat/pkg/ge"
)

func Load(fileName string, jc interface{}) error {
	txtConfigFile := fico.TxtFile(fileName)

	data, err := txtConfigFile.Read()
	if err != nil {
		return ge.Pin(err)
	}

	err = json.Unmarshal([]byte(data), jc)
	if err != nil {
		return ge.Pin(err)
	}

	return nil
}
