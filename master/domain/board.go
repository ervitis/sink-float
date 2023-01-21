package domain

const (
	N = 7

	W = N
	H = N
)

type Direction int

const (
	Vertical Direction = iota + 1
	Horizontal
)

type Board struct {
	Width  uint64
	Height uint64
	panel  [][]rune
}

type Position struct {
	X int
	Y int
}

type ShipPositions struct {
	Init Position
	End  Position
}

type Game struct {
	board Board
	ships map[IBattleShip]ShipPositions
}

func New() *Game {
	panel := make([][]rune, W)
	for i := 0; i < H; i++ {
		panel[i] = make([]rune, H)
	}

	x1 := GenerateRandNumber(0, N-ShipSize)
	y1 := GenerateRandNumber(0, N-ShipSize)
	d1 := GenerateDirection()

	xf := x1
	yf := y1
	acc := 0
	for i := 0; i < 3; i++ {
		if d1 == Horizontal {
			acc = x1 + i
			panel[x1][acc] = 'S'
		} else {
			acc = y1 + i
			panel[acc][y1] = 'S'
		}
	}

	if d1 == Horizontal {
		xf = acc
	} else {
		yf = acc
	}
	ships := map[IBattleShip]ShipPositions{
		NewSubmarine(): {
			Init: Position{
				X: x1,
				Y: y1,
			},
			End: Position{
				X: xf,
				Y: yf,
			},
		},
	}

	return &Game{
		board: Board{
			Width:  W,
			Height: H,
			panel:  panel,
		},
		ships: ships,
	}
}
