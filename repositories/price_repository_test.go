package repositories

import (
	"testing"
	"gopkg.in/mgo.v2"
	"microgo/models"
	"github.com/stretchr/testify/assert"
)

const DBNAME = "new_test-db"

func TestPriceRepository(t *testing.T) {
	session := setup()
	repository := NewPriceRepository(session, DBNAME)

	expectedPrice := models.HotelPrice{
		RoomName: "одноместный",
		Price:    1.212,
		Tariff:   "new2017",
	}

	id, _ := repository.Save(&expectedPrice)
	assert.NotNil(t, id)

	actualPrice, _ := repository.FindById(id)

	assert.Equal(t, expectedPrice.RoomName, actualPrice.RoomName)
	assert.Equal(t, expectedPrice.Price, actualPrice.Price)
	assert.Equal(t, expectedPrice.Tariff, actualPrice.Tariff)

	cleanup(session)
}

func setup() *mgo.Session {
	session, _ := mgo.Dial("localhost")
	session.SetMode(mgo.Monotonic, true)
	session.DB(DBNAME).DropDatabase()
	return session
}

func cleanup(session *mgo.Session)  {
	session.DB(DBNAME).DropDatabase()
	session.Close()
}
