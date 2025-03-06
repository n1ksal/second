package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/personaldata"
	"github.com/Yandex-Practicum/go1fl-sprint5-final/internal/spentenergy"
)

// создал структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создал метод Parse()
func (t *Training) Parse(datastring string) (err error) {
	// разделил строку на части и возвратил ошибку формата данных
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("неправильный формат данных")
	}

	//Преобразовал количество шагов и возвращаю ошибку
	t.Steps, err = strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return fmt.Errorf("ошибка преобразования количества шагов: %v", err)
	}

	// Проверил второй элемент слайса на соответствие известных нам типов
	trainingType := strings.TrimSpace(parts[1])
	if trainingType != "Бег" && trainingType != "Ходьба" {
		return errors.New("неизвестный тип тренировки")
	}
	t.TrainingType = trainingType

	// Преобразовал третий элемент слайса
	t.Duration, err = time.ParseDuration(strings.TrimSpace(parts[2]))
	if err != nil {
		return fmt.Errorf("ошибка преобразования длительности тренировки: %v", err)
	}

	return nil
}

// создал метод ActionInfo()
func (t Training) ActionInfo() string {
	// Вычислил дистанцию
	distance := spentenergy.Distance(t.Steps)

	// Вычислил среднюю скорость
	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64

	// расчитал калории в зависимости от типа тренировки
	switch t.TrainingType {
	case "Бег":
		calories = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)

	case "Ходьба":
		calories = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)

	default:
		return "неизвестный тип тренировки"
	}

	// создал строку с информацией о тренировке
	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), distance, speed, calories)

}
