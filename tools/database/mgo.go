package database

import (
	"time"

	"github.com/globalsign/mgo"
	"githup.com/bitini111/leafcom/log"
)

func NewMongoSession(host string) *mgo.Session {
	session, err := mgo.DialWithTimeout(host, 3*time.Second)
	if err != nil {
		log.Fatal("fail to init mongo connection")
	}
	session.SetPoolLimit(300)
	return session
}
