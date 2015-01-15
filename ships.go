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
	//Those are russian rules. May need to add rules for field size later.
	s := make(map[string]ship)
	s["battleship"] = newShip(4, 1)
	s["cruiser"] = newShip(3, 2)
	s["destroyer"] = newShip(2, 3)
	s["cutter"] = newShip(1, 4)
	return &s
}
