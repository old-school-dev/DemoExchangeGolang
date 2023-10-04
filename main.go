package main

import (
	"demox/exchange"
	"fmt"
	"log"
)

func main() {
	log.Println("hello world")
	ex := exchange.NewExchange("Famp","Visarut")
	order := ex.CreateBuyOrder("BTC", 999)
	log.Println(ex.GetOrderDetail(order.Id))
	for i:=0; i < 10; i ++ {
		log.Println(ex.GetPrice("BTC"))
	}
	log.Println(ex.GetPairs())
	fmt.Scanf("h")
}