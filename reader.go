package main

import (
	"io/ioutil"
	"strings"
	"fmt"
)

func ReadFile(filename string) (Puzzle, error) {

	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return Puzzle{}, err
	}

	return createPuzzle(string(dat[:]))
}

func createPuzzle(model string) (Puzzle, error) {

	rows := strings.Split(strings.Trim(strings.ToUpper(model), " \t\n\r"), "\n")
	var maxLen int8 = 0
	var grid = make(Grid, len(rows), len(rows[0]))
	for index, row := range rows {
		if len(row) == 0 {
			continue
		}
		var rowValues = []uint8{}
		var counter int8 = 0
		for _, char := range row {
			val := int8(char)
			if val == 32 {
				continue
			}
			counter ++
			if val > 48 && val <= 57 {
				rowValues = append(rowValues, uint8(val-48))
			} else if val >= 65 && val <= 90 {
				rowValues = append(rowValues, uint8(val-55))
			} else {
				return Puzzle{}, fmt.Errorf("only numbers [1-9] and latin characters allowed in model. Loaded model [%s].", filename)
			}
		}
		if maxLen < counter {
			maxLen = counter
		}
		grid[index] = rowValues
	}

	pieces := GetPiecesFromGrid(grid)
	var solutions []Grid

	var puzzle = Puzzle{
		pieces,
		copyGrid(grid[:]),
		grid[:],
		max(int8(len(rows)), maxLen),
		minPieceSize(pieces),
		&solutions,
		false,
		false,
		&WinInfo{},
	}

	return puzzle, nil
}

func max(a int8, b int8) int8 {
	if a >= b {
		return a
	}
	return b
}

func minPieceSize(pieces []Piece) int {
	var min int = 255
	for _, piece := range pieces {
		if piece.Size < min {
			min = piece.Size
		}
	}

	return min
}
