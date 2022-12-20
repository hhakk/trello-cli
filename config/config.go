package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Token        string `json:"token"`
	Key          string `json:"key"`
	DefaultBoard string `json:"default_board"`
}

func (t *Config) Input() {
	var token, key, board string
	fmt.Println("Let's setup required information.")
	fmt.Printf("Enter API key: ")
	fmt.Scanln(&key)
	fmt.Printf("Enter API token: ")
	fmt.Scanln(&token)
	fmt.Printf("Enter default board URL: ")
	fmt.Scanln(&board)
	t.Token = token
	t.Key = key
	t.DefaultBoard = board
}

func (t *Config) Save(filename string) error {
	bytes, err := json.Marshal(*t)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, bytes, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (t *Config) Load(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, t)
	if err != nil {
		return err
	}
	return nil
}
