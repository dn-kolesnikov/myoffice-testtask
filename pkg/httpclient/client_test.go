package httpclient

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mockClient struct {
	suite.Suite
	c *Client
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(mockClient))
}

func (m *mockClient) SetupTest() {
	m.c = New()
}

func (m *mockClient) TestGet() {
	r, err := m.c.Get("http://myoffice.ru")
	m.Require().NoError(err)
	m.Assert().Less(r.HandleTime(), _defaultTimeout, "handle time should be less than default timeout")
	m.Assert().Greater(r.ContentLength(), 0, "content length should be greater than 0")
}

func (m *mockClient) TestValidateURL() {
	testCases := []struct {
		Description string
		URL         string
		ExpectedErr bool
	}{
		{
			Description: "url with valid url",
			URL:         "http://myoffice.ru",
			ExpectedErr: false,
		},
		{
			Description: "url with invalid scheme",
			URL:         "the://myoffice/ru",
			ExpectedErr: true,
		},
		{
			Description: "url without scheme",
			URL:         "myoffice.ru",
			ExpectedErr: true,
		},
		{
			Description: "nil url",
			URL:         "",
			ExpectedErr: true,
		},
		{
			Description: "empty url",
			URL:         " ",
			ExpectedErr: true,
		},
	}
	for _, tc := range testCases {
		tc := tc
		m.T().Run(tc.Description, func(t *testing.T) {
			t.Parallel()
			err := m.c.ValidateURL(tc.URL)
			m.Require().Equal(tc.ExpectedErr, err != nil)
		})
	}
}
