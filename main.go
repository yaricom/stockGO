package main

import (
	"fmt"
	"github.com/yaricom/stockGO/trade"
)

func main() {
	var money float32
	var nStocks, daysLeft int
	engine := trade.New()
	for daysLeft > 0 {
		fmt.Scanf("%f %d %d", &money, &nStocks, &daysLeft)
		inputs := readTradeInputs(nStocks)
		orders, err := engine.Trade(inputs, money, daysLeft)
		if err != nil {
			fmt.Println(err)
		} else {
			outputTradeOrders(orders)
		}
	}
}

func outputTradeOrders(orders []trade.TradeOrder) {
	fmt.Println(len(orders))
	for _, order := range orders {
		if order.Operation == trade.BUY {
			fmt.Printf("%s %s %d", order.StockName, "BUY", order.Amount)
		} else {
			fmt.Printf("%s %s %d", order.StockName, "SELL", order.Amount)
		}
	}
}

func readTradeInputs(nStocks int) []trade.TradeInput {
	inputs := make([]trade.TradeInput, nStocks)
	for i := 0; i < nStocks; i++ {
		hist := make([]float32, 5)
		input := trade.TradeInput{}
		fmt.Scanf("%s %d %.2f %.2f %.2f %.2f %.2f", &input.StockName, &input.Owned,
			&hist[0], &hist[1], &hist[2], &hist[3], &hist[4])
		input.History = hist

		inputs[i] = input
	}
	return inputs
}


