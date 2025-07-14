package main

import "fmt"

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	data := make(map[int64][]float64)

	// Заполнение мапы
	for i := 0; i < len(temperatures); i++ {
		valueInt := int64(temperatures[i]) - int64(temperatures[i])%10

		slice := data[valueInt]
		if slice == nil {
			data[valueInt] = []float64{}
		}

		data[valueInt] = append(data[valueInt], temperatures[i])
	}

	// Вывод мапы
	for key, value := range data {
		fmt.Printf("%d:", key)
		for i := 0; i < len(value); i++ {
			fmt.Printf(" %.1f", value[i])
		}
		fmt.Println()
	}
}
