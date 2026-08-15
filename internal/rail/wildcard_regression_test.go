package rail

import "testing"

func TestParseRangeAppliesOperatorToWildcardBounds(t *testing.T) {
	tests := []struct {
		constraint string
		version    string
		want       bool
	}{
		{"<1.x", "1.5.0", false},
		{"<1.x", "0.9.0", true},
		{"<=1.x", "1.5.0", true},
		{">1.x", "1.5.0", false},
		{">1.x", "2.0.0", true},
		{">=1.x", "1.0.0", true},
		{"1.x", "1.5.0", true},
		{"1.x", "2.0.0", false},
	}
	for _, tt := range tests {
		constraint, err := ParseRange(tt.constraint)
		if err != nil {
			t.Fatalf("ParseRange(%q): %v", tt.constraint, err)
		}
		version, err := ParseVersion(tt.version)
		if err != nil {
			t.Fatalf("ParseVersion(%q): %v", tt.version, err)
		}
		if got := constraint.Contains(version); got != tt.want {
			t.Fatalf("ParseRange(%q).Contains(%s) = %t, want %t", tt.constraint, tt.version, got, tt.want)
		}
	}
}
