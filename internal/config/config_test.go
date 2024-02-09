package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mockConfig struct {
	suite.Suite
	FilePath string
}

func TestConfig(t *testing.T) {
	suite.Run(t, new(mockConfig))
}

func (m *mockConfig) SetupTest() {
	m.FilePath = "test/url1000.txt"
}
func (m *mockConfig) TestParse() {

}
