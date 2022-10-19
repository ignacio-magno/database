package saveOrReplace

import (
	"github.com/ignacioMagno/database/mongodb"
)

type SaveOrReplace interface {
	GetDatabase() mongodb.Db
	GetFilter() interface{}
	GetObject() interface{}
	IsNew() bool
}

func Save(s SaveOrReplace) error {
	if s.IsNew() {
		return saveNew(s)
	} else {
		return replace(s)
	}
}

func saveNew(s SaveOrReplace) error {
	db := s.GetDatabase()

	save := mongodb.Save{
		Db: db,
		Add: mongodb.Add{
			Data: s.GetObject(),
		},
	}

	_, err := save.SaveOne()
	return err

}

func replace(s SaveOrReplace) error {
	var i interface{}
	db := s.GetDatabase()

	db.Filter = s.GetFilter()

	save := mongodb.FindAndModify{
		Db:   db,
		Data: s.GetObject(),
		Bind: &i,
	}

	return save.FindOneAndReplace()
}
