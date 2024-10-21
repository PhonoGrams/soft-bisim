package soft_bisim

// func main() {
// 	data := [][]string{
// 		{"Xanax", "Zanax"},
// 		{"Advil", "Advel"},
// 		{"Tylenol", "Tylanol"},
// 		{"Ibuprofen", "Ibuprufen"},
// 	}

// 	// Example initial weights
// 	weights := map[string]float64{
// 		"match":    0,
// 		"replace":  1.0,
// 		"insert":   1.0,
// 		"delete":   1.0,
// 		"transpos": 0.5,
// 	}

// 	optimizedWeights := GeneticAlgorithm(data, 1000, 50)

// 	fmt.Printf("Optimized Weights: %v\n", optimizedWeights)

// 	// Test similarity using optimized weights
// 	name1 := "Zyrtec"
// 	name2 := "Zantac"
// 	distance := SoftBisimDistance(name1, name2, optimizedWeights)
// 	fmt.Printf("Soft-Bisim Distance between '%s' and '%s': %.2f\n", name1, name2, distance)
// }
