package config_test

import (
	"distivity/config"
	"distivity/types"
	"testing"
)

func TestGetConfig(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want types.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.GetConfig()
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
