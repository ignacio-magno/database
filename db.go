package database

type Db struct {
	NameCollection string
	Database       string
	Filter         interface{}
}

func (db *Db) getNameCollection() string {
	return db.NameCollection
}

func (db *Db) getDatabase() string {
	return db.Database
}

func (db *Db) getFilter() interface{} {
	return db.Filter
}
