package spentenergy

import (
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
	speed     = 1.39  // средняя скорость в м/с
)

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) float64 {
	// Продолжительность duration перевел в часы
	Hours := duration.Seconds() / 3600
	// Рассчитал и вернул количество калорий.
	return ((walkingCaloriesWeightMultiplier * weight) + (speed*speed/height)*walkingSpeedHeightMultiplier) * Hours * minInH
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	// Рассчитал среднюю скорость с помощью meanSpeed().
	speedrun := MeanSpeed(steps, duration)
	// Рассчитал и вернул количество калорий.
	return ((runningCaloriesMeanSpeedMultiplier * speedrun) - runningCaloriesMeanSpeedShift) * weight
}

func MeanSpeed(steps int, duration time.Duration) float64 {
	// Проверил, что продолжительность duration больше 0.
	if duration <= 0 {
		return 0
	}

	// Вычислил дистанцию с помощью Distance().
	dist := Distance(steps)
	// Вычислил и вернул среднюю скорость.
	hours := duration.Hours()
	return dist / hours
}

func Distance(steps int) float64 {
	// Для вычисления дистанции умножил шаги
	// на длину шага lenStep и разделил на mInKm
	return float64(steps) * lenStep / mInKm
}
