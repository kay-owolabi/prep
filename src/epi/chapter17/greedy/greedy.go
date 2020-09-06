package greedy

import "sort"

// 17.00 Greedy Algorithms and Invariants
func ChangeMaking(cents int) int {
	coins := []int{100, 50, 25, 10, 5, 1}
	var numCoins int
	for _, coin := range coins {
		numCoins += cents / coin
		cents %= coin
	}
	return numCoins
}

// 17.04 The 3-sum problem
func HasThreeSum(A []int, t int) bool {
	//sort.Search()
	sort.Ints(A)
	for _, a := range A {
		if hasTwoSum(A, t-a) {
			return true
		}
	}
	return false
}

func hasTwoSum(A []int, t int) bool {
	i, j := 0, len(A)-1
	for i <= j {
		if A[i]+A[j] == t {
			return true
		} else if A[i]+A[j] < t {
			i++
		} else {
			j--
		}
	}
	return false
}

type Grid [][]bool

func (g Grid) NumCol() int  { return len(g[0]) }
func (g Grid) NumRows() int { return len(g) }
func (g Grid) IsSafe(row, col int) bool {
	if g.RowCheck(row, col) && g.ColCheck(row, col) && g.DiagonalCheck(row, col) {
		return true
	}
	return false
}
func (g Grid) PlaceQueen(row, col int)  { g[row][col] = true }
func (g Grid) RemoveQueen(row, col int) { g[row][col] = false }
func (g Grid) Solve() bool              { return g.solve(0) }

func (g Grid) solve(col int) bool {
	if col >= g.NumCol() {
		return true
	}

	for rowToTry := 0; rowToTry < g.NumRows(); rowToTry++ {
		if g.IsSafe(rowToTry, col) {
			g.PlaceQueen(rowToTry, col)
			if g.solve(col + 1) {
				return true
			}
			g.RemoveQueen(rowToTry, col)
		}
	}
	return false
}

func (g Grid) RowCheck(row, col int) bool {
	for j := 0; j < len(g[row]); j++ {
		if j != col && g[row][j] {
			return false
		}
	}
	return true
}

func (g Grid) ColCheck(row, col int) bool {
	for i := 0; i < len(g); i++ {
		if i != row && g[i][col] {
			return false
		}
	}
	return true
}

func (g Grid) DiagonalCheck(row, col int) bool {
	zero := min(row, col)
	for i, j := row-zero, col-zero; i < len(g) && j < len(g[row]); i, j = i+1, j+1 {
		if (i != row || j != col) && g[i][j] {
			return false
		}
	}

	for i, j := row-zero, col-zero; i < len(g) && 0 <= j; i, j = i+1, j-1 {
		if (i != row || j != col) && g[i][j] {
			return false
		}

	}
	return true
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

type Sudoku [][]int

/*
func Solve(conf configuration) {
	if no_more_choices {
		return conf_is_goal_state
	}
	for all_available_choices {
		try one choice c
			// solve from here, if work out, you're done
		if Solve(conf with choice c made) {
			return true
		}
		unmake choice c
	}
}
*/

func (s Sudoku) SolveSudoku() bool {
	var row, col int
	if !s.findUnassignedLocation(&row, &col) {
		return true
	}
	for num := 1; num <= 9; num++ {
		if s.noConflicts(row, col, num) {
			s[row][col] = num
			if s.SolveSudoku() {
				return true
			}
			s[row][col] = 0
		}
	}
	return false
}

func (s Sudoku) findUnassignedLocation(row, col *int) bool {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] == 0 {
				*row, *col = i, j
				return true
			}
		}
	}
	return false
}

func (s Sudoku) noConflicts(row int, col int, num int) bool {
	return s.rowCheck(row, num) && s.colCheck(col, num) && s.boxCheck(row, col, num)
}

func (s Sudoku) rowCheck(row, num int) bool {
	for j := 0; j < len(s[row]); j++ {
		if s[row][j] == num {
			return false
		}
	}
	return true
}

func (s Sudoku) colCheck(col int, num int) bool {
	for i := 0; i < len(s); i++ {
		if s[i][col] == num {
			return false
		}
	}
	return true
}

func (s Sudoku) boxCheck(row int, col int, num int) bool {
	row0, col0 := row/3, col/3
	for i := row0 * 3; i < (row0+1)*3; i++ {
		for j := col0 * 3; j < (col0+1)*3; j++ {
			if s[i][j] == num {
				return false
			}
		}
	}
	return true
}
