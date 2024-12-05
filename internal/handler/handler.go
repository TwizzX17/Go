package handler

import (
	"encoding/json"
	"net/http"
	"stock_data/internal/service"
)

func GetSymbolsHandler(w http.ResponseWriter, r *http.Request) {
	if !isGetRequest(r) {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	symbols, err := service.GetSymbols()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(symbols)
}

func FetchSymbolsHandler(w http.ResponseWriter, r *http.Request) {
	err := service.SyncSymbols()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func isGetRequest(r *http.Request) bool {
	return r.Method == http.MethodGet
}
