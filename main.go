package main

import (
	"demo-strategy/constant"
	"demo-strategy/exchange"
	_"demo-strategy/exchange/demox/utils/history"
	_ "fmt"
	"log"
)

func main() {
	// history.LoadHistory()
	TestConnection()
}
func TestConnection() {
	ex := exchange.NewExchange()
	api1 := exchange.NewAPI("Famp", "Visarut", ex)
	log.Println(api1.GetBalances())
	api1.Deposit("THB", 100000)
	// api2 := exchange.NewAPI("FFFFF", "AAAAA", ex)
	log.Println(api1.GetPrice("BTC/THB"))
	log.Println(api1.GetBalances())
	bid := api1.CreateBuyOrder("BTC/THB", 1700, 1).Id
	
	for i := 0; i<40; i++ {
		status := api1.GetOrderDetail(bid).Status
		if status == constant.STATUS_MATCH {
			api1.CreateSellOrder("BTC/THB", 1001, 1)
			break
		}
		api1.GetPrice("BTC/THB")
	}

	for i := 0; i<300000; i++ {
		log.Println("Price:",api1.GetPrice("BTC/THB"))
	}
	log.Println(api1.GetOpenOrders())
	log.Println(api1.GetBalances())}