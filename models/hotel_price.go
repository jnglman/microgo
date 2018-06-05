package models

import "gopkg.in/mgo.v2/bson"

type HotelPrice struct {
	Id       bson.ObjectId `json:"Id" bson:"_id,omitempty"`
	RoomName string        `json:"RoomName" bson:"room_name"`
	Price    float64       `json:"Price" bson:"price"`
	Tariff   string        `json:"Tarif" bson:"tariff"`
}
