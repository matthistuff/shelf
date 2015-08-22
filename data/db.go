package data

import (
	"gopkg.in/mgo.v2"
	"os"
)

var session *mgo.Session

func DB() (*mgo.Database, *mgo.Session) {
	if session == nil {
		s, err := mgo.Dial(os.Getenv("SHELF_DB_HOST"))
		if err != nil {
			panic(err)
		}

		session = s
	}

	return session.DB("shelf"), session
}

func Objects() *mgo.Collection {
	db, _ := DB()

	return db.C("objects")
}

func Files() *mgo.GridFS {
	db, _ := DB()

	return db.GridFS("fs")
}
