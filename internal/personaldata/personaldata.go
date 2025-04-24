// Package personaldata contains the Personal structure
// and its Print() method for displaying information on the screen
package personaldata

import "fmt"

// Personal contains user data
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print displays the structure data on the screen
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %0.2f кг.\nРост: %0.2f м.\n\n", p.Name, p.Weight, p.Height)
}
