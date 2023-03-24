package main

import (
	"flag"
	"fmt"

	"github.com/nunsanity/funtemp/funtemp/conv"
	"github.com/nunsanity/funtemp/funtemp/funfacts"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var ffacts string
var celsius float64
var kelvin float64
var temptype string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/
	flag.Float64Var(&celsius, "C", 0.0, "temperatur i grader celcius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperatur i grader kelvin")
	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	//flag.StringVar(&out, "out", "F", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	//flag.StringVar(&out, "out", "K", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&ffacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises
	flag.StringVar(&temptype, "t", "C", "temperatur i C - celsius, F - farhenheit, K- Kelvin")

}

func main() {

	flag.Parse()
	type FunFact struct {
		name string
		temp float64
	}

	/**
	    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
	    pakkene implementeres.

	    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
	    av flagg. flag-pakken har funksjoner som man kan bruke for å teste
	    hvor mange flagg og argumenter er spesifisert på kommandolinje.

	        fmt.Println("len(flag.Args())", len(flag.Args()))
			    fmt.Println("flag.NFlag()", flag.NFlag())

	    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
	    brukes for å utelukke ugyldige kombinasjoner:
	    -F, -C, -K kan ikke brukes samtidig
	    disse tre kan brukes med -out, men ikke med -funfacts
	    -funfacts kan brukes kun med -t
	    ...
	    Jobb deg gjennom alle tilfellene. Vær obs på at det er en del sjekk
	    implementert i flag-pakken og at den vil skrive ut "Usage" med
	    beskrivelsene av flagg-variablene, som angitt i parameter fire til
	    funksjonene Float64Var og StringVar
	*/

	// Her er noen eksempler du kan bruke i den manuelle testingen
	/*fmt.Println(fahr, out, funfacts)

	fmt.Println("len(flag.Args())", len(flag.Args()))
	fmt.Println("flag.NFlag()", flag.NFlag())

	fmt.Println(isFlagPassed("out"))*/

	// Eksempel på enkel logikk
	if out == "C" && isFlagPassed("F") {
		// Kalle opp funksjonen FahrenheitToCelsius(fahr), som da
		// skal returnere °C
		//fmt.Println("0°F er -17.78°C")
		fmt.Println(fahr, "°F er", conv.FarhenheitToCelsius(fahr), "°C")
	}
	if out == "F" && isFlagPassed("C") {

		fmt.Println(celsius, "°C er", conv.CelsiusToFarhenheit(celsius), "°F")
	}
	if out == "K" && isFlagPassed("C") {

		fmt.Println(celsius, "°C er", conv.CelsiusToKelvin(celsius), "K")
	}
	if out == "C" && isFlagPassed("K") {

		fmt.Println(kelvin, "K er", conv.KelvinToCelsius(kelvin), "°C")
	}
	if out == "F" && isFlagPassed("K") {

		fmt.Println(kelvin, "K er", conv.KelvinToFarhenheit(kelvin), "°F")
	}
	if out == "K" && isFlagPassed("F") {

		fmt.Println(fahr, "°F er", conv.FarhenheitToKelvin(fahr), "K")
	}

	if isFlagPassed("funfacts") && isFlagPassed("t") {
		for _, fact := range funfacts.GetFunFacts(ffacts) {
			if temptype == "C" && fact.TempType == "C" {
				fmt.Printf("%s %.2f %s \n", fact.Name, fact.Value, fact.TempType)
			}
			if temptype == "F" && fact.TempType == "F" {
				fmt.Printf("%s %.2f %s \n", fact.Name, fact.Value, fact.TempType)
			}
			if temptype == "K" && fact.TempType == "K" {
				fmt.Printf("%s %.2f %s \n", fact.Name, fact.Value, fact.TempType)
			}
			if temptype == "C" && fact.TempType == "F" {
				fmt.Println(fact.Name, conv.FarhenheitToCelsius(fact.Value), temptype)
			}
			if temptype == "F" && fact.TempType == "C" {
				fmt.Println(fact.Name, conv.CelsiusToFarhenheit(fact.Value), temptype)
			}
			if temptype == "K" && fact.TempType == "C" {
				fmt.Println(fact.Name, conv.CelsiusToKelvin(fact.Value), temptype)
			}
			if temptype == "C" && fact.TempType == "K" {
				fmt.Println(fact.Name, conv.KelvinToCelsius(fact.Value), temptype)
			}
			if temptype == "F" && fact.TempType == "K" {
				fmt.Println(fact.Name, conv.KelvinToFarhenheit(fact.Value), temptype)
			}
			if temptype == "K" && fact.TempType == "F" {
				fmt.Println(fact.Name, conv.FarhenheitToKelvin(fact.Value), temptype)
			}

		}
	}
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
