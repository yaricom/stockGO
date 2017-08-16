package trade

import (
	"errors"
	"math/rand"
	"github.com/yaricom/stockGO/trade/stock"
)

// The trade operations
const (
	SELL int = iota
	BUY
)
// The money spend factor
const spendFactor = 0.8

// The trade input for particular stock
type TradeInput struct {
	StockName string
	History float32
	Owned int
}

// The trade order to be placed
type TradeOrder struct {
	StockName string
	Amount int
	// The operation to perform SELL, BUY
	Operation int
}

// The trade engine
type TradeEngine struct {
	stocks map[string]stock.Stock
}

func New() TradeEngine {
	return TradeEngine{
		stocks: make(map[string]stock.Stock),
	}
}

func (t *TradeEngine) addStock(name string, priceHistory float32, owned int) error {
	if _, ok := t.stocks[name]; ok {
		return errors.New("Stock already registered: " + name)
	}
	s := stock.New(name)
	s.Update(priceHistory, owned)
	t.stocks[name] = s

	return nil
}

func (t *TradeEngine) Trade(trades []TradeInput, money float32, daysLeft int) ([]TradeOrder, error) {
	// add / update stocks
	for _,tr := range trades {
		if s, ok := t.stocks[tr.StockName]; ok {
			s.Update(tr.History, tr.Owned)
		} else {
			err := t.addStock(tr.StockName, tr.History, tr.Owned)
			if err != nil{
				return nil, err
			}
		}
	}

	// find stocks to sell / buy
	var toSell, toBuy []stock.Stock
	for _,tr := range trades {
		s := t.stocks[tr.StockName]
		if s.CanBeSold(tr.History) {
			toSell = append(toSell, s)
		}
		if s.CanBeBought(tr.History) {
			toBuy = append(toBuy, s)
		}
	}

	// create trade orders
	var orders []TradeOrder
	for _, s := range toSell {
		amount := int((s.Owned + 1) * (rand.Float32() + 1.0) / 2.0) // owned * [0.5, 1]
		order := TradeOrder {
			StockName: s.Name,
			Amount: amount,
			Operation: SELL,
		}
		orders = append(orders, order)
	}

	moneyToSpent := money * rand.Float32() * spendFactor
	if moneyToSpent == 0 { moneyToSpent = money * spendFactor }
	ordersByStock := make(map[string]TradeOrder)
	for moneyToSpent > 0 {
		permIndxs := rand.Perm(len(trades)) // permutation of stock indexes to give chance to every stock randomly
		buyFailed := 0
		for _, indx := range(permIndxs) {
			st, ok := t.stocks[trades[indx].StockName]
			if !ok {
				return nil, errors.New("BUY| Stock not found: " + trades[indx].StockName)
			}
			if st.CurrPrice < moneyToSpent {
				// buy
				st.BidPrice = st.CurrPrice
				moneyToSpent -= st.CurrPrice
				order, ok := ordersByStock[st.Name]
				if !ok {
					order = TradeOrder {
						StockName: st.Name,
						Amount: 1,
						Operation: BUY,
					}
					ordersByStock[st.Name] = order
				} else {
					order.Amount += 1
				}
			} else {
				buyFailed += 1
			}
		}

		if buyFailed == len(trades) {
			break // failed to buy any stock - not enough money left
		}
	}

	// copy SELL orders
	for _, order := range(ordersByStock) {
		orders = append(orders, order)
	}

	return orders, nil
}


