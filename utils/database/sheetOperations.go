package database

import (
	"google.golang.org/api/sheets/v4"
	"log"
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
func (service SheetService) PutData(spreadsheetId string, writeRange string, valuesToWrite []interface{}) (AppendValuesResponse, error) {
	var valueRange sheets.ValueRange
	valueRange.Values = append(valueRange.Values, valuesToWrite)
	response, err := service.Service.Spreadsheets.Values.Append(spreadsheetId, writeRange, &valueRange).ValueInputOption("RAW").Do()
	if err != nil {
		log.Fatalf("Error occurred while writing data into the spreadsheet: %v", err)
	}
	return AppendValuesResponse{AppendValuesResponse: response}, err
}

// PutDataBatch puts data into the given Google Sheet in bulk.
func (service SheetService) PutDataBatch(spreadsheetId string, writeRange string, valuesToWrite [][]interface{}) {

}
