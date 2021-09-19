package services

import (
	"errors"
	"sync"
)

const (
	MAX_SUBSEQUENCES = 2

	SIZE_SUBSEQUENCE = 4
)

type Mutant struct {
	Adn         []string
	counter     int
	lockCounter sync.Mutex
}

func (mutant *Mutant) isValidAdn() bool {
	n := len(mutant.Adn)
	m := len(mutant.Adn[0])
	return n == m
}

func (mutant *Mutant) IsMutant() (bool, error) {

	if !mutant.isValidAdn() {
		return false, errors.New("this Adn can't be analize, wrong dimensions")
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		mutant.runHorizontalSearch()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mutant.runVerticalSearch()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mutant.runDiagonalSearch()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		mutant.runDiagonalInvSearch()
	}()

	wg.Wait()

	return mutant.counter >= MAX_SUBSEQUENCES, nil
}

func (mutant *Mutant) verifySequence(sequence []byte) int {

	prevLetter := sequence[0]
	keepSequence := true
	lastIndex := 0 //Prevent reuse cases when we know that this will fail
	for k := 1; k < SIZE_SUBSEQUENCE; k++ {
		currLetter := sequence[k]

		if prevLetter == currLetter {
			prevLetter = currLetter
			lastIndex = k
		} else {
			keepSequence = false
			break
		}
	}

	if keepSequence {
		mutant.lockCounter.Lock()
		mutant.counter += 1
		mutant.lockCounter.Unlock()
		lastIndex = 0
	}
	return lastIndex
}

func (mutant *Mutant) runHorizontalSearch() {

	n := len(mutant.Adn)
	for i := 0; i < n; i++ {
		row := mutant.Adn[i]
		for j := 0; j < n-SIZE_SUBSEQUENCE+1; {
			if mutant.counter >= MAX_SUBSEQUENCES {
				return
			}
			lastIndex := mutant.verifySequence([]byte(row[j : j+SIZE_SUBSEQUENCE]))
			j += 1 + lastIndex
		}
	}
}

func (mutant *Mutant) runVerticalSearch() {

	n := len(mutant.Adn)
	for i := 0; i < n; i++ {
		for j := 0; j < n-SIZE_SUBSEQUENCE+1; {
			if mutant.counter >= MAX_SUBSEQUENCES {
				return
			}
			seq := mutant.getColumnFromRow(i, j)
			lastIndex := mutant.verifySequence(seq)
			j += 1 + lastIndex
		}
	}
}

func (mutant *Mutant) getColumnFromRow(col int, row int) []byte {
	seq := []byte{}
	for i := 0; i < SIZE_SUBSEQUENCE; i++ {
		seq = append(seq, mutant.Adn[row+i][col])
	}
	return seq
}

func (mutant *Mutant) runDiagonalSearch() {

	n := len(mutant.Adn)
	for i := 0; i < n-SIZE_SUBSEQUENCE+1; i++ { // horizontal / vertical movs

		if mutant.counter >= MAX_SUBSEQUENCES {
			return
		}
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			for j := 0; j < n-SIZE_SUBSEQUENCE+1; j++ { // horizontal movs
				if mutant.counter >= MAX_SUBSEQUENCES {
					return
				}

				seq := mutant.getDiagonalFrom(i, j)
				lastIndex := mutant.verifySequence(seq)
				j += 1 + lastIndex

			}
		}()
		go func() {
			defer wg.Done()

			for j := 0; j < n-SIZE_SUBSEQUENCE; { // vertical movs
				if mutant.counter >= MAX_SUBSEQUENCES {
					return
				}
				seq := mutant.getDiagonalFrom(j+1, i)
				lastIndex := mutant.verifySequence(seq)
				j += 1 + lastIndex

			}
		}()
		wg.Wait()

	}
}
func (mutant *Mutant) getDiagonalFrom(row int, col int) []byte {
	seq := []byte{}
	for i := 0; i < SIZE_SUBSEQUENCE; i++ {
		seq = append(seq, mutant.Adn[row+i][col+i])
	}
	return seq
}

func (mutant *Mutant) runDiagonalInvSearch() {

	n := len(mutant.Adn)
	for i := 0; i < n-SIZE_SUBSEQUENCE+1; i++ { // horizontal / vertical movs

		if mutant.counter >= MAX_SUBSEQUENCES {
			return
		}
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			for j := 0; j < n-SIZE_SUBSEQUENCE+1; { // horizontal movs
				if mutant.counter >= MAX_SUBSEQUENCES {
					return
				}
				k := n - j - 1
				seq := mutant.getDiagonalInvFrom(i, k)
				lastIndex := mutant.verifySequence(seq)
				j += 1 + lastIndex

			}
		}()
		go func() {
			defer wg.Done()

			for j := 0; j < n-SIZE_SUBSEQUENCE; { // vertical movs
				if mutant.counter >= MAX_SUBSEQUENCES {
					return
				}
				k := n - i - 1

				seq := mutant.getDiagonalInvFrom(j+1, k)
				lastIndex := mutant.verifySequence(seq)
				j += 1 + lastIndex

			}
		}()
		wg.Wait()

	}
}
func (mutant *Mutant) getDiagonalInvFrom(row int, col int) []byte {
	seq := []byte{}
	for i := 0; i < SIZE_SUBSEQUENCE; i++ {
		seq = append(seq, mutant.Adn[row+i][col-i])
	}
	return seq
}
