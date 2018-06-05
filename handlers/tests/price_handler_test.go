package tests

import (
	"testing"
	"microgo/mocks"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"bytes"
	"microgo/handlers"
	"github.com/gorilla/mux"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"microgo/models"
)

func TestPriceHandler_SaveTest(t *testing.T) {
	repository := new(mocks.IPriceRepository)
	id := bson.NewObjectId()
	hex := id.Hex()
	repository.On("Save", mock.Anything).Return(hex, nil)

	payload := []byte(`{ "RoomName": "одноместный", "Price" :1.01, "Tarif": "new2017" }`)
	req := httptest.NewRequest("POST", "/", bytes.NewBuffer(payload))
	response := httptest.NewRecorder()

	handler := handlers.NewPriceHandler(repository)
	router := mux.NewRouter()
	router.HandleFunc("/", handler.SavePrice)
	router.ServeHTTP(response, req)

	var actualResult handlers.IdView
	json.NewDecoder(response.Body).Decode(&actualResult)

	assert.Equal(t, hex, actualResult.ID)
	assert.Equal(t, 201, response.Code)
}

func TestPriceHandler_GetTest(t *testing.T)  {
	repository := new(mocks.IPriceRepository)
	hexId := "4d88e15b60f486e428412dc9"
	expectedPrice := models.HotelPrice{
		Id:       bson.ObjectIdHex(hexId),
		RoomName: "одноместный",
		Price:    1.212,
		Tariff:   "new2017",
	}
	repository.On("FindById", mock.Anything).Return(&expectedPrice, nil)

	request := httptest.NewRequest("GET", "/"+hexId, nil)
	response := httptest.NewRecorder()

	handler := handlers.NewPriceHandler(repository)
	router := mux.NewRouter()
	router.HandleFunc("/{id}", handler.GetPrice)
	router.ServeHTTP(response, request)

	var actualResult models.HotelPrice
	json.NewDecoder(response.Body).Decode(&actualResult)

	assert.Equal(t, expectedPrice, actualResult)
}
