package sdf

import (
	"encoding/xml"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSdf(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		sdf      Sdf
	}{
		{
			name:     "Test version 1.6",
			fileName: "test_data/root.sdf",
			sdf: Sdf{
				Root: Root{
					Version: "1.6",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bytes, err := ioutil.ReadFile(tt.fileName)
			assert.NoError(t, err)
			sdf := Sdf{}
			err = xml.Unmarshal(bytes, &sdf)
			assert.NoError(t, err)
			assert.Equal(t, tt.sdf.Root.Version, sdf.Root.Version)
		})
	}
}
