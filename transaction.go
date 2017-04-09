package mgomap

import (
	"errors"
)

type Transaction struct {
	Mapper *Mapper
}

func (self *Transaction) Save(model interface{}) error {
	ref, ok := model.(DocumentActions)
	if !ok {
		return errors.New(DocumentFormatError)
	}
	
	if self.Mapper == nil {
		return errors.New(MapperError)
	}
	
	if self.Mapper.Session == nil {
		return errors.New(SessionError)
	}
	
	session := self.Mapper.Session.Copy()
	defer session.Close()
	
	colName := ref.GetDocumentName()
	if colName == "" {
		colName = getName(model)
	}
	
	// Call before actions
	ref.Before()
	
	// Verify the documnet has id, if has id return
	if ref.GetId().Hex() != "" {
		return errors.New(DocumentExisting)
	}
	
	// Save the document
	ref.CreateId()
	col := session.DB(self.Mapper.DatabaseName).C(colName)
	err := col.Insert(ref)
	if err != nil {
		return err
	}
	
	// Call after actions
	ref.After()
	
	return nil
}