package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"interactive-configurator/internal/scenario"
	"interactive-configurator/internal/writer"
)

const (
	ColorReset = "\033[0m"
	ColorRed   = "\033[31m"
	ColorCyan  = "\033[36m"
)

// clearConsole очищает консоль перед выводом нового шага.
// ANSI-код \033[H\033[2J перемещает курсор в верхний левый угол и очищает экран.
func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

// AskStep — интерактивный опрос одного шага с очисткой экрана.
// После успешной валидации значение сразу записывается в указанный файл.
func AskStep(step scenario.Step, stepNum int) string {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearConsole() // очищаем экран перед отображением шага

		// Выводим информацию о шаге
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

		// Пользователь решил пропустить шаг
		if strings.ToLower(input) == "skip" {
			return ""
		}

		// Валидация введённого значения через тип шага
		if err := step.Type.ValidateValue(input, step.EnumValues); err != nil {
			fmt.Printf("%sОшибка: %s%s\n", ColorRed, err, ColorReset)
			fmt.Println("Попробуйте ещё раз или введите 'skip'.")
			reader.ReadString('\n') // ждем Enter перед повтором
			continue
		}

		// ------------------------------
		// Запись значения в файл
		// ------------------------------
		if err := writer.SaveValue(step.File, step.Key, input); err != nil {
			// Если произошла ошибка при записи (файл не найден, права и т.п.)
			fmt.Printf("%sОшибка записи в файл: %s%s\n", ColorRed, err, ColorReset)
			fmt.Println("Вы можете пропустить шаг или попробовать снова.")
			reader.ReadString('\n')
			continue
		}

		// Если всё прошло успешно — возвращаем значение
		return input
	}
}

// AskScenario — интерактивный ввод всех шагов сценария.
// Использует AskStep для каждого шага и сохраняет результаты в map.
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
