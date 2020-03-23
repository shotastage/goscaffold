package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Init() {

	buf, err := ioutil.ReadFile("Scaffoldfile")

	if err != nil {
		return
	}

	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}
}
