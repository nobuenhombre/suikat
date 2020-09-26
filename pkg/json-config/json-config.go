package jsonconfig

import (
	"encoding/json"

	"github.com/nobuenhombre/suikat/pkg/fico"
	"github.com/nobuenhombre/suikat/pkg/ge"
)

type LoadError struct {
	ge.IdentityPlaceError
}

func Load(fileName string, jc interface{}) error {
	txtConfigFile := fico.TxtFile(fileName)

	data, err := txtConfigFile.Read()
	if err != nil {
		return &LoadError{
			IdentityPlaceError: ge.IdentityPlaceError{
				Place:  "txtConfigFile.Read",
				Parent: err,
			},
		}
	}

	err = json.Unmarshal([]byte(data), jc)
	if err != nil {
		return &LoadError{
			IdentityPlaceError: ge.IdentityPlaceError{
				Place:  "json.Unmarshal",
				Parent: err,
			},
		}
	}

	return nil
}
