package models

type Points struct {
	points []*CorrPoint
}

func (p *Points) Append(cp *CorrPoint) {
	p.points = append(p.points, cp)
}

// func (p *Points) Correlations() ([]string, error) {
// 	ans := make([]string, 0, 10)

// 	// metPar := make([][]float64, 0, 6)
// 	// ligPar := make([][]float64, 0, 6)

// 	ws := make([]float64, 0, len(p.points))
// 	mws := make([]float64, 0, len(p.points))
// 	r := make([]float64, 0, len(p.points))
// 	mr := make([]float64, 0, len(p.points))
// 	rr := make([]float64, 0, len(p.points))
// 	mrr := make([]float64, 0, len(p.points))

// 	c := make([]float64, 0, len(p.points))
// 	mps := make([]float64, 0, len(p.points))
// 	mns := make([]float64, 0, len(p.points))
// 	ps := make([]float64, 0, len(p.points))
// 	ns := make([]float64, 0, len(p.points))
// 	cr := make([]float64, 0, len(p.points))
// 	gr := make([]float64, 0, len(p.points))

// 	for _, el := range p.points {
// 		ws = append(ws, el.WindSpeed)
// 		mws = append(mws, el.MaxWindSpeed)
// 		r = append(r, el.Rain)
// 		mr = append(mr, el.MaxRain)
// 		rr = append(rr, el.RainRate)
// 		mrr = append(mrr, el.MaxRainRate)

// 		c = append(c, float64(el.count))
// 		mps = append(mps, float64(el.maxPozitiveSignal))
// 		mns = append(mns, float64(el.maxNegativeSignal))
// 		ps = append(ps, float64(el.pozitiveSignal))
// 		ns = append(ns, float64(el.negativeSignal))
// 		cr = append(cr, el.cloudTypeRelation)
// 		gr = append(gr, el.groundTypeRelation)
// 	}

// 	// metPar = append(metPar, ws, mws, r, rr, mr, mrr)

// 	// ligPar = append(ligPar, c, mps, mns, ps, ns, cr, gr)

// 	// for i:=0;i<len(metPar);i++{
// 	// 	for j:=0;j<len(ligPar);j++{

// 	// 		a, _ := stats.Correlation(metPar[i], ligPar[j])
// 	// 		rounded, _ := stats.Round(a, 5)
// 	// 		fmt.Printf(" correlation betwen WindSpeed and count of lightning: %v", rounded)
// 	// 	}
// 	// }

// 	a, _ := stats.Correlation(ws, c)
// 	rounded, _ := stats.Round(a, 5)
// 	fmt.Printf(" correlation betwen WindSpeed and count of lightning: %v", rounded)

// 	return ans, nil
// }
