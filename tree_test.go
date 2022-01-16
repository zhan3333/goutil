package util_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	util "github.com/zhan3333/goutil"
)

func TestNewTree(t *testing.T) {
	tree := util.NewTree[int](1)
	tree.SetLeft(2)
	tree.SetRight(3)
	assert.Equal(t, 1, tree.Val())
	assert.NotNil(t, tree.Left())
	assert.Equal(t, 2, tree.Left().Val())
	assert.NotNil(t, tree.Right())
	assert.Equal(t, 3, tree.Right().Val())
}
