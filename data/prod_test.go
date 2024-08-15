package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{Name: "Varun", Price: 69, SKU: "abs-abd-abcc-ass"}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
