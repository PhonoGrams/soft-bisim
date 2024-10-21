package soft_bisim

type Weights struct {
	Match          float64
	Replace        float64
	Insert         float64
	Delete         float64
	Transposition  float64
	Merge          float64
	Split          float64
	CaseChange     float64
	PhoneticChange float64
}
