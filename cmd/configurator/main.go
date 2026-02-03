package main

import (
	"fmt"

	"interactive-configurator/internal/input"
	"interactive-configurator/internal/scenario"
)

func main() {
	// Загружаем сценарий
	sc, err := scenario.Load("examples/scenario.json")
	if err != nil {
		fmt.Println("Ошибка загрузки сценария:", err)
		return
	}

	// Проверяем корректность сценария
	if err := sc.Validate(); err != nil {
		fmt.Println("Ошибка сценария:", err)
		return
	}

	fmt.Println("Сценарий загружен.")

	// Интерактивный ввод
	responses := input.AskScenario(sc)

	fmt.Println("Ввод завершён. Результаты:")
	for k, v := range responses {
		fmt.Printf("  %s = %s\n", k, v)
	}
}
