package conv

import "testing"

func TestFahrenheitToCelsius(t *testing.T) {
    fahrenheit := 68.0
    expectedCelsius := 20.0
    actualCelsius := FahrenheitToCelsius(fahrenheit)

    if actualCelsius != expectedCelsius {
        t.Errorf("FahrenheitToCelsius(%f) = %f; expected %f", fahrenheit, actualCelsius, expectedCelsius)
    }
}

func FahrenheitToCelsius(fahrenheit float64) float64 {
    return (fahrenheit - 32) * 5 / 9
}


func TestCelsiusToKelvin(t *testing.T) {
    celsius := 25.0
    expectedKelvin := 298.15
    actualKelvin := CelsiusToKelvin(celsius)

    if actualKelvin != expectedKelvin {
        t.Errorf("CelsiusToKelvin(%f) = %f; expected %f", celsius, actualKelvin, expectedKelvin)
    }
}

func CelsiusToKelvin(celsius float64) float64 {
    return celsius + 273.15
}
func TestCelsiusToFahrenheit(t *testing.T) {
    // Define some test cases
    testCases := []struct {
        celsius    float64
        fahrenheit float64
    }{
        {0, 32},
        {25, 77},
        {100, 212},
    }

    // Iterate through each test case and check the result
    for _, tc := range testCases {
        // Convert Celsius to Fahrenheit
        result := (tc.celsius * 9 / 5) + 32

        // Check if the result matches the expected Fahrenheit value
        if result != tc.fahrenheit {
            t.Errorf("Celsius %f should be Fahrenheit %f, but got %f", tc.celsius, tc.fahrenheit, result)
        }
    }
}