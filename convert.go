package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Product struct {
	ImageUrl  string `json:"image_url"`
	MobileUrl string `json:"mobile_url"`
	Price     int    `json:"price"`
	ShopEn    string `json:"shop_en"`
	ShopKo    string `json:"shop_ko"`
	Title     string `json:"title"`
	Url       string `json:"url"`
}

type Ranking struct {
	Description string `json:"description"`
	NameEn string `json:"name_en"`
	NameKo string `json:"name_ko"`
	ImageUrl string `json:"image_url"`
	Url string `json:"url"`
}

func Convert() {
	raw, err := ioutil.ReadFile("./product.json")
	if err != nil {
		fmt.Println(err)
	}

	var jsonData []Product
	err = json.Unmarshal([]byte(raw), &jsonData)
	if err != nil {
		fmt.Println(err)
	}

	csvData, err := os.Create("./product.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvData.Close()

	writer := csv.NewWriter(csvData)

	for _, worker := range jsonData {
		var record []string
		record = append(record, worker.ImageUrl)
		record = append(record, worker.MobileUrl)
		record = append(record, strconv.Itoa(worker.Price))
		record = append(record, worker.ShopEn)
		record = append(record, worker.ShopKo)
		record = append(record, worker.Title)
		record = append(record, worker.Url)
		writer.Write(record)
	}
	writer.Flush()
}

func ConvertRankings() {
	raw, err := ioutil.ReadFile("./ranking.json")
	if err != nil {
		fmt.Println(err)
	}

	var jsonData []Ranking
	err = json.Unmarshal([]byte(raw), &jsonData)
	if err != nil {
		fmt.Println(err)
	}

	csvData, err := os.Create("./ranking.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer csvData.Close()

	writer := csv.NewWriter(csvData)

	for _, worker := range jsonData {
		var record []string
		record = append(record, worker.Description)
		record = append(record, worker.NameEn)
		record = append(record, worker.NameKo)
		record = append(record, worker.ImageUrl)
		record = append(record, worker.Url)
		writer.Write(record)
	}
	writer.Flush()
}