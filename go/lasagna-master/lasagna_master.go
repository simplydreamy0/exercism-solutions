package lasagna

func PreparationTime(layers []string, layerAverageTime int) int {
	if layerAverageTime == 0 {
		layerAverageTime = 2;
	}
	return len(layers) * layerAverageTime;
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodleQuanity := 0;
	sauceQuantity := 0.0;
	for _, layer := range layers {
		switch layer {
			case "noodles":
				noodleQuanity += 50;
			case "sauce":
				sauceQuantity += 0.2;
		}
	}
	return noodleQuanity, sauceQuantity;
}

func AddSecretIngredient(friendsList, ownList []string) {
	secretIngredient := friendsList[len(friendsList) -1];
	ownList[len(ownList) -1] = secretIngredient;
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	var scaledQuantities []float64;
	for _, quantity := range quantities {
		scaledQuantities =  append(scaledQuantities, (quantity / 2) * float64(numberOfPortions));
	}
	return scaledQuantities;
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
