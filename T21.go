package main

import (
	"fmt"
)

// Существующий класс/структура
type CelsiusSensor struct{}

func (c *CelsiusSensor) GetTemperatureC() float64 {
	return 25.0
}

// Целевой интерфейс, который ожидает клиент
type FahrenheitSensor interface {
	GetTemperatureF() float64
}

// Адаптер: реализует целевой интерфейс и использует существующий объект
type CelsiusToFahrenheitAdapter struct {
	celsiusSensor *CelsiusSensor
}

func (a *CelsiusToFahrenheitAdapter) GetTemperatureF() float64 {
	return a.celsiusSensor.GetTemperatureC()*9/5 + 32
}

func displayTemperature(sensor FahrenheitSensor) {
	fmt.Printf("Temperature: %.2f°F\n", sensor.GetTemperatureF())
}

func main() {
	cSensor := &CelsiusSensor{}

	adapter := &CelsiusToFahrenheitAdapter{celsiusSensor: cSensor}

	displayTemperature(adapter)
}
