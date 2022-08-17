package database

type Add struct {
	// optional field
	Data  interface{}
	Datas []interface{}
}

func (a *Add) getData() interface{} {
	return a.Data
}

func (a *Add) getDatas() []interface{} {
	return a.Datas
}
