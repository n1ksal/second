package actioninfo

import (
	"fmt"
)

// создал интерфейс DataParser
type DataParser interface {
	Parse(data string) error
	ActionInfo() string
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	// Перебрал все значения слайса dataset в цикле
	for _, data := range dataset {
		// Распарсил каждое значение
		err := dp.Parse(data)
		// Обработал ошибку парсинга
		if err != nil {
			fmt.Printf("Ошибка парсинга данных '%s': %v\n", data, err)
			continue
		}
		// Сформировал и вывел строку с информацией об активности
		info := dp.ActionInfo()
		fmt.Println(info)
	}
}
