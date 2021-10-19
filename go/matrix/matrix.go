package matrix

type Matrix struct {
}

func New(s string) (*Matrix, error) {
	return &Matrix{}, nil
}

func (m Matrix) Cols() [][]int {
	return [][]int{}
}

func (m Matrix) Rows() [][]int {
	return [][]int{}
}

func (m Matrix) Set(row int, column int, value int) bool {
	return true
}
