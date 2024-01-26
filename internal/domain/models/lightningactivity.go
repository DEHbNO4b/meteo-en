package models

type LightningActivity struct {
	count              int
	maxSignal          int
	signal             int
	cloudTypeRelation  float64
	groundTypeRelation float64
}

func (la *LightningActivity) SetCount(c int) {
	la.count = c
}

func (la *LightningActivity) SetMaxSignal(s int) {
	if s > la.maxSignal {
		la.maxSignal = s
	}
}

func (la *LightningActivity) SetSignal(s int) {
	// la.signal = (la.signal + s) / 2
	la.signal = s
}

func (la *LightningActivity) SetCloudTypeRelation(r float64) {
	la.cloudTypeRelation = r
}

func (la *LightningActivity) SetGroundTypeRelation(r float64) {
	la.groundTypeRelation = r
}
