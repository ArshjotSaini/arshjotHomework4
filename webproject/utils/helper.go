package utils

import (
	"os"

	"github.com/xuri/excelize/v2"
)

// method to open and load data from xlsx file
func Load_xlsx() [][]string {
	file, err := excelize.OpenFile("./assets/Project3Data.xlsx")
	ThrowError(err)

	rowss, err := file.GetRows("Comp490 Jobs")
	ThrowError(err)
	return rowss
}

// check if file exist or not
func IsFileExist(filename string) bool {
	res, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !res.IsDir()
}
