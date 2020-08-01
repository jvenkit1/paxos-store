package db

import (
	"errors"
	"paxos_store/utils"
)

type operations interface {
	GetData()
	InsertData()
	DeleteData() error
}

var record map[string]string

type data struct {
	key string
	value string
}

func init() {
	record = make(map[string]string)  // check if required
}

// For a particular key, return the value
func (d *data) GetData(){
	val := record[d.key]
	d.value = val
}

// Insert the value corresponding to d.key into the data store
func (d *data) InsertData() {
	record[d.key] = d.value
}

// Remove the value corresponding to d.key from the data store
// Return error if value does not exist
func (d *data) DeleteData() error{
	_, exists := record[d.key]
	if exists != true {
		return errors.New(utils.RecordNotFound)
	}
	delete(record, d.key)
	return nil
}