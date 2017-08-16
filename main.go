package main

import (
	"fmt"
	"math/rand"
	"github.com/yaricom/stockGO/trade"
)

func main() {
	rand.Seed(157)
	var money float64
	var nStocks, daysLeft int
	engine := trade.New()
	fmt.Scanf("%f %d %d", &money, &nStocks, &daysLeft)
	//fmt.Printf("%f %d %d\n", money, nStocks, daysLeft)
	inputs := readTradeInputs(nStocks)
	//fmt.Println(len(inputs))
	orders, err := engine.Trade(inputs, money, daysLeft)
	//fmt.Println(len(orders))
	if err != nil {
		fmt.Println(err)
	} else {
		outputTradeOrders(orders)
	}
}

func outputTradeOrders(orders []trade.TradeOrder) {
	fmt.Println(len(orders))
	for _, order := range orders {
		if order.Operation == trade.BUY {
			fmt.Printf("%s %s %d\n", order.StockName, "BUY", order.Amount)
		} else {
			fmt.Printf("%s %s %d\n", order.StockName, "SELL", order.Amount)
		}
	}
}

func readTradeInputs(nStocks int) []trade.TradeInput {
	inputs := make([]trade.TradeInput, nStocks)
	for i := 0; i < nStocks; i++ {
		hist := make([]float64, 5)
		input := trade.TradeInput{}
		fmt.Scanf("%s %d %f %f %f %f %f", &input.StockName, &input.Owned,
			&hist[0], &hist[1], &hist[2], &hist[3], &hist[4])
		input.History = hist

		inputs[i] = input
	}
	return inputs
}


