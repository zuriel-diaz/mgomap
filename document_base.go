package mgomap

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type DocumentActions interface {
	Before()		// Before call
	After()			// After call
	
	CreateId()
	GetId() bson.ObjectId
	GetDocumentName() string
}

type DocumentBase struct {
	Id 			bson.ObjectId 	`json:"id,omitempty" bson:"_id,omitempty"`
	CreatedAt 	time.Time 		`json:"created_at" bson:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at" bson:"updated_at"`
}

func (self *DocumentBase) Before()  {

}

func (self *DocumentBase) After() {

}

func (self *DocumentBase) CreateId()  {
	self.Id = bson.NewObjectId()
}

func (self *DocumentBase) GetId() bson.ObjectId {
	return self.Id
}

func (self *DocumentBase) GetDocumentName() string {
	return ""
}