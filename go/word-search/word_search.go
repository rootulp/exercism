package wordsearch

import "fmt"

func Solve(words []string, puzzle []string) (result map[string][2][2]int, err error) {
	grid := NewGrid(puzzle)

	result = make(map[string][2][2]int)
	fmt.Printf("grid %v\n", grid)
	for _, word := range words {
		result[word] = grid.Search(word)
	}
	for k, v := range result {
		if v == [2][2]int{{-1, -1}, {-1, -1}} {
			err = fmt.Errorf("word %s not found", k)
		}
	}
	return result, err
}

type Grid struct {
	grid [][]rune
}

func NewGrid(puzzle []string) *Grid {
	grid := make([][]rune, len(puzzle))
	for i, row := range puzzle {
		grid[i] = []rune(row)
	}
	return &Grid{grid: grid}
}

func (g Grid) String() string {
	return fmt.Sprintf("%v", g.grid)
}

func (g Grid) Search(word string) (result [2][2]int) {
	for i, row := range g.grid {
		for j, v := range row {
			if v == rune(word[0]) {
				loc := dfs(g, i, j, word)
				if loc != [2][2]int{{-1, -1}, {-1, -1}} {
					return result
				}
			}
		}
	}
	return [2][2]int{{-1, -1}, {-1, -1}}
}

func dfs(grid Grid, row int, col int, word string) (result [2][2]int) {
	if len(word) == 0 {
		fmt.Printf("found %v at %v %v", word, row, col)
		return [2][2]int{{row, col}, {row, col}}
	}
	if row < 0 || col < 0 || row >= len(grid.grid) || col >= len(grid.grid[0]) || grid.grid[row][col] != rune(word[0]) {
		return [2][2]int{{-1, -1}, {-1, -1}}
	}
	temp := grid.grid[row][col]
	grid.grid[row][col] = ' '
	result = dfs(grid, row+1, col, word[1:])
	if result != [2][2]int{{-1, -1}, {-1, -1}} {
		return result
	}
	result = dfs(grid, row-1, col, word[1:])
	if result != [2][2]int{{-1, -1}, {-1, -1}} {
		return result
	}
	result = dfs(grid, row, col+1, word[1:])
	if result != [2][2]int{{-1, -1}, {-1, -1}} {
		return result
	}
	result = dfs(grid, row, col-1, word[1:])
	if result != [2][2]int{{-1, -1}, {-1, -1}} {
		return result
	}
	grid.grid[row][col] = temp
	return [2][2]int{{-1, -1}, {-1, -1}}
}

// type Point struct {
// 	x, y int
// }

// func (p Point) ToSlice() [2]int {
// 	return [2]int{p.x, p.y}
// }
