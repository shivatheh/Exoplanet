package dao

type DataBase struct {
	ShortUrlDB       map[string]string
	DomainCountDB    map[string]int
	ShorttoOriginUrl map[string]string
}

// will create instance of DB
func initDB() (*DataBase, error) {
	db := new(DataBase)
	db.ShortUrlDB = make(map[string]string)
	db.DomainCountDB = make(map[string]int)
	db.ShorttoOriginUrl = make(map[string]string)
	return db, nil
}
