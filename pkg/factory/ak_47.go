package factory

type Ak47 struct {
	Gun
}

func newAK47() IGun {
	return &Ak47{
		Gun: Gun{
			Name:  "Ak47 gun",
			Power: 4,
		},
	}
}
