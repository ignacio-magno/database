package mongodb

type Find struct {
	Db
	Bind interface{}
}

func (db *Find) getBind() interface{} {
	return db.Bind
}

func (db *Find) Find() error {
	res, err := client.Database(database).Collection(db.getNameCollection()).Find(ctx, db.getFilter())
	if err != nil {
		return err
	}

	if res.Err() != nil {
		return res.Err()
	}

	return res.All(ctx, db.getBind())
}

func (db *Find) FindOne() error {
	res := client.Database(database).Collection(db.getNameCollection()).FindOne(ctx, db.getFilter())

	if res.Err() != nil {
		return res.Err()
	}

	return res.Decode(db.getBind())
}
