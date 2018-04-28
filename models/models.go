package models

// Models returns one of every model the database must create
func Models() []interface{} {
	return []interface{}{
		&String{},
		&Item{},
		&Field{},
		&ItemType{},
		&User{},
		&Role{},
		&Project{},
		&Account{},
	}
}
