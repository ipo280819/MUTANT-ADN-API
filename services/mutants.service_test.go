package services

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMutant(t *testing.T) {
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "AACCCA"}
	human := Mutant{Adn: dna}
	isMutant, _ := human.IsMutant()
	assert.Equal(t, true, isMutant)
}

/*
	this test verify that the subsequence was found
	and that lock of counter is working
*/
func TestVerifySubsequence(t *testing.T) {
	sequence := "AAAA"
	calls := 4
	wg := sync.WaitGroup{}
	human := Mutant{}

	for i := 0; i < calls; i++ {
		wg.Add(1)
		func() {
			defer wg.Done()
			human.verifySequence([]byte(sequence))
		}()
	}
	wg.Wait()
	assert.Equal(t, calls, human.counter)
}

func TestVerifySubsequenceNotFound(t *testing.T) {
	sequence := "AABA"
	human := Mutant{}
	human.verifySequence([]byte(sequence))

	assert.Equal(t, 0, human.counter)
}

func TestRunHorizontalSearch(t *testing.T) {
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "AACCCC"}
	human := Mutant{Adn: dna}

	human.runHorizontalSearch()

	assert.Equal(t, 2, human.counter)
}

/*
	Testing case for CCCCCT
	[CCCC]CT
	C[CCCC]T
*/
func TestHorizontalSubsequencesLinked(t *testing.T) {
	dna := []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCCT", "AATCCC"}
	human := Mutant{Adn: dna}

	human.runHorizontalSearch()

	assert.Equal(t, 2, human.counter)
}

func TestGetCol(t *testing.T) {
	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"AACCTC"}
	human := Mutant{Adn: dna}

	seq := human.getColumnFromRow(3, 2)

	assert.Equal(t, "TACC", string(seq))
}

func TestRunVerticalSearch(t *testing.T) {
	dna := []string{
		"ATGCGA",
		"CAGTCT",
		"TTATGT",
		"AGAAGT",
		"CCCCTT",
		"AACCCT",
	}
	human := Mutant{Adn: dna}

	human.runVerticalSearch()

	assert.Equal(t, 2, human.counter)
}

func TestGetDiag(t *testing.T) {
	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAAG",
		"CCCCTA",
		"AACCTC"}
	human := Mutant{Adn: dna}

	seq := human.getDiagonalFrom(0, 1)

	assert.Equal(t, "TGTA", string(seq))
}

func TestRunDiagonalSearch(t *testing.T) {
	dna := []string{
		"ATGCGA",
		"TGGGGC",
		"TTATGT",
		"ATGAAG",
		"CCTTTG",
		"CCCTTG",
	}
	human := Mutant{Adn: dna}

	human.runDiagonalSearch()

	assert.Equal(t, 2, human.counter)
}

func TestGetDiagInv(t *testing.T) {
	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAAG",
		"CCCCTA",
		"AACCTC",
	}
	human := Mutant{Adn: dna}

	seq := human.getDiagonalInvFrom(0, 5)

	assert.Equal(t, "AGTA", string(seq))
}

func TestRunDiagonalInvSearch(t *testing.T) {
	dna := []string{
		"CTGTGCT",
		"CCTCGCT",
		"CTGTATT",
		"TTACGCA",
		"AGCATAT",
		"CGTTAAT",
		"AAGATCA",
	}
	human := Mutant{Adn: dna}

	human.runDiagonalInvSearch()

	assert.Equal(t, 2, human.counter)
}
func TestIsValidAdn(t *testing.T) {
	dna := []string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAAG",
		"CCCCTA",
		"AACCTC",
	}
	human := Mutant{Adn: dna}

	assert.Equal(t, true, human.isValidAdn())
}

func TestIsNotValidAdn(t *testing.T) {
	dna := []string{
		"ATGCGA",
	}
	human := Mutant{Adn: dna}

	assert.Equal(t, false, human.isValidAdn())
}
