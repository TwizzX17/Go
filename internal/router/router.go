package router

import (
	"net/http"
	"stock_data/internal/handler"
	"stock_data/internal/middleware"
)

func SymbolsRouter() http.Handler {
	mux := http.NewServeMux()

	// Symbol routes
	mux.HandleFunc("/api/symbols", handler.GetSymbolsHandler)

	// Scheduled sync
	mux.HandleFunc("/api/scheduledSync", handler.FetchSymbolsHandler)

	return middleware.JSONMiddleware(mux)
}
