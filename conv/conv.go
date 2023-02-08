package conv

/*
  I denne pakken skal alle konverteringfunksjonene
  implementeres. Bruk engelsk.
    FarhenheitToCelsius
    CelsiusToFahrenheit
    KelvinToFarhenheit
    ...
*/

// Konverterer Farhenheit til Celsius
func FarhenheitToCelsius(value float64) float64 {
	// Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return (value - 32) * 5 / 9

}

// Konverterer Celsius til Farhenheit
func CelsiusToFarhenheit(value float64) float64 {

	return (value * 9 / 5) + 32
}

// Konverterer Kelvin til Farhenheit
func KelvinToFarhenheit(value float64) float64 {

	//return (value-273.15)*(9/5) + 32
	return 1.8*value - 459.67
}

// Konverterer Kelvin til Celsius
func KelvinToCelsius(value float64) float64 {

	return value - 273.15
}

// Konverterer Celsius til Kelvin
func CelsiusToKelvin(value float64) float64 {

	return value + 273.15
}

// Konverterer Farhenheit til Kelvin
func FarhenheitToKelvin(value float64) float64 {

	return (value-32)*5/9 + 273.15
}
