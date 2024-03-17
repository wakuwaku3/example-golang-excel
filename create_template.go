package main

import "github.com/xuri/excelize/v2"

func createTemplate() (errResult error) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			if errResult == nil {
				errResult = err
			}
		}
	}()

	errResult = func() error {
		if err := f.SetCellValue("Sheet1", "A1", "Name"); err != nil {
			return err
		}
		if err := f.SetCellValue("Sheet1", "B1", "Age"); err != nil {
			return err
		}

		if err := f.SetSheetName("Sheet1", "users"); err != nil {
			return err
		}

		return f.SaveAs("template.xlsx")
	}()
	return
}
