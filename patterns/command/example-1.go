package command

import "fmt"

type Command interface {
	execute()
}

// the restaurant contains the total dishes and the total cleaned dishes
type Restaurant struct {
	TotalDishes   int
	CleanedDishes int
}

func (r *Restaurant) MakePizza(n int) Command {
	return &MakePizzaCommand{
		n:          n,
		restaurant: r,
	}
}

func (r *Restaurant) MakeSalad(n int) Command {
	return &MakeSaladCommand{
		n:          n,
		restaurant: r,
	}
}

func (r *Restaurant) CleanDishes() Command {
	return &CleanDishesCommand{restaurant: r}
}

type MakePizzaCommand struct {
	n          int
	restaurant *Restaurant
}

func (mpc *MakePizzaCommand) execute() {
	mpc.restaurant.CleanedDishes -= mpc.n
	fmt.Println("made", mpc.n, "pizzas")
}

type MakeSaladCommand struct {
	n          int
	restaurant *Restaurant
}

func (msc *MakeSaladCommand) execute() {
	msc.restaurant.CleanedDishes -= msc.n
	fmt.Println("made", msc.n, "salads")
}

type CleanDishesCommand struct{ restaurant *Restaurant }

func (clc *CleanDishesCommand) execute() {
	clc.restaurant.CleanedDishes = clc.restaurant.TotalDishes
	fmt.Println("all dishes cleaned")
}

type Cook struct {
	Commands []Command
}

func (c *Cook) executeCommands() {
	for _, c := range c.Commands {
		c.execute()
	}
}

func NewRestaurant() *Restaurant {
	const totalDishes = 10
	r := &Restaurant{
		totalDishes,
		totalDishes,
	}
	return r
}

func Example1() {
	r := NewRestaurant()

	tasks := []Command{
		r.MakePizza(2),
		r.MakeSalad(1),
		r.MakePizza(3),
		r.CleanDishes(),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	// create the cooks that will execute the tasks
	cooks := []*Cook{
		&Cook{},
		&Cook{},
	}

	// assigning tasks to cooks alternating between the existing cooks
	for i, task := range tasks {
		// using the modulus of the current task index, we can alternate between different cooks
		cook := cooks[i%len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommands()
	}
}
