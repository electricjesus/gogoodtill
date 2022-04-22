package util_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/electricjesus/gogoodtill/pkg/util"
)

type jsonTestSuite struct {
	suite.Suite
}

type Subject struct {
	Title string `json:"title"`
}

func (suite *jsonTestSuite) TestJSONMarshalUnmarshal() {
	item := &Subject{Title: "hello json"}
	rdr, err := util.JSONMarshalToReader(item)
	assert.NoError(suite.T(), err)

	target := &Subject{}
	assert.NoError(suite.T(), util.JSONUnmarshalFromReader(rdr, target))
	assert.Equal(suite.T(), item.Title, target.Title)
}

func TestJSONSuitee(t *testing.T) {
	suite.Run(t, &jsonTestSuite{suite.Suite{}})
}
