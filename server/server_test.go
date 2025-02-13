package server_test

import (
	"distivity/server"
	"distivity/types"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		config types.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server.Run(tt.config)
		})
	}
}
