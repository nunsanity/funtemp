package funfacts

/*
*

	Implementer funfacts-funksjon:
	  GetFunFacts(about string) []string
	    hvor about kan ha en av tre testverdier, -
	      sun, luna eller terra

	Sett inn alle Funfucts i en struktur
	type FunFacts struct {
	    Sun []string
	    Luna []string
	    Terra []string
	}
*/

type Fact struct {
	Name     string
	Value    float64
	TempType string
}

type FunFacts struct {
	Sun   []Fact
	Luna  []Fact
	Terra []Fact
}

func GetFunFacts(about string) []Fact {

	funfacts := FunFacts{
		Sun: []Fact{
			{Name: "Temperatur i Solens kjerne", Value: 15000000, TempType: "C"},
			{Name: "Temperatur på ytre lav av solen", Value: 5778, TempType: "C"},
		},
		Luna: []Fact{
			{Name: "Temperatur på Månens overflate om natten", Value: -183, TempType: "C"},
			{Name: "Temperatur på Månens overflate om dagen", Value: 106, TempType: "C"},
		},
		Terra: []Fact{
			{Name: "Høyeste temperatur målt på Jordens overflate", Value: 56.7, TempType: "C"},
			{Name: "Laveste temperatur målt på Jordens overflate", Value: -89.4, TempType: "C"},
			{Name: "Temperatur i Jordens indre kjerne", Value: 9392, TempType: "K"},
		},
	}
	if about == "sun" {
		return funfacts.Sun
	} else if about == "luna" {
		return funfacts.Luna
	} else if about == "terra" {
		return funfacts.Terra
	} else {
		return []Fact{}
	}
}
