package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	for _, c := range []struct {
		name    string
		testing bool
		answer  string
	}{
		{
			name:    "Test",
			testing: true,
			answer:  "123",
		},
		{
			name:    "Quiz",
			testing: false,
			answer:  "5502",
		},
	} {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.answer, run(c.testing))
		})
	}
}
