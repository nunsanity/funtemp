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
		{input: "luna", want: []Fact{
			{Name: "Temperatur på Månens overflate om natten", Value: -183, TempType: "C"},
			{Name: "Temperatur på Månens overflate om dagen", Value: 106, TempType: "C"}}},
		{input: "terra", want: []Fact{
			{Name: "Høyeste temperatur målt på Jordens overflate", Value: 56.7, TempType: "C"},
			{Name: "Laveste temperatur målt på Jordens overflate", Value: -89.4, TempType: "C"},
			{Name: "Temperatur i Jordens indre kjerne", Value: 9392, TempType: "K"}}},
	}

	for _, tc := range tests {
		got := GetFunFacts(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
