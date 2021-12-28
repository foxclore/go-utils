package config

import (
	"encoding/json"
	"io/ioutil"
)

func ReadConfig(parseTo interface{}, names ...string) error {
	var name string
	if len(names) == 0 {
		name = "config.json"
	} else {
		name = names[0]
	}

	d, err := ioutil.ReadFile(name)

	if err != nil {
		return err
	}

	return json.Unmarshal(d, parseTo)
}
