package soft_bisim

import (
	"strings"
)

// PhoneticTransform applies phonetic rules based on simplified and language-independent transformations.
func PhoneticTransform(s string) string {
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

// TransliterateNonLatin transliterates Cyrillic or Hebrew scripts into Latin characters (simplified).
func TransliterateNonLatin(s string) string {
	// For demonstration, we'll use a small set of Cyrillic to Latin mappings
	cyrillicToLatin := map[string]string{
		"ш": "sh", "ч": "ch", "ц": "ts", "ж": "zh",
		"ю": "yu", "я": "ya", "х": "kh", "й": "y",
	}

	for cyrillic, latin := range cyrillicToLatin {
		s = strings.ReplaceAll(s, cyrillic, latin)
	}

	// Add more transliterations for Hebrew or other scripts as needed
	return s
}

// ApplyPhoneticSimilarity applies the phonetic transformation, vowel reduction, and transliteration rules
func ApplyPhoneticSimilarity(name string) string {
	// Step 1: Transliterate if non-Latin script detected (simplified for demo purposes)
	name = TransliterateNonLatin(name)

	// Step 2: Apply the phonetic transformation
	name = PhoneticTransform(name)

	// Step 3: Reduce vowels
	name = VowelReduction(name)

	return name
}
