package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "dan",
		Price: 1.00,
		SKU:   "abcd-abcd-abcd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
