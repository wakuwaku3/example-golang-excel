package main

import (
	_ "embed"

	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/xuri/excelize/v2"
)

//go:embed template.xlsx
var template []byte

type user struct {
	name string
	age  int
}

var users = []user{
	{"Tom", 18},
	{"Jerry", 20},
}

func add() (errResult error) {
	f, err := excelize.OpenReader(bytes.NewReader(template))
	if err != nil {
		errResult = err
		return
	}

	defer func() {
		if err := f.Close(); err != nil {
			if errResult == nil {
				errResult = err
			}
		}
	}()

	const startRow = 2
	errResult = func() error {
		for i, u := range users {
			nameAddress, err := excelize.JoinCellName("A", i+startRow)
			if err != nil {
				return err
			}
			if err := f.SetCellValue("users", nameAddress, u.name); err != nil {
				return err
			}

			ageAddress, err := excelize.JoinCellName("B", i+startRow)
			if err != nil {
				return err
			}
			if err := f.SetCellValue("users", ageAddress, u.age); err != nil {
				return err
			}
		}

		// out がなければ作成
		if _, err := os.Stat("out"); os.IsNotExist(err) {
			if err := os.Mkdir("out", 0755); err != nil {
				return err
			}
		}

		outputPath := fmt.Sprintf("out/%s.xlsx", time.Now().Format("2006-01-02-15-04-05"))

		// sample コードなのであえて byte 配列に変換してからファイルに書き出している。
		// 直接書き出す場合 SaveAs を使う。
		// f.SaveAs(outputPath)
		output := bytes.NewBuffer(nil)
		if err := f.Write(output); err != nil {
			return err
		}

		return os.WriteFile(outputPath, output.Bytes(), 0644)
	}()
	return
}
