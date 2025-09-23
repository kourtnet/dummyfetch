package entities

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var mockRead = func(str string, err error) func(string) ([]byte, error) {
	return func(s string) ([]byte, error) {
		return []byte(str), err
	}
}

var mockOpen = func(str string, errO error) (string, func(string) (*os.File, error)) {
	file, err := os.CreateTemp("./", "mockOpenTmp")
	if err != nil {
		panic(err)
	}

	if _, err := file.WriteString(str); err != nil {
		panic(err)
	}

	file.Seek(0, 0)

	return file.Name(), func(s string) (*os.File, error) {
		return file, errO
	}
}

func Test_getOsReleaseInfo(t *testing.T) {
	oldOpen := open

	testCases := []struct {
		name         string
		fileContents string
		prefix       string
		wantErr      bool
		wantStr      string
	}{
		{
			name:    "open returns error",
			wantErr: true,
		},
		{
			name:         "open returns empty file",
			fileContents: "",
			prefix:       "ID=",
			wantStr:      "",
		},
		{
			name:         "open returns file without required prefix",
			fileContents: `PRETTY_NAME="Mock Linux"`,
			prefix:       "ID=",
			wantStr:      "",
		},
		{
			name:         "open returns file with required prefix",
			fileContents: "PRETTY_NAME=\"Mock Linux\"\nID=\"mock\"",
			prefix:       "ID=",
			wantStr:      "mock",
		},
		{
			name:         "open returns file with required prefix and value without quotes",
			fileContents: "PRETTY_NAME=\"Mock Linux\"\nID=mock",
			prefix:       "ID=",
			wantStr:      "mock",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var errO error
			if tt.wantErr {
				errO = errors.New("error")
			}

			var filepath string
			filepath, open = mockOpen(tt.fileContents, errO)
			defer os.Remove(filepath)

			str, err := getOSReleaseInfo(tt.prefix)

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.wantStr, str)
			}
		})
	}

	open = oldOpen
}

func Test_getKernel(t *testing.T) {
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
			name:    "two trimmed valid numbers",
			readF:   mockRead("100.10 50.10", nil),
			wantErr: false,
			wantStr: "1 min",
		},
		{
			name:    "two trimmed invalid numbers",
			readF:   mockRead("100q.10 50.10", nil),
			wantErr: true,
		},
		{
			name:    "one trimmed invalid number",
			readF:   mockRead("100q.10", nil),
			wantErr: true,
		},
		{
			name:    "one untrimmed valid number",
			readF:   mockRead("60   ", nil),
			wantErr: false,
			wantStr: "1 min",
		},
		{
			name:    "two untrimmed valid numbers",
			readF:   mockRead("60     100", nil),
			wantErr: false,
			wantStr: "1 min",
		},
		{
			name:    "two untrimmed numbers, second is invalid",
			readF:   mockRead("60     10s0", nil),
			wantErr: false,
			wantStr: "1 min",
		},
		{
			name:    "one trimmed valid number representing 1s",
			readF:   mockRead("1", nil),
			wantErr: false,
			wantStr: "0 mins",
		},
		{
			name:    "one trimmed valid number representing 59s",
			readF:   mockRead("59", nil),
			wantErr: false,
			wantStr: "0 mins",
		},
		{
			name:    "one trimmed valid number representing 59s, 99ms",
			readF:   mockRead("59.99", nil),
			wantErr: false,
			wantStr: "0 mins",
		},
		{
			name:    "one trimmed valid number representing 1min",
			readF:   mockRead("60", nil),
			wantErr: false,
			wantStr: "1 min",
		},
		{
			name:    "one trimmed valid number representing 1min, 59s, 99ms",
			readF:   mockRead("119.99", nil),
			wantErr: false,
			wantStr: "1 min",
		},
		{
			name:    "one trimmed valid number representing 2mins",
			readF:   mockRead("120", nil),
			wantErr: false,
			wantStr: "2 mins",
		},
		{
			name:    "one trimmed valid number representing 59 mins, 59s",
			readF:   mockRead("3599", nil),
			wantErr: false,
			wantStr: "59 mins",
		},
		{
			name:    "one trimmed valid number representing 1h",
			readF:   mockRead("3600", nil),
			wantErr: false,
			wantStr: "1 hour",
		},
		{
			name:    "one trimmed valid number representing 2h",
			readF:   mockRead("7200", nil),
			wantErr: false,
			wantStr: "2 hours",
		},
		{
			name:    "one trimmed valid number representing 1 day",
			readF:   mockRead("86400", nil),
			wantErr: false,
			wantStr: "1 day",
		},
		{
			name:    "one trimmed valid number representing 2 days",
			readF:   mockRead("172800", nil),
			wantErr: false,
			wantStr: "2 days",
		},
		{
			name:    "one trimmed valid number representing 2 days, 3 hours, 5 mins, 10s, 5ms",
			readF:   mockRead("183910.05", nil),
			wantErr: false,
			wantStr: "2 days, 3 hours, 5 mins",
		},
		{
			name:    "one trimmed valid number representing 1 day, 1 hour, 1 min",
			readF:   mockRead("90060", nil),
			wantErr: false,
			wantStr: "1 day, 1 hour, 1 min",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			readFile = tt.readF

			str, err := getUptime()

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
