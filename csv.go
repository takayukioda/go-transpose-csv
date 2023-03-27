package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 入力ファイルを開く
	inputfile, err := os.Open("input.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer inputfile.Close()

	// CSVリーダーを作成
	reader := csv.NewReader(inputfile)

	// 出力ファイルを作成
	outputfile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outputfile.Close()

	// CSVライターを作成
	writer := csv.NewWriter(outputfile)
	defer writer.Flush()

	order := []string{
		"id",
		"content",
	}

	// 列のマッピング
	columns := map[string]int{
		"id":      0,
		"content": 1,
	}

	var rows [][]string
	for {
		// CSVファイルから1行読み込む
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for index, title := range order {
			titleAcc := fmt.Sprintf("精度:%s", title)
			if indexAcc, ok := columns[titleAcc]; ok {
				fmt.Println(index, ":", title, ":", columns[title])
				rows = append(rows, []string{title, record[columns[title]], record[indexAcc]})
			} else if strings.HasPrefix(title, "精度:") {
				continue
			} else {
				fmt.Println(index, ":", title, ":", columns[title])
				rows = append(rows, []string{title, record[columns[title]]})
			}
		}
	}

	// 列の値を転置してCSVファイルに書き込む
	for _, row := range rows {
		writer.Write(row)
	}
}
