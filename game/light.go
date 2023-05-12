package game

type Light struct {
	Radius int
}

func NewLight(r int) Light {
	return Light{
		Radius: r,
	}
}
