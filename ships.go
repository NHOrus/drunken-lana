// ships
package main

type ship struct {
	Fatness int
	Numbers int
}

func newShip(F int, N int) ship {
	sh := new(ship)
	sh.Fatness = F
	sh.Numbers = N
	return *sh
}

func allShips() *map[string]ship {
	s := make(map[string]ship)
	s["battleship"] = newShip(4, 1)
	s["cruiser"] = newShip(3, 2)
	s["destroyer"] = newShip(2, 3)
	s["cutter"] = newShip(1, 4)
	return &s
}

type rules struct {
	fWid, fLen int
	Ships      *map[string]ship
}

func getRules() (rules, error) {
	var r rules
	r.Ships = allShips()
	r.fWid, r.fLen = 10, 10
	return r, nil
}
