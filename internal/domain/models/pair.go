package models

import (
	"fmt"

	"github.com/montanaflynn/stats"
)

type Pair struct {
	mp   []float64
	la   []float64
	corr float64
	name string
	err  error
}

func NewPair(name string) *Pair {

	mp := make([]float64, 0, 100)
	la := make([]float64, 0, 100)

	return &Pair{mp: mp, la: la, name: name}
}

func (p *Pair) String() string {
	if p.err != nil {
		return p.err.Error()
	}
	return fmt.Sprintf("Data len: %d, correlation coefficient  %s = %v", len(p.mp), p.name, p.corr)
}

func (p *Pair) AddPair(first, second float64) {
	p.mp = append(p.mp, first)
	p.la = append(p.la, second)
}

func (p *Pair) Calculate() error {
	corr, err := stats.Correlation(p.mp, p.la)
	if err != nil {
		p.err = err
		return err
	}
	r, err := stats.Round(corr, 5)
	if err != nil {
		p.corr = corr
		return nil
	}
	p.corr = r
	return nil
}

func (p *Pair) CorrCoef() float64 {
	return p.corr
}

// type Pairs struct {
// 	pairs []*Pair
// }

// func NewPairs() *Pairs {
// 	ps := make([]*Pair, 0, 10)
// 	return &Pairs{
// 		pairs: ps,
// 	}
// }
// func (ps *Pairs) Add(p *Pair) {

// 	ps.pairs = append(ps.pairs, p)

// }

// func (ps *Pairs) Pairs() []*Pair {
// 	return ps.pairs
// }

// func (ps *Pairs) Len() int {
// 	return len(ps.pairs)
// }
// func (ps *Pairs) Less(i, j int) bool {
// 	return false
// }

// // func (ps *Pairs) Swap(i, j int)
