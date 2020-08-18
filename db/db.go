package db

import (
	"errors"
	"paxos-store/utils"
)

type operations interface {
	GetData()
	InsertData()
	DeleteData() error
}

var record map[string]string

type Data struct {
	Key   string
	Value string
}

func init() {
	record = make(map[string]string)  // check if required
}

// For a particular Key, return the value
func (d *Data) GetData(){
	val := record[d.Key]
	d.Value = val
}

// Insert the value corresponding to d.Key into the Data store
func (d *Data) InsertData() {
	record[d.Key] = d.Value
}

// Remove the value corresponding to d.Key from the Data store
// Return error if value does not exist
func (d *Data) DeleteData() error{
	_, exists := record[d.Key]
	if exists != true {
		return errors.New(utils.RecordNotFound)
	}
	delete(record, d.Key)
	return nil
}