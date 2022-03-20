package database

import (
	"context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
	"log"
)

const (
	clientSecretPath = "secret.json"
)

type SheetService struct {
	*sheets.Service
}

// NewSpreadsheetService creates a new, authenticated spreadsheet service.
func NewSpreadsheetService() (SheetService, error) {
	context := context.Background()
	service, error := sheets.NewService(context, option.WithCredentialsFile(clientSecretPath), option.WithScopes(sheets.SpreadsheetsScope))
	if error != nil {
		log.Fatalf("Couldn't initiate service - %v", error)
	}
	return SheetService{service}, error
}
