package database

import (
	"context"
	"google.golang.org/api/sheets/v4"
	"log"
	"sync"
)

type SpreadsheetResponse struct {
	*sheets.ValueRange
}

type AppendValuesResponse struct {
	*sheets.AppendValuesResponse
}

// GetData gets data from the given Google Sheet.
func (service SheetService) GetData(spreadsheetId string, readRange string) (SpreadsheetResponse, error) {
	log.Printf("🤞 Getting data from Google Sheets...")
	response, err := service.Service.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Error while reading data from spreadsheet: %v", err)
	}
	return SpreadsheetResponse{
		ValueRange: response,
	}, err
}

// PutData puts data into the given Google Sheet.
func (service SheetService) PutData(spreadsheetId string, writeRange string, valuesToWrite []interface{}, wg *sync.WaitGroup) error {

	log.Printf("🤞 Writing data to Google Sheets...")
	var valueRange sheets.ValueRange
	valueRange.Values = append(valueRange.Values, valuesToWrite)
	_, err := service.Service.Spreadsheets.Values.Append(spreadsheetId, writeRange, &valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Error occurred while writing data into the spreadsheet: %v", err)
	}
	defer wg.Done()
	return err
}

// PutDataBatch puts data into the given Google Sheet in bulk.
func (service SheetService) putDataBatch(spreadsheetId string, writeRange string, valuesToWrite [][]interface{}, wg *sync.WaitGroup) error {
	defer wg.Done()
	log.Printf("🤞 Writing a batch of data to Google Sheets...")
	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "RAW",
	}
	rb.Data = append(rb.Data, &sheets.ValueRange{Range: writeRange, Values: valuesToWrite})
	_, err := service.Service.Spreadsheets.Values.BatchUpdate(spreadsheetId, rb).Context(context.Background()).Do()
	if err != nil {
		log.Fatalf("Error occurred while BULK writing data into the spreadsheet: %v", err)
	}
	return err
}
