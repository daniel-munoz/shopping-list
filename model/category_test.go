package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmtpy(t *testing.T) {
	assert.True(t, Category{}.IsEmpty(), "Category should be empty")
	assert.False(t, Category{Name: "test"}.IsEmpty(), "Category should not be empty")
}
