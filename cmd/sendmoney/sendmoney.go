//+build ignore

// This is a dummy program for solving the cannonical constraint programming
// puzzle SEND+MORE=MONEY.
//
// Along with other dummy programs, the intention is to explore what the API of
// the constraint engine may look like from the users perspective.

package main

import (
	"fmt"
	"log"

	"github.com/mewmew/constraint"
)

func main() {
	solutions, err := foo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("solutions:", solutions)
}

func foo() ([]constraint.Solution, error) {
	// Define variables within the search space.
	s := constraint.NewIntVar("s")
	e := constraint.NewIntVar("e")
	n := constraint.NewIntVar("n")
	d := constraint.NewIntVar("d")
	m := constraint.NewIntVar("m")
	o := constraint.NewIntVar("o")
	r := constraint.NewIntVar("r")
	y := constraint.NewIntVar("y")
	s.SetDomain(0, 9)
	e.SetDomain(0, 9)
	n.SetDomain(0, 9)
	d.SetDomain(0, 9)
	m.SetDomain(0, 9)
	o.SetDomain(0, 9)
	r.SetDomain(0, 9)
	y.SetDomain(0, 9)
	variables := []constraint.Variable{s, e, n, d, m, o, r, y}

	// Define constraints for the variables.
	sNotZero := constraint.NewRelation(s, constraint.NE, 0)
	mNotZero := constraint.NewRelation(m, constraint.NE, 0)
	distinct := constraint.NewDistinct(s, e, n, d, m, o, r, y)
	constraints := []constraint.Constraint{sNotZero, mNotZero, distinct}

	// Search for solutions satisfying the constraints.
	solutions := constraint.Solve(variables, constraints)

	return solutions, nil
}
