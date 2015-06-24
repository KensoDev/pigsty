package pigsty

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestBootstrapParser(t *testing.T) { TestingT(t) }

type BootstrapParserTestSuite struct{}

var _ = Suite(&BootstrapParserTestSuite{})

func (s *BootstrapParserTestSuite) TestBootstrapStepParser(c *C) {
	jsonBlob := []byte(`[
        { "Name": "Some Name", "Path": "s3://some_bucket_name", "Args": ["1", "2", "3"] }
	]`)
	bootstrapActions := NewBootstrapActions(jsonBlob)
	c.Assert(len(bootstrapActions), Equals, 1)
	c.Assert(len(bootstrapActions[0].Args), Equals, 3)
}

func (s *BootstrapParserTestSuite) TestBootstrapActionWithNoArgs(c *C) {
	jsonBlob := []byte(`[
        { "Name": "Some Name", "Path": "s3://some_bucket_name" }
	]`)
	bootstrapActions := NewBootstrapActions(jsonBlob)
	c.Assert(len(bootstrapActions), Equals, 1)
	c.Assert(len(bootstrapActions[0].Args), Equals, 0)
}

func (s *BootstrapParserTestSuite) TestEmptyArrayIfNoJson(c *C) {
	jsonBlob := []byte("")
	bootstrapActions := NewBootstrapActions(jsonBlob)
	c.Assert(len(bootstrapActions), Equals, 0)
}

func (s *BootstrapParserTestSuite) TestReadJsonFileWhenFileNameWrong(c *C) {
	content := ReadJsonFile("some-missing-file-name")
	c.Assert(len(content), Equals, 0)
}
