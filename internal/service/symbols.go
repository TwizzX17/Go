package service

import (
	"fmt"
	"io"
	"os"
	"stock_data/internal/model"
	"stock_data/utils"
	"strconv"
	"time"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var apiKey string = os.Getenv("API_KEY")
var baseUri string = os.Getenv("BASE_URI")

func GetSymbols() ([]model.Symbol, error) {
	return model.QuerySymbols()
}

/*
Run to fetch and save latest data from the remote endpoint
*/
func SyncSymbols() error {
	symbols, err := fetchSymbols()
	if err != nil {
		return fmt.Errorf("error fetching symbols: %v", err)
	}

	err = model.SaveSymbols(symbols)
	if err != nil {
		return fmt.Errorf("error saving symbols: %v", err)
	}

	return nil
}

/*
Execute get request, parse and return symbols from csv
Docs: https://www.alphavantage.co/documentation/
*/
func fetchSymbols() ([]model.Symbol, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API_KEY is not set")
	}

	url := fmt.Sprintf("%s/query?function=LISTING_STATUS&apikey=%s", baseUri, apiKey)

	csvReader, err := utils.GetCsvFromRemote(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching symbols: %v", err)
	}

	symbols := []model.Symbol{}
	/* Loop over csv rows */
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading CSV record: %v", err)
		}

		id, _ := strconv.Atoi(record[0])
		createdTs, _ := time.Parse("2006-01-02", record[4])
		status := record[6] == "Active"

		symbol := model.Symbol{
			Id:        id,
			Name:      record[1],
			Symbol:    record[0],
			CreatedTs: createdTs,
			Exchange:  record[2],
			AssetType: record[3],
			Status:    status,
			DeletedTs: nil,
		}
		symbols = append(symbols, symbol)
	}

	return symbols, nil
}
