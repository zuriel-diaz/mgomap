package mgomap

import (
	"time"
	"gopkg.in/mgo.v2"
	"reflect"
)

type Mapper struct {
	// Database connection vars
	Hosts 			[]string
	TimeOut 		time.Duration
	DatabaseName 	string
	DatabaseUser 	string
	DatabasePass 	string
	
	// Pointer to the active session of MongoDB
	Session *mgo.Session
}

func (self *Mapper) Connect() error {
	info := &mgo.DialInfo{
		Addrs: self.Hosts,
		Timeout: self.TimeOut,
		Database: self.DatabaseName,
		Username: self.DatabaseUser,
		Password: self.DatabasePass,
	}
	
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		return err
	}
	
	session.SetMode(mgo.Monotonic, true)
	self.Session = session
	
	return nil
}

func newTransaction(mapper *Mapper) *Transaction {
	return &Transaction{
		Mapper: mapper,
	}
}

func getName(model interface{}) string {
	t := reflect.TypeOf(model)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	
	return t.Name()
}