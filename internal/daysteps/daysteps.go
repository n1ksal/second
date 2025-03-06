package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

const (
	StepLength = 0.65 // Длина шага в метрах
)

// Структура DaySteps
type DaySteps struct {
	Steps                 int
	Duration              time.Duration
	personaldata.Personal // Вложенная структура
}

// Метод Parse() для парсинга строки данных
func (ds *DaySteps) Parse(datastring string) error {
	// Разделил строку на слайс строк
	parts := strings.Split(datastring, ",") // Исправлено: "data" на "datastring"

	// Проверил что длина слайса равна 2
	if len(parts) != 2 {
		return errors.New("неверный формат данных")
	}

	// Преобразовал первый элемент в int (количество шагов)
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось преобразовать количество шагов: %v", err)
	}
	ds.Steps = steps // Сохранил количество шагов в структуру

	// Преобразовал второй элемент в time.Duration
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("не удалось преобразовать продолжительность: %v", err)
	}
	ds.Duration = duration // Сохранил продолжительность в структуру

	return nil
}

// Метод ActionInfo() для получения информации о прогулке
func (ds DaySteps) ActionInfo() string {
	// Вычислил дистанцию
	distance := float64(ds.Steps) * StepLength / 1000
	// Вычислил количество сожженных калорий
	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	// Сформировал и вернул строку с информацией
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, distance, calories)
}
