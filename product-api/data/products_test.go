package data

import "testing"

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "iphone",
		Price: 29.91,
		SKU:   "a-a-a",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
