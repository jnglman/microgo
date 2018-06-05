package repositories

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"microgo/models"
	"errors"
	"log"
	"fmt"
)

const (
	TABLE = "prices"
)

type IPriceRepository interface {
	Save(price *models.HotelPrice) (string, error)
	FindById(id string) (*models.HotelPrice, error)
}

type PriceRepository struct {
	session *mgo.Session
	dbName string
}

func NewPriceRepository(session *mgo.Session, dbName string) IPriceRepository {
	repository := new(PriceRepository)
	repository.session = session
	repository.dbName = dbName
	return repository
}

func (repo *PriceRepository) Save(price *models.HotelPrice) (string, error) {
	session := repo.getMongoSession()
	defer session.Close()
	price.Id = bson.NewObjectId()
	err := repo.getCollection(session).Insert(&price)
	return price.Id.Hex(), err
}

func (repo *PriceRepository) FindById(id string) (*models.HotelPrice, error) {
	var price models.HotelPrice
	if !bson.IsObjectIdHex(id) {
		log.Println()
		return nil, errors.New(fmt.Sprintf("id=%s isn't valid ObjectIdHex", id))
	}
	session := repo.getMongoSession()
	defer session.Close()
	err := repo.getCollection(session).FindId(bson.ObjectIdHex(id)).One(&price)
	return &price, err
}

func (repo *PriceRepository) getMongoSession() *mgo.Session {
	if repo.session == nil {
		panic("session closed")
	}
	return repo.session.Clone()
}

func (repo *PriceRepository) getCollection(session *mgo.Session) *mgo.Collection {
	return session.DB(repo.dbName).C(TABLE)
}