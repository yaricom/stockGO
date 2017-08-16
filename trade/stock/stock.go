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
	MinPrice float32
	MinPriceHistorical float32
	MaxPrice float32
	MaxPriceHistorical float32
	CurrPrice float32
	BidPrice float32
	Owned int
}

// Creates new stock with given name
func New(name string) Stock {
	return Stock{
		Name: name,
	}
}

func (s Stock) CanBeSold(history[]float32) bool {
	if s.Owned == 0 {
		return false
	}
	delta := s.bidMargin()
	return s.BidPrice < s.MaxPrice - delta
}

func (s Stock) CanBeBought(history[]float32) bool {
	if s.CurrPrice < s.MinPriceHistorical {
		return true
	}
	if s.Owned > 0 {
		return s.CurrPrice < s.BidPrice
	}
	delta := s.bidMargin()
	return s.CurrPrice < s.MinPrice + delta
}

func (s *Stock) Update(history[]float32, owned int) {
	// fill values
	s.CurrPrice = history[len(history) - 1]
	s.Owned = owned

	// update min /max
	sort.Float64s(history)// sort in increasing order
	min := history[0]
	if s.MinPrice > min {
		s.MinPrice = min
	}
	if s.MinPriceHistorical > min {
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

func (s Stock) bidMargin() float32 {
	margin := rand.Float32() * marginMultiplier
	if margin == 0 { margin = marginMultiplier }
	return (s.MaxPrice - s.MinPrice) * margin
}


