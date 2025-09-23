package entities

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockRead = func(str string, err error) func(string) ([]byte, error) {
	return func(s string) ([]byte, error) {
		return []byte(str), err
	}
}

func Test_gerKernel(t *testing.T) {
	oldReadFile := readFile

	testCases := []struct {
		name    string
		readF   func(string) ([]byte, error)
		wantErr bool
		wantStr string
	}{
		{
			name:    "readFile returns error",
			readF:   mockRead("", errors.New("error")),
			wantErr: true,
		},
		{
			name:    "readFile returns 1-word trimmed name",
			readF:   mockRead("mock", nil),
			wantErr: false,
			wantStr: "mock",
		},
		{
			name:    "readFile returns trimmed name with version",
			readF:   mockRead("6.16.7-mock1-1", nil),
			wantErr: false,
			wantStr: "6.16.7-mock1-1",
		},
		{
			name:    "readFile returns untrimmed 1-word name",
			readF:   mockRead("mock \n", nil),
			wantErr: false,
			wantStr: "mock",
		},
		{
			name:    "readFile returns untrimmed name with version",
			readF:   mockRead("6.16.7-mock1-1 \n\t", nil),
			wantErr: false,
			wantStr: "6.16.7-mock1-1",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			readFile = tt.readF

			str, err := getKernel()

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.wantStr, str)
			}
		})
	}

	readFile = oldReadFile
}

func Test_getUptime(t *testing.T) {
	oldReadFile := readFile

	testCases := []struct {
		name    string
		readF   func(string) ([]byte, error)
		wantErr bool
		wantStr string
	}{
		{
			name:    "readFile returns error",
			readF:   mockRead("", errors.New("error")),
			wantErr: true,
		},
		{
			name:    "readFile returns empty slice",
			readF:   mockRead("", nil),
			wantErr: true,
		},
		{
			name:    "readFile sclice of spaces",
			readF:   mockRead("          ", nil),
			wantErr: true,
		},
		{
			name:    "one trimmed valid number representing 1s",
			readF:   mockRead("1", nil),
			wantErr: false,
			wantStr: "mock",
		},
		{
			name:    "readFile returns untrimmed name with version",
			readF:   mockRead("6.16.7-mock1-1 \n\t", nil),
			wantErr: false,
			wantStr: "6.16.7-mock1-1",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			readFile = tt.readF

			str, err := getKernel()

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.wantStr, str)
			}
		})
	}

	readFile = oldReadFile
}
