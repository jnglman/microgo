package container

import (
	"microgo/handlers"
	"gopkg.in/mgo.v2"
	"log"
	"microgo/repositories"
	"sync"
)

const DBNAME = "priceDB"

type IDependencyContainer interface {
	InjectHandler() handlers.PriceHandler
}

type kernel struct{}

func (k *kernel) InjectHandler() handlers.PriceHandler {
	session := ConnectDB()
	priceDao := repositories.NewPriceRepository(session, DBNAME)
	priceHandler := handlers.NewPriceHandler(priceDao)
	return priceHandler
}

var (
	k             *kernel
	containerOnce sync.Once
	mgoSession    *mgo.Session
)

func DependencyContainer() IDependencyContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}

func ConnectDB() *mgo.Session {
	var err error
	mgoSession, err = mgo.Dial("localhost")
	mgoSession.SetMode(mgo.Monotonic, true)

	if err != nil {
		log.Fatal("Failed to start the Mongo session.")
	}
	return mgoSession
}
