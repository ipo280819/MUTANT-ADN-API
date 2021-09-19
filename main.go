package main

import (
	"fmt"

	"github.com/ipo280819/MUTANT-ADN-API/services"
)

func main() {
	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"AACCCA",
	}
	human := services.Mutant{Adn: dna}
	isMutant, _ := human.IsMutant()
	fmt.Printf("Is mutant: %v\n", isMutant)
}
