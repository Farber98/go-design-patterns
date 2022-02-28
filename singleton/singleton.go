package singleton

/*
SINGLETON:
- Provides a single instance of an object. Guarantees there are no duplicates.
- At first call to use the instance, it is created.
- Subsequent calls are reused between all parts in the app that need to use that behavior.

WHEN TO USE IT:
- We need single, shared value, of some particular type.
- We need to restrict object creation of some type to a single unit along entire program.

SITUATIONS:
- Use same connection (eg. SSH, DB)
- Limit the access/number of calls to some place.

EXAMPLE: unique counter
- When no counter has been created before, a new one is created with value 0.
- If a counter has already been created, return this instance that holds the actual count.
- If we call the method AddOne, the count must be incremented by 1.
*/

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

// Go can't initialize a struct as nil.
// Need to define a pointer to a struct of type Singleton as nil.
var instance *singleton

func GetInstance() Singleton {
	if instance == nil {
		// Create pointer to an instance of type singleton
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
