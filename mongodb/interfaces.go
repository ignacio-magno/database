package mongodb

type IDb interface {
	getNameCollection() string
	getDatabase() string
	getFilter() interface{}
}

type IAdd interface {
	getData() interface{}
	getDatas() []interface{}
}

type IFindAndModify interface {
	IDb
	IAdd
	getBind() interface{}
}

type IFind interface {
	IDb
	getBind() interface{}
}

type IDelete interface {
	IDb
}

type ISaveOrModify interface {
	IDb
	IAdd
}
