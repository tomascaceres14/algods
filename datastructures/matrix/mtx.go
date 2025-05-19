package mtx

import (
	"errors"
	"fmt"
)

//
// func (m *Matrix) Clone() *Matrix
// func (m *Matrix) Clear()
// func (m *Matrix) ToSlice() [][]any
// func (m *Matrix) String() string

type Matrix struct {
	rows, cols int
	isEmpty    bool
	Elements   [][]any
}

func NewMatrix(rows, cols int) *Matrix {
	m := Matrix{}
	m.rows, m.cols = rows, cols
	m.isEmpty = true
	mtx := make([][]any, rows)
	for i := range rows {
		lst := make([]any, cols)
		for j := range cols {
			lst[j] = nil
		}

		mtx[i] = lst
	}
	m.Elements = mtx
	return &m
}

func (m *Matrix) Rows() int {
	return m.rows
}

func (m *Matrix) Cols() int {
	return m.cols
}

func (m *Matrix) Shape() (int, int) {
	return m.rows, m.cols
}

func (m *Matrix) IsEmpty() bool {
	return m.isEmpty
}

func (m *Matrix) IsSquare() bool {
	return m.cols == m.rows
}

func (m *Matrix) Set(row, col int, val any) error {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		return errors.New("cell out of bounds.")
	}
	m.isEmpty = false
	m.Elements[row][col] = val

	return nil
}

func (m *Matrix) Get(row, col int) (any, error) {
	if row < 0 || row >= m.rows || col < 0 || col >= m.cols {
		return nil, errors.New("cell out of bounds.")
	}

	return m.Elements[row][col], nil
}

func (m *Matrix) Transpose() *Matrix {
	if m.IsEmpty() {
		return m
	}

	mtx := NewMatrix(m.cols, m.rows)

	for i := range mtx.cols {
		for j := range mtx.rows {
			mtx.Elements[j][i] = m.Elements[i][j]
		}
	}

	return mtx
}

func (m *Matrix) AppendMatrix(mtx *Matrix) error {
	if m.cols != mtx.cols {
		return errors.New("Number of columns don't match")
	}

	m.Elements = append(m.Elements, mtx.Elements...)
	return nil
}

func (m *Matrix) AddRow(row []any) {
	if len(row) <= 0 {
		return
	}

	m.Elements = append(m.Elements, row)
	m.rows += 1
}

func (m *Matrix) AddHeaders(headers []string) error {

	if len(headers) != m.cols {
		return errors.New("Wrong number of parameters")
	}

	newHeaders := make([]any, m.cols)
	for i := range newHeaders {
		newHeaders[i] = headers[i]
	}

	m.Elements = append(m.Elements, []any{})
	copy(m.Elements[1:], m.Elements)
	m.Elements[0] = newHeaders

	return nil
}

func (m *Matrix) String() string {
	result := ""

	mtx := m.Elements
	for i := range m.rows {
		row := ""
		for j := range m.rows {
			row += fmt.Sprintf("|%v|", mtx[i][j])
		}
		row += "\n"
		result += row
	}

	return result
}
