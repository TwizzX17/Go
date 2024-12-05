package model

import (
	"context"
	"database/sql"
	"fmt"
	"stock_data/db"
	dbSetup "stock_data/internal/db"
	"time"
)

type Symbol struct {
	Id        int
	CreatedTs time.Time
	Symbol    string
	Name      string
	Exchange  string
	AssetType string
	Status    bool
	DeletedTs *time.Time
}

/*
Query all symbols from the db
*/
func QuerySymbols() ([]Symbol, error) {

	queries := db.New(dbSetup.DBConn)

	dbSymbols, err := queries.ListSymbols(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error querying symbols: %v", err)
	}

	var symbols []Symbol
	for _, dbSymbol := range dbSymbols {
		/* Convert query records to struct symbol values */
		symbols = append(symbols, Symbol{
			Id:        int(dbSymbol.ID),
			CreatedTs: *dbSetup.ToTime(dbSymbol.CreatedTs),
			Name:      dbSetup.ToString(dbSymbol.Name),
			Symbol:    dbSetup.ToString(dbSymbol.Symbol),
			Exchange:  dbSetup.ToString(dbSymbol.Exchange),
			AssetType: dbSetup.ToString(dbSymbol.AssetType),
			Status:    dbSetup.ToBool(dbSymbol.Status),
			DeletedTs: nil,
		})
	}

	return symbols, nil

}

/*
Save array of symbols to the db
*/
func SaveSymbols(symbols []Symbol) error {

	queries := db.New(dbSetup.DBConn)

	for _, symbol := range symbols {
		var deletedTs sql.NullTime
		if symbol.DeletedTs != nil {
			deletedTs = sql.NullTime{Time: *symbol.DeletedTs, Valid: true}
		} else {
			deletedTs = sql.NullTime{Valid: false}
		}

		err := queries.InsertSymbol(context.Background(), db.InsertSymbolParams{
			CreatedTs: sql.NullTime{Time: symbol.CreatedTs, Valid: true},
			Symbol:    sql.NullString{String: symbol.Symbol, Valid: true},
			Name:      sql.NullString{String: symbol.Name, Valid: true},
			Exchange:  sql.NullString{String: symbol.Exchange, Valid: true},
			AssetType: sql.NullString{String: symbol.AssetType, Valid: true},
			Status:    sql.NullBool{Bool: symbol.Status, Valid: true},
			DeletedTs: deletedTs,
		})

		if err != nil {
			return fmt.Errorf("error inserting symbol: %v", err)
		}

		fmt.Println("Inserted symbol: ", symbol.Symbol)
	}

	return nil
}
