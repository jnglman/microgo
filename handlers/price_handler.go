package handlers

import (
	"net/http"
	"microgo/models"
	"encoding/json"
	"microgo/repositories"
	"log"
	"github.com/gorilla/mux"
	"microgo/strategies"
	"errors"
)

type PriceHandler struct {
	priceRepository repositories.IPriceRepository
	strategyHolder  strategies.StrategyHolder
}

type IdView struct {
	ID string `json:"Id"`
}

func NewPriceHandler(repository repositories.IPriceRepository) PriceHandler {
	handler := PriceHandler{strategyHolder: strategies.StrategyHolder{}}
	handler.priceRepository = repository
	return handler
}

func (handler *PriceHandler) SavePrice(w http.ResponseWriter, r *http.Request) {
	var price models.HotelPrice

	if err := json.NewDecoder(r.Body).Decode(&price); err != nil {
		log.Println(err)
		http.Error(w, "Wrong request syntax", http.StatusBadRequest)
		return
	}

	if err := ResolveTariffStrategy(price.Tariff, &handler.strategyHolder); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		price.Price = handler.strategyHolder.CurrentStrategy.CalculatePrice(price.Price)
	}

	if newId, err := handler.priceRepository.Save(&price); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		jsonResponse(w, http.StatusCreated, IdView{newId})
	}
}

func (handler *PriceHandler) GetPrice(w http.ResponseWriter, r *http.Request) {
	requestId := mux.Vars(r)["id"]

	if price, err := handler.priceRepository.FindById(requestId); err != nil {
		log.Println(err)
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	} else {
		jsonResponse(w, http.StatusOK, &price)
	}
}

func ResolveTariffStrategy(tariff string, strategyHolder *strategies.StrategyHolder) error {
	switch tariff {
	case "old2017":
		strategyHolder.CurrentStrategy = strategies.NewMultiplierStrategy(1.13)
	case "new2017":
		strategyHolder.CurrentStrategy = strategies.NewMultiplierStrategy(1.20)
	case "2018":
		strategyHolder.CurrentStrategy = strategies.NewMultiplierStrategy(1.5)
	default:
		return errors.New("tariff doesn't exist")
	}
	return nil
}

func jsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
