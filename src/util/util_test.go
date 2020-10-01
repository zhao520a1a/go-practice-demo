package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UtilTestSuite struct {
	suite.Suite
}

func (u *UtilTestSuite) TestJsonEscape() {
	assert := assert.New(u.T())
	var cases = []struct {
		raw, escaped string
	}{
		{`test double quotes "`, "test double quotes \""},
		{`test braces {`, "test braces {"},
		{`test normal string`, "test normal string"},
	}
	for _, testCase := range cases {
		assert.Equal(testCase.escaped, testCase.raw)
	}
}

func TestUtilTestSuite(t *testing.T) {
	suite.Run(t, new(UtilTestSuite))
}
