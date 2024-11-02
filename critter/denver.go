package critter

type Denver struct {
	Name      string
	Strength  int
	Speed     int
	Health    int
	MaxHealth int
}

func NewDenver() *Denver {
	// TODO make random
	return &Denver{
		Name:      "MyDenver",
		Strength:  10,
		Speed:     10,
		Health:    10,
		MaxHealth: 10,
	}
}
