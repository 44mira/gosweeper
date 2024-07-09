package main

type Matrix struct {
	Cells [][]uint
}

// GetAdjacencyMatrix returns the corresponding adjacency matrix of a field
// which can then be used as a lookup table for the mine neighbors of a tile.
func (f *Field) GetAdjacencyMatrix() Matrix {
	tiles := f.Tiles
	x, y := len(tiles), len(tiles[0])

	cells := make([][]uint, x) // Create columns
	for i := range cells {
		cells[i] = make([]uint, y) // Create rows
	}

	// Mark all mines
	for i := range tiles {
		for j := range tiles[i] {
			if tiles[i][j].IsMine {
				cells[i][j] = 1
			}
		}
	}

	matrix := Matrix{cells}
	original := deepCopy(matrix.Cells)

	matrix.VShift(original)
	matrix.HShift(original)
	matrix.DShift(original)

	return matrix
}

// VShift is a function that increments a matrix in place 1 step up and down.
func (m *Matrix) VShift(original [][]uint) {
	cells := m.Cells

	for i := range cells {
		for j := range cells[i] {

			// if the cell has no mine, early return
			if original[i][j] < 1 {
				continue
			}

			// Increment above
			if i > 0 {
				cells[i-1][j]++
			}

			// Increment below
			if i < len(cells)-1 {
				cells[i+1][j]++
			}
		}
	}
}

// HShift is a function that increments a matrix inplace 1 step left and right.
func (m *Matrix) HShift(original [][]uint) {
	cells := m.Cells

	for i := range cells {
		for j := range cells[i] {

			// if the cell has no mine, early return
			if original[i][j] < 1 {
				continue
			}

			// Increment left
			if j > 0 {
				cells[i][j-1]++
			}

			// Increment right
			if j < len(cells[i])-1 {
				cells[i][j+1]++
			}
		}
	}
}

// HShift is a function that increments a matrix inplace on both diagonals.
func (m *Matrix) DShift(original [][]uint) {
	cells := m.Cells
	colLength := len(cells) - 1

	for i := range cells {
		for j := range cells[i] {
			// if the cell has no mine, early return
			if original[i][j] < 1 {
				continue
			}

			rowLength := len(cells[i]) - 1

			// Increment top-left
			if i > 0 && j > 0 {
				cells[i-1][j-1]++
			}

			// Increment bottom-left
			if i < colLength && j > 0 {
				cells[i+1][j-1]++
			}

			// Increment top-right
			if i > 0 && j < rowLength {
				cells[i-1][j+1]++
			}

			if i < colLength && j < rowLength {
				cells[i+1][j+1]++
			}
		}
	}
}

func deepCopy(matrix [][]uint) [][]uint {
	x, y := len(matrix), len(matrix[0])

	result := make([][]uint, x)
	for i := range result {
		result[i] = make([]uint, y) // Create rows
	}

	for i := range matrix {
		for j, v := range matrix[i] {
			result[i][j] = v
		}
	}

	return result
}
