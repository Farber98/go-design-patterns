package decorator

import (
	"errors"
	"fmt"
)

/*
DECORATOR:
- Allows you to decorate an already existing type with more functional features without actually touching it.
- Approach similar to matryoshka dolls, where you have a small doll that you can put inside a doll of the same shape but bigger and so on.
- Decorator type impleemnts the same interface of the type it decorates, and stores an instance of that type in its members. This way, you can stack as many decorators (dolls) as you want by simple storing the old decorator ina a filed of the new one.

OBJECTIVE:
- When you need to add funcionality to some code that you don't have access to, or you don't want to modify to avoid a negative effect on the code.
- When you want the functionality of an object to be created or altered dynamically, and the number of features is unknown and could grow fast.

SITUATIONS:
- Extending legacy code without the risk of breaking something.
- Creating types with lots of features based on user inputs, preferences or similar. (Similar to a swiss knife, you have a base type and a lot of unfold functionalitites)

EXAMPLE: Prepare pizza.
- The core is the pizza
- The ingredientes are the decorating types..

ACCEPTANCE CRITERIA: Have a common interface and a core type.
- We must have the main interface that all decorators will implement. This interface will be called IngredientAdd, and it will have the AddIngredient method.
- We must have a core PizzaDecorator type (the decorator) that we will add ingredients to.
- We must have an ingredient "onion" implementing the same IngredientAdd interface that will add the string onion to the returned pizza.
- We must have a ingredient "meat" implementing the IngredientAdd interface that will add the string meat to the returned pizza.
- When calling AddIngredient method on the top object, it must return a fully decorated pizza with the text: Pizza wih the following ingredients: meat, onion
*/

type IngredientAdd interface {
	AddIngredient() (string, error)
}

type PizzaDecorator struct {
	Ingredient IngredientAdd
}

func (p *PizzaDecorator) AddIngredient() (string, error) {
	return "Pizza with the following ingredients:", nil
}

type Onion struct {
	Ingredient IngredientAdd
}

func (o *Onion) AddIngredient() (string, error) {
	if o.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Onion")
	}
	s, err := o.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "onion"), nil
}

type Meat struct {
	Ingredient IngredientAdd
}

func (m *Meat) AddIngredient() (string, error) {
	if m.Ingredient == nil {
		return "", errors.New("An IngredientAdd is needed in the Ingredient field of the Meat")
	}
	s, err := m.Ingredient.AddIngredient()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %s,", s, "meat"), nil
}
