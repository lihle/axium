package storage

import (
	"axium/types"
)

//GetAllPersons : returns all persons
func GetAllPersons() (persons []types.Person, err error) {
	raws, err := db.Query("SELECT * FROM persons")
	if err != nil {
		return
	}
	defer raws.Close()

	for raws.Next() {
		var p types.Person
		err = raws.Scan(&p.PersonID, &p.Firstname, &p.Lastname, &p.Birthdate, &p.Gender)
		if err != nil {
			return
		}
		persons = append(persons, p)
	}
	return
}
