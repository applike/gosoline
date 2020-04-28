package cfg_test

import (
	"github.com/applike/gosoline/pkg/cfg"
	"github.com/stretchr/testify/suite"
	"testing"
)

type MergeTestSuite struct {
	suite.Suite
}

func (s *MergeTestSuite) TestMerge() {
	type Embedded struct {
		F32 float32
	}

	type object struct {
		Embedded
		I  int
		S1 string
		S2 string
	}

	a := object{
		I:  1,
		S1: "string",
		S2: "foo",
		Embedded: Embedded{
			F32: 1.1,
		},
	}
	b := object{
		I:  2,
		S1: "",
		S2: "bar",
		Embedded: Embedded{
			F32: 1.2,
		},
	}

	err := cfg.Merge(&a, b)
	s.NoError(err, "there should be no error on merge")

	s.Equal(2, a.I)
	s.Equal("string", a.S1)
	s.Equal("bar", a.S2)
	s.Equal(float32(1.2), a.F32)
}

func TestMergeTestSuite(t *testing.T) {
	suite.Run(t, new(MergeTestSuite))
}
