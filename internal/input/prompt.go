package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"interactive-configurator/internal/scenario"
	"interactive-configurator/internal/writer"
)

// AskStep — интерактивный опрос одного шага.
// Линейный вывод, без очистки экрана и без цветов.
// После успешной валидации значение записывается в указанный файл.
func AskStep(step scenario.Step, stepNum int) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		// Линейный вывод шага
		fmt.Printf("Шаг %d\n", stepNum)
		fmt.Printf("Файл       : %s\n", step.File)
		fmt.Printf("Ключ       : %s\n", step.Key)
		fmt.Printf("Тип        : %s\n", step.Type)
		if step.Comment != "" {
			fmt.Printf("Комментарий: %s\n", step.Comment)
		}
		if step.Default != nil {
			fmt.Printf("Значение по умолчанию: %s\n", *step.Default)
		}

		// Ввод пользователя
		fmt.Print("Введите значение (или 'skip' чтобы пропустить): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Пропуск шага
		if strings.ToLower(input) == "skip" {
			fmt.Println("Шаг пропущен.")
			return ""
		}

		// Валидация значения
		if err := step.Type.ValidateValue(input, step.EnumValues); err != nil {
			fmt.Printf("Ошибка: %s\n", err)
			fmt.Printf("Попробуйте ещё раз или введите 'skip'.\n\n")
			continue
		}

		// Запись в файл
		if err := writer.SaveValue(step.File, step.Key, input); err != nil {
			fmt.Printf("Ошибка записи в файл: %s\n", err)
			fmt.Println("Попробуйте ещё раз или введите 'skip'.")
			continue
		}

		// Успешный ввод
		fmt.Println("Значение сохранено.")
		return input
	}
}

// AskScenario — последовательный интерактивный ввод всех шагов сценария.
// Возвращает map ключ => введённое значение.
func AskScenario(sc *scenario.Scenario) map[string]string {
	responses := make(map[string]string)

	for i, step := range sc.Steps {
		val := AskStep(step, i+1)
		if val != "" {
			responses[step.Key] = val
		}
	}

	return responses
}
