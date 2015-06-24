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
        { "Path": "s3://some_bucket_name", "Args": ["1", "2", "3"] }
	]`)
	bootstrapActions := NewBootstrapActionsParser(jsonBlob)
	c.Assert(len(bootstrapActions), Equals, 1)
	c.Assert(len(bootstrapActions[0].Args), Equals, 3)
}
