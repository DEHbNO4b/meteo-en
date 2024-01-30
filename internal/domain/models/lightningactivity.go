package models

import "math"

type LightningActivity struct {
	// Strokes            []*StrokeEN
	count              int
	maxPozitiveSignal  int64
	maxNegativeSignal  int64
	signalAbs          float64
	pozitiveSignal     int64
	negativeSignal     int64
	cloudTypeRelation  float64
	groundTypeRelation float64
}

func NewLActivity(strokes []*StrokeEN) LightningActivity {

	la := LightningActivity{}

	var (
		sigAbs  float64
		maxPoz  int64
		maxNeg  int64
		poz     int64
		neg     int64
		clouds  float64
		grounds float64
	)
	for _, el := range strokes {
		if el.Signal() > 0 {
			poz = (poz + el.Signal()) / 2
			if el.Signal() > maxPoz {
				maxPoz = el.Signal()
			}
		} else if el.Signal() < 0 {
			neg = (neg + el.Signal()) / 2
			if el.Signal() < maxNeg {
				maxNeg = el.Signal()
			}
		}
		sigAbs = (sigAbs + math.Abs(float64(el.signal))/2)
		if el.Cloud() {
			clouds++
		} else {
			grounds++
		}
	}

	la.signalAbs = sigAbs
	la.count = len(strokes)
	la.maxPozitiveSignal = maxPoz
	la.maxNegativeSignal = maxNeg
	la.pozitiveSignal = poz
	la.negativeSignal = neg
	la.cloudTypeRelation = clouds / float64(len(strokes))
	la.groundTypeRelation = grounds / float64(len(strokes))

	return la
}

func (la *LightningActivity) SetCount(c int) {
	la.count = c
}

func (la *LightningActivity) SetMaxPozitiveSignal(s int64) {
	if s > la.maxPozitiveSignal {
		la.maxPozitiveSignal = s
	}
}

func (la *LightningActivity) SetMaxNegativeSignal(s int64) {
	if s < la.maxNegativeSignal {
		la.maxNegativeSignal = s
	}
}

func (la *LightningActivity) SetPozSignal(s int64) {
	// la.signal = (la.signal + s) / 2
	la.pozitiveSignal = s
}

func (la *LightningActivity) SetNegSignal(s int64) {
	// la.signal = (la.signal + s) / 2
	la.negativeSignal = s
}

func (la *LightningActivity) SetCloudTypeRelation(r float64) {
	la.cloudTypeRelation = r
}

func (la *LightningActivity) SetGroundTypeRelation(r float64) {
	la.groundTypeRelation = r
}

func (la *LightningActivity) Count() int {
	return la.count
}

func (la *LightningActivity) AbsSig() float64 {
	return la.signalAbs
}

func (la *LightningActivity) MaxPozSig() int64 {
	return la.maxPozitiveSignal
}

func (la *LightningActivity) MaxNegSig() int64 {
	return la.maxNegativeSignal
}

func (la *LightningActivity) PozSig() int64 {
	return la.pozitiveSignal
}

func (la *LightningActivity) NegSig() int64 {
	return la.negativeSignal
}

func (la *LightningActivity) CloudTypeRel() float64 {
	return la.cloudTypeRelation
}

func (la *LightningActivity) GroundTypeRel() float64 {
	return la.groundTypeRelation
}
