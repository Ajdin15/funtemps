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
	Celsius = (Farhrenheit - 32)*(5/9)
      // Her skal du implementere funksjonen
	// Du skal ikke formattere float64 i denne funksjonen
	// Gjør formattering i main.go med fmt.Printf eller
	// lignende
	return Celsius = 0
}

func FarhenheitToKelvin(value float64) float64 {
       Kelvin = (Farhenheit - 32) * (5/9) + 32
       return Kelvin = 0
}

func CelsiusToFarhenheit(value float64) float64 {
       Fahrenheit = Celsius*(9/5) + 32
       return Fahrenheit = 0
}

func CelsiusToKelvin(value float64) float64 {
       Kelvin = Celsius + 273.15
       return Kelvin = 0
}

func KelvinToFarhenheit(value float64) float64 {
        Farhrenheit = (Kelvin - 273.15)*(9/5) + 32
        return Fahrenheit = 0
}

func KelvinToCelsius(value float64) float64 {
         Celsius = Kelvin - 273.15
        return Celsius = 0
}

// De andre konverteringsfunksjonene implementere her
// ...
