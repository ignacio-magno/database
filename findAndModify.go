package database

import "errors"

type FindAndModify struct {
	Db
	Data interface{}
	Bind interface{}
}

func (f *FindAndModify) getData() interface{} {
	return f.Data
}

func (f *FindAndModify) getBind() interface{} {
	return f.Bind
}

func (f *FindAndModify) getFilter() interface{} {
	return f.Filter
}

// replace all document
func (db *FindAndModify) FindOneAndReplace() error {
	if db.getData() != nil {
		res := client.Database(database).Collection(db.getNameCollection()).FindOneAndReplace(ctx, db.getFilter(), db.getData())
		if res.Err() != nil {
			return res.Err()
		}

		return res.Decode(db.getBind())
	} else {
		return errors.New("data is nil")
	}
}

// need update operators
func (db *FindAndModify) FindOneAndUpdate() error {
	res := client.Database(database).Collection(db.getNameCollection()).FindOneAndUpdate(ctx, db.getFilter(), db.getData())
	if res.Err() != nil {
		return res.Err()
	}
	return res.Decode(db.getBind())
}

func (db *FindAndModify) FindOneAndDelete() error {
	res := client.Database(database).Collection(db.getNameCollection()).FindOneAndDelete(ctx, db.getFilter())
	if res.Err() != nil {
		return res.Err()
	}
	return res.Decode(db.getBind())
}
