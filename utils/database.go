package utils

import (
	"context"
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

func DatabaseService() (Database, error) {
	ctx := context.Background()
	service, error := sheets.NewService(ctx, option.WithCredentialsFile(clientSecretPath), option.WithScopes(sheets.SpreadsheetsScope))
	if error != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", error)
	}
	return Database{service: service}, nil
}

func (db *Database) Append() {

}

type PushRequest struct {
	SpreadsheetId string
	Range         string
	Values        []interface{}
}
