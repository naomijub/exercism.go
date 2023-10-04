package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time == 0 {
		return len(layers) * 2
	}
	return len(layers) * time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(myFriendsList, myList []string) []string {
	newList := myList[:]
	secretIngredient := myFriendsList[len(myFriendsList)-1]

	newList[len(newList)-1] = secretIngredient
	return newList
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(input []float64, portions int) []float64 {
	scalingFactor := float64(portions) / 2.0

	portionsValues := make([]float64, len(input))
	for index, amount := range input {
		portionsValues[index] = amount * scalingFactor
	}

	return portionsValues
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
