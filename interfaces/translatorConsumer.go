package interfaces

/**
Enhance
How to:
* overload a method with the same logic
* A different way to use polimorphism
* another way to think interfaces in Java
*/

func PrepareGreeting(b Bot) string {
	s := "Hey!" + b.GetGreeting()
	return s
}
