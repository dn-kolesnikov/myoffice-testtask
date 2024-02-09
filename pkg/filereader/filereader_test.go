package filereader

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
)

type mockReader struct {
	suite.Suite
	buffer *bytes.Buffer
	reader *Reader
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(mockReader))
}

func (m *mockReader) SetupTest() {
	m.buffer = &bytes.Buffer{}
	m.reader = New(m.buffer)
}

func (m *mockReader) TestRead() {

	testCases := []struct {
		Description string
		Expected    string
	}{
		{
			Description: "not empty line",
			Expected:    "notEmptyLine",
		},
		{
			Description: "empty line",
			Expected:    " ",
		},
		{
			Description: "nil line",
			Expected:    "",
		},
	}

	for _, tc := range testCases {
		m.buffer.WriteString(tc.Expected + "\n")
	}

	ctx := context.TODO()

	l := m.reader.Read(ctx)

	i := 0

	for v := range l {
		m.Require().Equal(testCases[i].Expected, v)
		i++
	}
}
