package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func transpose(reader *csv.Reader, writer *csv.Writer) {
	order := []string{
		"id",
		"content",
	}

	// 列のマッピング
	columns := map[string]int{
		"id":      0,
		"content": 1,
	}

	var count int
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
		if count == 0 {
			count += 1
			continue
		}

		for index, title := range order {
			fmt.Println(index, ":", title, ":", columns[title])
			writer.Write([]string{title, record[columns[title]]})
		}

		count += 1
	}
}

func aggregate(reader *csv.Reader, writer *csv.Writer) {
	var count int
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
		if count == 0 {
			count += 1
			continue
		}

		fmt.Println(record)
		count += 1
	}
}

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

	transpose(reader, writer)

}
