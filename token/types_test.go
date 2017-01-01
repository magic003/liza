package token

import (
	"testing"
)

func TestString(t *testing.T) {
	ty := IDENT
	if ty.String() != "IDENT" {
		t.Errorf("bad string for token type: got %s, expected %s", ty.String(), "IDENT")
	}
}

func TestStringInvalid(t *testing.T) {
	ty := Type(999)
	if ty.String() != "token(999)" {
		t.Errorf("bad string for token type: got %s, expected %s", ty.String(), "token(999)")
	}
}

func TestLookupKeyword(t *testing.T) {
	if LookupKeyword("break") != BREAK {
		t.Errorf("bad token type for lookupKeyword: got %s, expected %s", LookupKeyword("break"), BREAK)
	}

	if LookupKeyword("main") != IDENT {
		t.Errorf("bad token type for lookupKeyword: got %s, expected %s", LookupKeyword("main"), IDENT)
	}
}
