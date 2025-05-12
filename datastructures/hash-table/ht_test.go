package ht

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuick(t *testing.T) {
	table := New(4)

	table.Put("uno", "1")
	table.Put("dos", "2")
	table.Put("otro", "mas")
	table.Put("atra", "atramas")

	assert.Equal(t, 8, table.Buckets())

	fmt.Println(table.Table)
	table.Delete("atra")
	fmt.Println(table.Table)
	assert.Equal(t, 3, table.Size())
	assert.Equal(t, nil, table.Get("atra"))
}

func TestPutAndGetBasic(t *testing.T) {
	table := New(4)

	table.Put("uno", "1")
	table.Put("dos", "2")
	table.Put("tres", "3")

	assert.Equal(t, "1", table.Get("uno"))
	assert.Equal(t, "2", table.Get("dos"))
	assert.Equal(t, "3", table.Get("tres"))
}

func TestOverwriteValue(t *testing.T) {
	table := New(4)

	table.Put("clave", "valor original")
	assert.Equal(t, "valor original", table.Get("clave"))

	table.Put("clave", "nuevo valor")
	assert.Equal(t, "nuevo valor", table.Get("clave"))
}

func TestGetNonExistentKey(t *testing.T) {
	table := New(4)

	table.Put("existe", "valor")
	assert.Nil(t, table.Get("no-existe"))
}

func TestHashCollisions(t *testing.T) {
	table := New(1) // fuerza colisiones

	table.Put("a", "valor1")
	table.Put("b", "valor2")
	table.Put("c", "valor3")

	assert.Equal(t, "valor1", table.Get("a"))
	assert.Equal(t, "valor2", table.Get("b"))
	assert.Equal(t, "valor3", table.Get("c"))
}

func TestValueTypePreservation(t *testing.T) {
	table := New(4)

	table.Put("entero", 123)
	table.Put("bool", true)

	assert.Equal(t, 123, table.Get("entero"))
	assert.Equal(t, true, table.Get("bool"))

	val := table.Get("entero")
	if i, ok := val.(int); ok {
		assert.Equal(t, 123, i)
	} else {
		t.Errorf("esperado int, recibido %T", val)
	}
}
