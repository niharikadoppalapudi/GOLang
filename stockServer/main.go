package main

import (
	//"fmt"
	"encoding/json"
	"net/http"
	//	"os"
	"github.com/gorilla/mux"
)

type Stock_ex struct {
	StockName string
}
type Stock struct {
	Symbol          string  `json:"symbol,omitempty"`
	Name            string  `json:"name,omitempty"`
	Price           float32 `json:"price,omitempty"`
	Close_yesterday float32 `json:"close_yesterday,omitempty"`
	Currency        string  `json:"currency"`
	Market_cap      int64   `json:"market_cap"`
	Volume          int     `json:"volume"`
	Timezone        string  `json:"timezone"`
	Timezone_name   string  `json:"timezone_name"`
	Gmt_offset      int     `json:"gmt_offset"`
	Last_trade_time string  `json:"last_trade_time"`
	Stock_ex        string
}

var de []Stock

func GetDefaultStockDetails(res http.ResponseWriter, req *http.Request) {
	//fmt.Println("test")

	json.NewEncoder(res).Encode(de)
}

var s []Stock

func GetStockDetails(res http.ResponseWriter, req *http.Request) {

	params := mux.Vars(req)
	for _, sym := range s {
		if sym.Symbol == params["symbol"] {
			json.NewEncoder(res).Encode(sym)
			return
		}
	}
	json.NewEncoder(res).Encode(s)
}

func main() {
	router := mux.NewRouter()
	de = append(de, Stock{Name: "Test Inc.",
		Price:           134.04,
		Close_yesterday: 134.04})
	s = append(s, Stock{Symbol: "AAPL", Name: "Apple Inc.", Price: 154.94, Close_yesterday: 154.94,
		Currency: "USD", Market_cap: 732835688367, Volume: 142022, Timezone: "EST", Gmt_offset: -18000, Last_trade_time: "2019-01-16 16:00:01"})
	s = append(s, Stock{Symbol: "SAMI", Name: "Samsung Inc.", Price: 194.90, Close_yesterday: 194.90})
	router.HandleFunc("/stock", GetDefaultStockDetails).Methods("GET")
	router.HandleFunc("/stock/{symbol}", GetStockDetails).Methods("GET")
	http.ListenAndServe(":8080", router)

}
