package main

import (
	"fmt"
	"my-go-project/daysteps"
	"my-go-project/spentcalories"
)

func main() {
	fmt.Println("=== Day Action Info ===")
	dayData := "678,0h50m"
	dayInfo := daysteps.DayActionInfo(dayData, 70.0, 1.75)
	fmt.Println(dayInfo)
	
	fmt.Println("\n=== Training Info ===")
	trainingData := "3456,Ходьба,3h00m"
	trainingInfo, err := spentcalories.TrainingInfo(trainingData, 70.0, 1.75)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(trainingInfo)
	}
	
	fmt.Println("\n=== Running Training ===")
	runningData := "5000,Бег,1h30m"
	runningInfo, err := spentcalories.TrainingInfo(runningData, 70.0, 1.75)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(runningInfo)
	}
}