package utils

import "testing"

func TestHandleError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name  string
		args  args
		wantB bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := HandleError(tt.args.err); gotB != tt.wantB {
				t.Errorf("HandleError() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}
