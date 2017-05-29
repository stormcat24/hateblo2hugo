package service

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestMovableTypeImpl_Parse(t *testing.T) {

	s := MovableTypeImpl{}
	entries, err := s.Parse("movabletype_test.txt")
	require.NoError(t, err)

	require.Equal(t, 1, len(entries))
}
