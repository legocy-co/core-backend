package helpers

import "testing"

func TestStringToIntArray(t *testing.T) {

	tests := []struct {
		name      string
		input     string
		delimiter string
		want      []int
	}{
		{
			name:      "Test 1",
			input:     "1,2,3,4",
			delimiter: ",",
			want:      []int{1, 2, 3, 4},
		},
		{
			name:      "Test 2",
			input:     "1,2,3,4,5",
			delimiter: ",",
			want:      []int{1, 2, 3, 4, 5},
		},
		// Add empty test case
		{
			name:      "Test 3",
			input:     "",
			delimiter: ",",
			want:      []int{},
		},
		// Add test case with different delimiter
		{
			name:      "Test 4",
			input:     "1;2;3;4;5",
			delimiter: ";",
			want:      []int{1, 2, 3, 4, 5},
		},
		// Add test case with non-integer values
		{
			name:      "Test 5",
			input:     "1,2,3,4,5,a,b,c,d,e",
			delimiter: ",",
			want:      []int{1, 2, 3, 4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringToIntArray(tt.input, tt.delimiter)
			if !compareIntSlices(got, tt.want) {
				t.Errorf("%s: StringToIntArray() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func compareIntSlices(got, want []int) bool {
	if len(got) != len(want) {
		return false
	}

	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}

	return true
}
