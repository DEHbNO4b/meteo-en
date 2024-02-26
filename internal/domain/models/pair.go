package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

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

func (p *Pair) OutputData() error {
	file, err := os.Create("./public/out/" + p.name + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewWriter(file)
	r.Comma = ';'

	record := make([]string, 2, 2)

	for i := range len(p.la) {
		record[0] = strconv.FormatFloat(p.mp[i], 'f', -1, 64)
		record[1] = strconv.FormatFloat(p.la[i], 'f', -1, 64)
		if err := r.Write(record); err != nil {
			continue
		}
	}
	r.Flush()

	if err := r.Error(); err != nil {
		return err
	}
	return nil
}
func (p *Pair) OutputData2() error {
	file, err := os.Create("./public/out/" + p.name + "_2" + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	r := csv.NewWriter(file)
	r.Comma = ';'

	record := make([]string, len(p.mp))

	for i := range len(p.mp) {
		record[i] = strconv.FormatFloat(p.mp[i], 'f', -1, 64)

	}
	if err := r.Write(record); err != nil {
		return err
	}

	record1 := make([]string, len(p.la))
	for i := range len(p.la) {
		record1[i] = strconv.FormatFloat(p.la[i], 'f', -1, 64)

	}
	if err := r.Write(record1); err != nil {
		return err
	}
	r.Flush()
	if err := r.Error(); err != nil {
		return err
	}
	return nil
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
