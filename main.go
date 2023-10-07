package main

import (
	"demo-strategy/exchange"
	_ "fmt"
	"log"
)

func main() {
	// log.Println("hello world")
	// ex := exchange.NewExchange("","")
	// order := ex.CreateBuyOrder("BTC", 999)
	// log.Println(ex.GetOrderDetail(order.Id))
	// for i:=0; i < 10; i ++ {
	// 	log.Println(ex.GetPrice("BTC"))
	// }
	// fmt.Scanf("h")
	TestConnection()
}
func TestConnection() {
	ex := exchange.NewExchange()
	api1 := exchange.NewAPI("Famp", "Visarut", ex)
	// api2 := exchange.NewAPI("FFFFF", "AAAAA", ex)
	log.Println(api1.GetPrice("BTC/THB"))
	api1.CreateBuyOrder("BTC/THB", 999)
	log.Println(api1.GetOrders())
	log.Println(ex)
}