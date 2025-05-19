package mtx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicOps(t *testing.T) {
	mtx := NewMatrix(3, 5)

	row1 := []any{"a", 2, 3, 4, 5}
	row2 := []any{"b", 6, 7, 8, 9}
	row3 := []any{"c", 10, 11, 12, 13}

	for j := 0; j < 5; j++ {
		mtx.Set(0, j, row1[j])
		mtx.Set(1, j, row2[j])
		mtx.Set(2, j, row3[j])
	}

	testMtx := make([][]any, 3)
	testMtx[0], testMtx[1], testMtx[2] = row1, row2, row3

	fmt.Println(mtx)

	assert.Equal(t, testMtx, mtx.Elements)

	testTranspose := [][]any{
		{"a", "b", "c"},
		{2, 6, 10},
		{3, 7, 11},
		{4, 8, 12},
		{5, 9, 13},
	}

	transposed := mtx.Transpose()

	assert.Equal(t, testTranspose, transposed.Elements)

	headers := []string{"LETRA", "NUM1", "NUM2", "NUM3", "NUM4"}

	if err := mtx.AddHeaders(headers); err != nil {
		t.Error(err)
	}

	for i := range mtx.Elements[0] {
		assert.Equal(t, headers[i], mtx.Elements[0][i])
	}
}
