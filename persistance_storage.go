package main

import (
	"encoding/json"
	"os"
)

// this file should be in charge of the persistance storage functionality

func (todos *Todos) Save(filename string) error { // should be used each time a function add, toggle, delete are called
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("storage.json", data, 0644)
}

func (todos *Todos) Load() error { // should be used each time a function (list)
	data, err := os.ReadFile("storage.json")
	if err != nil {
		return err
	}
	return json.Unmarshal(data, todos)
}
