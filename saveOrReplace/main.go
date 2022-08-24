package saveorreplace

import (
	"github.com/ignacioMagno/database"
)

type SaveOrReplace interface {
	GetDatabase() database.Db
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

	save := database.Save{
		Db: db,
		Add: database.Add{
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

	save := database.FindAndModify{
		Db:   db,
		Data: s.GetObject(),
		Bind: &i,
	}

	return save.FindOneAndReplace()
}
