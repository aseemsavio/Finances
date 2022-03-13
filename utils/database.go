package utils

import (
	"context"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
)

const (
	clientSecretPath = "secret.json"
)

type Database struct {
	service *sheets.Service
}

// DatabaseService creates a Database instance on which the various functionalities can be called.
func DatabaseService() (Database, error) {
	ctx := context.Background()
	service, error := sheets.NewService(ctx, option.WithCredentialsFile(clientSecretPath), option.WithScopes(sheets.SpreadsheetsScope))
	if error != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", error)
	}
	return Database{service: service}, nil
}

func (db *Database) GetSpreadSheet(sheetId string) *sheets.SpreadsheetsGetCall {
	return db.service.Spreadsheets.Get(sheetId)
}

func SpreadSheet() sheets.Spreadsheet {
	return sheets.Spreadsheet{
		DataSourceSchedules: nil,
		DataSources:         nil,
		DeveloperMetadata:   nil,
		NamedRanges:         nil,
		Properties:          nil,
		Sheets:              nil,
		SpreadsheetId:       "",
		SpreadsheetUrl:      "",
		ServerResponse:      googleapi.ServerResponse{},
		ForceSendFields:     nil,
		NullFields:          nil,
	}
}

// CreateDatabase creates a spreadsheet to be used as a database.
func (db *Database) createDatabase(spreadSheet *sheets.Spreadsheet) {
	db.service.Spreadsheets.Create(spreadSheet)
}

// AddData adds data to the database.
func (db *Database) AddData(request *PushRequest) {
	var valueRange sheets.ValueRange
	valueRange.Values = append(valueRange.Values, request.Values)
	response, error := db.service.Spreadsheets.Values.Append(request.SpreadsheetId, request.Range, &valueRange).ValueInputOption("RAW").Do()
	log.Printf("Spreadsheet Push: %v", response)
	if error != nil {
		log.Fatalf("Failed to push data to DB %v", error)
	}
}

type PushRequest struct {
	SpreadsheetId string
	Range         string
	Values        []interface{}
}
