package main

import (
	"os"
	"errors"
	"encoding/json"
)

var ErrInexistantBC =errors.New("No existing Blockchain found.  Create one first.")

// true if the file is existant
func bcFileExists(file string) bool {
	//get the stat of a file using it's name
	_, err := os.Stat(file)
	//if there is an error about it not existing the it doesn't exist
	return !os.IsNotExist(err)
}

func LoadBlockchain(file string) (*Blockchain,error){
	var bc Blockchain
	//check if the file exist
	if !bcFileExists(file) {
		return nil, ErrInexistantBC
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &bc)
	if err != nil {
		return nil, err
	}

	return &bc, nil

}

func SaveBlockchain(bc *Blockchain, file string) error {
	jsonData, err := json.Marshal(bc)
	if err != nil {
        return err
    }
	err = os.WriteFile(file, jsonData, 0666)
	return err
}
