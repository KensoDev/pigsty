package pigsty

import (
	. "gopkg.in/check.v1"
	"testing"
)

func TestStepParer(t *testing.T) { TestingT(t) }

type StepParserTestSuite struct{}

var _ = Suite(&StepParserTestSuite{})

func (s *StepParserTestSuite) TestReadJsonFileWhenFileNameWrong(c *C) {
	reader := JsonReader{FileName: "steps_example.json"}
	content := reader.ReadFile()
	steps := NewStepsParser(content)
	c.Assert(len(content), Not(Equals), 0)
	c.Assert(len(steps), Equals, 3)
}
