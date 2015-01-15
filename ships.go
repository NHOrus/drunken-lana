// ships
package main

type ship struct {
	Fatness int //Amount of cells in the field covered by ship in straight line
	Numbers int //Number of ships total
}

func newShip(F int, N int) ship {
	//creating new ship, could be done better, too
	// lazy to redo in the light of next one
	sh := new(ship)
	sh.Fatness = F
	sh.Numbers = N
	return *sh
}

func allShips() *map[string]ship {
	//Those are russian rules. Field size declared elsewhere
	s := make(map[string]ship)
	s["battleship"] = newShip(4, 1)
	s["cruiser"] = newShip(3, 2)
	s["destroyer"] = newShip(2, 3)
	s["cutter"] = newShip(1, 4)
	return &s
}

type rules struct { //rules declare what the hell is going on here.
	fWid, fLen int              //field parameters
	Ships      *map[string]ship //ship parameters
}

func getRules() (rules, error) {
	var r rules
	r.Ships = allShips()
	r.fWid, r.fLen = 10, 10 //again, russian rules, need some rework and unification
	return r, nil
}
