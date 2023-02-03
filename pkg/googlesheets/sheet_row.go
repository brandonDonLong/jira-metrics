package googlesheets

import (
	"fmt"
	"reflect"
)

type MySheetRow struct {
	Sprint       string `json:"Sprint"`
	Dicipline    string `json:"Dicipline"`
	TicketNumber string `json:"Ticket Number"`
	Title        string `json:"Title"`
	Link         string `json:"Link"`
	Commited     int    `json:"Commited"`
	Dropped      int    `json:"Dropped"`
	Added        int    `json:"Added"`
	Adjusted     int    `json:"Adjusted"`
	CarriedOver  int    `json:"Carried Over"`
	Completed    int    `json:"Completed"`
}

type MySheetRowArray []MySheetRow

type GoogleSheetValues [][]interface{}

func (m MySheetRowArray) Convert() GoogleSheetValues {

	// for key, value := range m {
	// 	// fmt.Printf("\n brandon sValue: %d values: %s \n", key, value);
	// }

	// result := make(GoogleSheetValues, len(m))

	// i := 0
	for _, s := range m {

		// transforming the sheet struct in a generic interface array ([]interface{})
		sValue := reflect.ValueOf(s)
		values := make([]interface{}, sValue.NumField())
		// fmt.Printf("\n brandon sValue: %s values: %s \n", sValue, values);
		for j := 0; j < sValue.NumField(); j++ {
			// copy struct field value into interface
			if (j == sValue.NumField()-1){
				fmt.Printf("%v\n", sValue.Field((j)))
			} else {
				fmt.Printf("%v|", sValue.Field((j)))
			}

			if sValue.Field(j).CanInterface() {
				values[j] = sValue.Field(j).Interface()
			}
		}

		// // result[i] = values
		// i++
	}

	return nil
}
