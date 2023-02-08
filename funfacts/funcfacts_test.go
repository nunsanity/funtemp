package funfacts

import (
	"reflect"
	"testing"
)

/*
*

	Mal for TestGetFunFacts funksjonen.
	Definer korrekte typer for input og want,
	og sette korrekte testverdier i slice tests.
*/
func TestGetFunFacts(t *testing.T) {
	type test struct {
		input string
		want  []Fact
	}

	// Her må du legge inn korrekte testverdier
	tests := []test{
		{input: "sun", want: []Fact{
			{Name: "Temperatur i Solens kjerne", Value: 15000000, TempType: "C"},
			{Name: "Temperatur på ytre lav av solen", Value: 5778, TempType: "C"}}},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
