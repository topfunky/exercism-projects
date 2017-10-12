package igpay

import "testing"

const targetTestVersion = 1

var tests = []struct{ pl, in string }{
	{"appleay", "apple"},
	{"earay", "ear"},
	{"igpay", "pig"},
	{"oalakay", "koala"},
	{"airchay", "chair"},
	{"eenquay", "queen"},
	{"aresquay", "square"},
	{"erapythay", "therapy"},
	{"ushthray", "thrush"},
	{"oolschay", "school"},
	{"ickquay astfay unray", "quick fast run"},
	{"ellowyay", "yellow"},
	{"yttriaay", "yttria"},
	{"enonxay", "xenon"},
	{"xrayay", "xray"},
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestPigLatin(t *testing.T) {
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			if pl := PigLatin(test.in); pl != test.pl {
				t.Fatalf("PigLatin(%q) = %q, want %q.", test.in, pl, test.pl)
			}
		})
	}
}

func BenchmarkPigLatin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PigLatin("school xray things all finger alligator bubble gum")
	}
}
