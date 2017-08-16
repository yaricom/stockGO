// Provides utilities to create and operate over stock data
package stock

import (
	"math/rand"
	"sort"
)

// The stock bid margin multiplier
const marginMultiplier = 0.25

// The stock data holder to hold information about particular
// stock
type Stock struct {
	Name string
	MinPrice float64
	MinPriceHistorical float64
	MaxPrice float64
	MaxPriceHistorical float64
	CurrPrice float64
	BidPrice float64
	Owned int
}

// Creates new stock with given name
func New(name string) Stock {
	return Stock{
		Name: name,
	}
}

func (s Stock) CanBeSold(history[]float64) bool {
	if s.Owned == 0 {
		return false
	}
	delta := s.bidMargin()
	return s.BidPrice < s.CurrPrice - delta
}

func (s Stock) CanBeBought(history[]float64) bool {
	if s.CurrPrice < s.MinPriceHistorical {
		return true
	}
	if s.Owned > 0 {
		return s.CurrPrice < s.BidPrice
	}
	delta := s.bidMargin()
	return s.CurrPrice < s.MinPrice + delta
}

func (s *Stock) Update(history[]float64, owned int) {
	// fill values
	s.CurrPrice = history[len(history) - 1]
	s.Owned = owned

	// update min /max
	sort.Float64s(history)// sort in increasing order
	min := history[0]
	if s.MinPrice > min || s.MinPrice == 0 {
		s.MinPrice = min
	}
	if s.MinPriceHistorical > min || s.MinPriceHistorical == 0 {
		s.MinPriceHistorical = min
	}
	max := history[len(history) - 1]
	if s.MaxPrice < max {
		s.MaxPrice = max
	}
	if s.MaxPriceHistorical < max {
		s.MaxPriceHistorical = max
	}
}

func (s Stock) bidMargin() float64 {
	margin := rand.Float64() * marginMultiplier
	if margin == 0 { margin = marginMultiplier }
	return (s.MaxPrice - s.MinPrice) * margin
}


