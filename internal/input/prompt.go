package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"interactive-configurator/internal/scenario"
)

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
	ColorGreen = "\033[32m"
	ColorCyan  = "\033[36m"
)

// AskStep — интерактивный опрос одного шага
func AskStep(step scenario.Step, stepNum int) string {
	reader := bufio.NewReader(os.Stdin)
	for {
		// Заголовок шага
		fmt.Printf("%sШаг %d%s\n", ColorCyan, stepNum, ColorReset)
		fmt.Printf("%sФайл   :%s %s\n", ColorCyan, ColorReset, step.File)
		fmt.Printf("%sКлюч    :%s %s\n", ColorCyan, ColorReset, step.Key)
		fmt.Printf("%sТип     :%s %s\n", ColorCyan, ColorReset, step.Type)
		if step.Comment != "" {
			fmt.Printf("%sКомментарий:%s %s\n", ColorCyan, ColorReset, step.Comment)
		}
		if step.Default != nil {
			fmt.Printf("%sЗначение по умолчанию:%s %s\n", ColorCyan, ColorReset, *step.Default)
		}

		// Ввод пользователя
		fmt.Print("Введите значение (или 'skip' чтобы пропустить): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "skip" {
			fmt.Println(ColorGreen + "Шаг пропущен.\n" + ColorReset)
			return ""
		}

		// Проверка значения
		if err := step.Type.ValidateValue(input, step.EnumValues); err != nil {
			fmt.Printf("%sОшибка:%s %s\n", ColorRed, ColorReset, err)
			fmt.Println("Попробуйте ещё раз или введите 'skip' чтобы пропустить.\n")
			continue
		}

		fmt.Printf("%sПринято:%s %s\n\n", ColorGreen, ColorReset, input)
		return input
	}
}

// AskScenario — интерактивный ввод по всем шагам сценария
func AskScenario(sc *scenario.Scenario) map[string]string {
	responses := make(map[string]string)
	fmt.Println("Начинаем интерактивный ввод...\n")

	for i, step := range sc.Steps {
		val := AskStep(step, i+1)
		if val != "" {
			responses[step.Key] = val
		}
	}
	return responses
}
