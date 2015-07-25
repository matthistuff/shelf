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