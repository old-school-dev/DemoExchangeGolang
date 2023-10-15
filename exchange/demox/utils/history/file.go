package history

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func LoadHistory() []PriceHistory {
	file, err := os.Open("pricehistory.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	var prices []PriceHistory
	for _, his := range records {
		open, _ := strconv.ParseFloat(his[2],64)
		high, _ := strconv.ParseFloat(his[3],64)
		low, _ := strconv.ParseFloat(his[4],64)
		close, _ := strconv.ParseFloat(his[5],64)
		prices = append(prices, PriceHistory{
			Date:  his[0],
			Time:  his[1],
			Open:  open,
			High:  high,
			Low:   low,
			Close: close,
		})
	}
	return prices
}
