package soft_bisim

import (
	"math"
	"math/rand"
	"strings"
	"unicode"
)

// ComputeBigrams generates bigrams from a given string, working with runes.
func ComputeBigrams(s string) []string {
	var bigrams []string
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		bigram := string(runes[i : i+2])
		bigrams = append(bigrams, bigram)
	}
	return bigrams
}

// DetectLanguage determines the language of the name based on specific patterns.
// This is a simplified implementation and should be expanded with more robust logic.
func DetectLanguage(name string) string {
	name = strings.ToLower(name)

	// Basic detection rules based on common character patterns
	if strings.Contains(name, "sch") || strings.HasSuffix(name, "mann") {
		return "german"
	} else if strings.Contains(name, "cz") || strings.Contains(name, "sz") {
		return "polish"
	} else if strings.HasPrefix(name, "ch") || strings.HasSuffix(name, "ez") {
		return "spanish"
	} else if containsCyrillic(name) {
		return "russian"
	} else if containsHebrew(name) {
		return "hebrew"
	}
	// Add more languages and patterns as needed

	return "generic" // Fallback if no specific language pattern is detected
}

// Helper function to check for Cyrillic characters
func containsCyrillic(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Cyrillic) {
			return true
		}
	}
	return false
}

// Helper function to check for Hebrew characters
func containsHebrew(s string) bool {
	for _, r := range s {
		if unicode.In(r, unicode.Hebrew) {
			return true
		}
	}
	return false
}

// ApplyPhoneticTransform applies phonetic transformations based on the detected language.
func ApplyPhoneticTransform(name string) string {
	language := DetectLanguage(name)

	switch language {
	case "german":
		return ApplyGermanPhoneticRules(name)
	case "polish":
		return ApplyPolishPhoneticRules(name)
	case "spanish":
		return ApplySpanishPhoneticRules(name)
	case "russian":
		return TransliterateRussianToLatin(name)
	case "hebrew":
		return TransliterateHebrewToLatin(name)
	default:
		// Apply generic phonetic transformation if no specific language is detected
		return GenericPhoneticTransform(name)
	}
}

// GenericPhoneticTransform applies language-independent phonetic transformations.
func GenericPhoneticTransform(s string) string {
	// Convert to lowercase for uniformity
	s = strings.ToLower(s)

	// Apply basic phonetic rules common across languages
	replacer := strings.NewReplacer(
		"ph", "f",
		"ck", "k",
		"gh", "g",
		"sch", "sh",
		"ch", "k",
		"th", "t",
		"sh", "s",
		"cz", "c",
		"qu", "k",
		"gn", "n",
		"wr", "r",
		"kn", "n",
		"wh", "w",
		"dg", "g",
	)
	s = replacer.Replace(s)

	// Remove silent consonants
	s = strings.ReplaceAll(s, "h", "")

	// Simplify double letters (e.g., tt -> t, ll -> l)
	for _, ch := range []string{"tt", "ll", "ss", "pp", "rr", "mm"} {
		s = strings.ReplaceAll(s, ch, string(ch[0]))
	}

	return s
}

// ApplyGermanPhoneticRules applies German-specific phonetic rules
func ApplyGermanPhoneticRules(name string) string {
	// Convert to lowercase
	name = strings.ToLower(name)

	// Example transformations specific to German
	replacer := strings.NewReplacer(
		"sch", "sh",
		"ch", "k",
		"tz", "ts",
		"z", "ts",
		"ss", "s",
		"eu", "oy",
		"ä", "ae",
		"ö", "oe",
		"ü", "ue",
		"ß", "ss",
	)
	return replacer.Replace(name)
}

// ApplyPolishPhoneticRules applies Polish-specific phonetic rules
func ApplyPolishPhoneticRules(name string) string {
	// Convert to lowercase
	name = strings.ToLower(name)

	// Example transformations specific to Polish
	replacer := strings.NewReplacer(
		"cz", "ch",
		"sz", "sh",
		"w", "v",
		"ł", "l",
		"ń", "n",
		"ś", "s",
		"ź", "z",
		"ż", "z",
		"ą", "a",
		"ę", "e",
	)
	return replacer.Replace(name)
}

// ApplySpanishPhoneticRules applies Spanish-specific phonetic rules
func ApplySpanishPhoneticRules(name string) string {
	// Convert to lowercase
	name = strings.ToLower(name)

	// Example transformations specific to Spanish
	replacer := strings.NewReplacer(
		"ll", "y",
		"ch", "k",
		"ñ", "n",
		"v", "b",
		"ce", "se",
		"ci", "si",
		"z", "s",
		"j", "h",
		"h", "",
	)
	return replacer.Replace(name)
}

// TransliterateRussianToLatin converts Russian Cyrillic characters to Latin (simplified)
func TransliterateRussianToLatin(name string) string {
	cyrillicToLatin := map[string]string{
		"ш": "sh", "ч": "ch", "ц": "ts", "ж": "zh",
		"ю": "yu", "я": "ya", "х": "kh", "й": "y",
		"ё": "yo", "э": "e", "ы": "i", "щ": "shch",
		"ъ": "", "ь": "",
		"а": "a", "б": "b", "в": "v", "г": "g",
		"д": "d", "е": "e", "з": "z", "и": "i",
		"к": "k", "л": "l", "м": "m", "н": "n",
		"о": "o", "п": "p", "р": "r", "с": "s",
		"т": "t", "у": "u", "ф": "f",
	}

	// Convert to lowercase
	name = strings.ToLower(name)

	// Replace Cyrillic characters with Latin equivalents
	for cyrillic, latin := range cyrillicToLatin {
		name = strings.ReplaceAll(name, cyrillic, latin)
	}
	return name
}

// TransliterateHebrewToLatin converts Hebrew characters to Latin (simplified)
func TransliterateHebrewToLatin(name string) string {
	hebrewToLatin := map[string]string{
		"א": "a", "ב": "b", "ג": "g", "ד": "d",
		"ה": "h", "ו": "v", "ז": "z", "ח": "ch",
		"ט": "t", "י": "y", "כ": "k", "ל": "l",
		"מ": "m", "נ": "n", "ס": "s", "ע": "a",
		"פ": "p", "צ": "tz", "ק": "k", "ר": "r",
		"ש": "sh", "ת": "t",
	}

	// Replace Hebrew characters with Latin equivalents
	for hebrew, latin := range hebrewToLatin {
		name = strings.ReplaceAll(name, hebrew, latin)
	}
	return name
}

// VowelReduction removes or reduces vowels to improve phonetic matching
func VowelReduction(s string) string {
	// Ignore vowels except for the first character
	if len(s) > 1 {
		vowels := "aeiou"
		reduced := string(s[0]) // Always keep the first character
		for _, char := range s[1:] {
			if !strings.ContainsRune(vowels, char) {
				reduced += string(char)
			}
		}
		return reduced
	}
	return s
}

// SoftBisimDistance computes the Soft-Bisim distance between two strings.
func SoftBisimDistance(s1, s2 string, weights Weights) float64 {
	// Apply phonetic transformation
	s1 = ApplyPhoneticTransform(s1)
	s2 = ApplyPhoneticTransform(s2)

	// Optionally apply vowel reduction
	s1 = VowelReduction(s1)
	s2 = VowelReduction(s2)

	bigrams1 := ComputeBigrams(s1)
	bigrams2 := ComputeBigrams(s2)

	// Distance matrix
	dist := make([][]float64, len(bigrams1)+1)
	for i := range dist {
		dist[i] = make([]float64, len(bigrams2)+1)
	}

	// Initialize base cases
	for i := 0; i <= len(bigrams1); i++ {
		dist[i][0] = float64(i) * weights.Delete
	}
	for j := 0; j <= len(bigrams2); j++ {
		dist[0][j] = float64(j) * weights.Insert
	}

	// Fill the distance matrix using Soft-Bisim rules
	for i := 1; i <= len(bigrams1); i++ {
		for j := 1; j <= len(bigrams2); j++ {
			if bigrams1[i-1] == bigrams2[j-1] {
				// Case 1: Exact match
				dist[i][j] = dist[i-1][j-1] + weights.Match
			} else {
				// Case 2: Substitution (including case change and phonetic change)
				cost := weights.Replace
				if strings.ToLower(bigrams1[i-1]) == strings.ToLower(bigrams2[j-1]) {
					cost = weights.CaseChange
				}
				// You can add conditions for phonetic change if applicable

				dist[i][j] = math.Min(
					math.Min(dist[i-1][j]+weights.Delete, dist[i][j-1]+weights.Insert),
					dist[i-1][j-1]+cost,
				)

				// Case 3: Transposition
				if i > 1 && j > 1 && bigrams1[i-1] == bigrams2[j-2] && bigrams1[i-2] == bigrams2[j-1] {
					dist[i][j] = math.Min(dist[i][j], dist[i-2][j-2]+weights.Transposition)
				}

				// Case 4: Merge and Split (not fully implemented; placeholder for future extension)
				// You can implement merge and split operations if necessary
			}
		}
	}

	return dist[len(bigrams1)][len(bigrams2)]
}

// FitnessFunction evaluates the performance of the weights using the F-measure
func FitnessFunction(data [][]string, weights Weights) float64 {
	var totalScore float64
	for _, pair := range data {
		name1, name2 := pair[0], pair[1]
		// Calculate similarity
		distance := SoftBisimDistance(name1, name2, weights)
		len1 := len([]rune(name1))
		len2 := len([]rune(name2))
		maxLen := float64(len1)
		if len2 > len1 {
			maxLen = float64(len2)
		}
		similarity := 1 - (distance / maxLen)
		totalScore += similarity
	}
	return totalScore / float64(len(data))
}

// GeneticAlgorithm optimizes the weights used in the Soft-Bisim calculation
func GeneticAlgorithm(data [][]string, iterations int, populationSize int) Weights {
	// Initialize population of weights
	population := make([]Weights, populationSize)
	for i := range population {
		population[i] = Weights{
			Match:          rand.Float64(),
			Replace:        rand.Float64(),
			Insert:         rand.Float64(),
			Delete:         rand.Float64(),
			Transposition:  rand.Float64(),
			Merge:          rand.Float64(),
			Split:          rand.Float64(),
			CaseChange:     rand.Float64(),
			PhoneticChange: rand.Float64(),
		}
	}

	bestWeights := population[0]
	bestScore := FitnessFunction(data, bestWeights)

	// Evolution process
	for iteration := 0; iteration < iterations; iteration++ {
		for i := 0; i < populationSize; i++ {
			weights := population[i]
			score := FitnessFunction(data, weights)
			if score > bestScore {
				bestScore = score
				bestWeights = weights
			}

			// Mutation process
			if rand.Float64() < 0.1 { // Mutation probability
				weights.Match = rand.Float64()
				weights.Replace = rand.Float64()
				weights.Insert = rand.Float64()
				weights.Delete = rand.Float64()
				weights.Transposition = rand.Float64()
				weights.Merge = rand.Float64()
				weights.Split = rand.Float64()
				weights.CaseChange = rand.Float64()
				weights.PhoneticChange = rand.Float64()
				population[i] = weights
			}
		}
	}
	return bestWeights
}
