package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSafeMap_Insert(t *testing.T) {
	m := New[int, string]()
	m.Insert(1, "value1")

	v, err := m.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, "value1", v)
}

func TestSafeMap_Get(t *testing.T) {
	m := New[int, string]()
	m.Insert(1, "value1")

	v, err := m.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, "value1", v)

	_, err = m.Get(2)
	assert.NotNil(t, err)
}

func TestSafeMap_Update(t *testing.T) {
	m := New[int, string]()
	m.Insert(1, "value1")

	err := m.Update(2, "value2")
	assert.NotNil(t, err)

	err = m.Update(1, "value2")
	assert.Nil(t, err)

	v, err := m.Get(1)
	assert.Nil(t, err)
	assert.Equal(t, "value2", v)
}

func TestSafeMap_Delete(t *testing.T) {
	m := New[int, string]()
	m.Insert(1, "value1")

	err := m.Delete(2)
	assert.NotNil(t, err)

	err = m.Delete(1)
	assert.Nil(t, err)

	_, err = m.Get(1)
	assert.NotNil(t, err)
}

func TestSafeMap_Has(t *testing.T) {
	m := New[int, string]()
	m.Insert(1, "value1")

	assert.True(t, m.Has(1))
	assert.False(t, m.Has(2))
}
