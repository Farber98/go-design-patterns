package prototype

import (
	"errors"
	"fmt"
)

/*
PROTOTYPE:
- Have an object that is already created at compilation time, which you can clone as many times as you want at runtime.
- Avoid repetitive object creation.
- Maintain a set of objects that will be cloned to create new instances.
- Provide a default value of some type to start working on top of it.
- Free CPU of complex object initialization to take more memory resources.

EXAMPLE: Small component of customized shirts shop that will have few shirts with their default colors and prices.
- Each shirt will have a Stock Keeping Unit, a system to identify items stored at a specificl location that will need an update.
- To have a shirt cloner object and interface to ask for differenet types of shirts (white, black anmd blue at 15.00, 16.00 and 17.00 dollars respectively)
- When you ask for a white shirt, a clone of the white shirt must be made, and the new instance must be different from the original one.
- The SKU of the created object shouldnt affect new object creation.
- An info method must give me all the information available on the instance fields, including the updated SKU.
*/

type ShirtCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

const (
	WHITE = 1
	BLACK = 2
	BLUE  = 3
)

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: WHITE,
}
var blackPrototype *Shirt = &Shirt{
	Price: 16.00,
	SKU:   "empty",
	Color: BLACK,
}
var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: BLUE,
}

func GetShirtsCloner() ShirtCloner {
	return nil
}

type ShirtsCache struct {
}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case WHITE:
		newItem := *whitePrototype
		return &newItem, nil
	case BLACK:
		newItem := *blackPrototype
		return &newItem, nil
	case BLUE:
		newItem := *bluePrototype
		return &newItem, nil
	default:
		return nil, errors.New("shirt model not recognized")
	}
}

type ItemInfoGetter interface {
	GetInfo() string
}

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("shirt with SKU '%s' and Color id %d that costs %f\n", s.SKU, s.Color, s.Price)
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}
