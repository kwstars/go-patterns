package main

import (
	"reflect"
	"testing"
)

func NewOption(optFuncs ...OptionFunc) (opt *Options) {
	opt = defaultOption
	for _, optFunc := range optFuncs {
		optFunc(opt)
	}
	return
}

func TestOption(t *testing.T) {
	var tests = []struct {
		input    OptionFunc
		excepted Options
	}{
		{input: WithA("AAA"), excepted: Options{A: "AAA", B: "B", C: 100}},
		{input: WithB("BB"), excepted: Options{A: "AAA", B: "BB", C: 100}},
		{input: WithC(11111), excepted: Options{A: "AAA", B: "BB", C: 11111}},
	}

	for _, test := range tests {
		option := NewOption(test.input)
		if !reflect.DeepEqual(*option, test.excepted) {
			t.Errorf("Test Failed: expected: %+v, target: %+v", test.excepted, *option)
		}
	}
}
