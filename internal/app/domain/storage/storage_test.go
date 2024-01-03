package storage

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemStorage(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  string
	}{
		{
			name: "test1",
			key:  "metr1",
		},
	}
	Stor := GetStorage()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, Stor.SetRelation(tt.key, tt.val))
		})
	}
}
