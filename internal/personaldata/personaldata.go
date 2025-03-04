package personaldata

import "fmt"

// Создал структуру Personal
type Personal struct {
	Name   string  // Имя пользователя
	Weight float64 // Вес пользователя
	Height float64 // Рост пользователя
}

// Ниже создал метод Print()
// Вывел данные структуры на экран
func (p Personal) Print() {
	fmt.Printf("Имя: %s\n", p.Name)
	fmt.Printf("Вес: %.2f\n", p.Weight)
	fmt.Printf("Рост: %.2f\n", p.Height)
}
