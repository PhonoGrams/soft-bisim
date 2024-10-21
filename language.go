package soft_bisim

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

// ApplyLanguageSpecificRules applies phonetic transformations based on the detected language.
func ApplyLanguageSpecificRules(name string) string {
	// Detect language using golang.org/x/text
	matcher := language.NewMatcher([]language.Tag{
		language.English, language.German, language.Polish, language.Spanish,
		language.Russian, language.Hebrew,
	})

	// Match the language based on the name
	tag, _, _ := matcher.Match(language.Make(name))
	lang := tag.String()

	// Log the detected language (optional for debugging)
	fmt.Printf("Detected language for '%s': %s (%s)\n", name, display.English.Tags().Name(tag), lang)

	// Apply transformations based on the detected language
	switch lang {
	case "de": // German
		return ApplyGermanPhoneticRules(name)
	case "pl": // Polish
		return ApplyPolishPhoneticRules(name)
	case "es": // Spanish
		return ApplySpanishPhoneticRules(name)
	case "ru": // Russian
		return TransliterateRussianToLatin(name)
	case "he": // Hebrew
		return TransliterateHebrewToLatin(name)
	default:
		// Apply generic phonetic transformation if no specific language is detected
		return PhoneticTransform(name)
	}
}

// ApplyGermanPhoneticRules, ApplyPolishPhoneticRules, ApplySpanishPhoneticRules,
// TransliterateRussianToLatin, TransliterateHebrewToLatin, and PhoneticTransform
// are assumed to be defined elsewhere, as per previous examples.

func main() {
	// Test examples
	name1 := "Schwarz"
	name2 := "Schwartz"

	fmt.Printf("Testing similarity between '%s' and '%s': %v\n", name1, name2, TestPhoneticSimilarity(name1, name2))

	name3 := "Juan"
	name4 := "Huan"

	fmt.Printf("Testing similarity between '%s' and '%s': %v\n", name3, name4, TestPhoneticSimilarity(name3, name4))

	name5 := "Шварц" // Cyrillic for Schwartz
	name6 := "Swarz"

	fmt.Printf("Testing similarity between '%s' and '%s': %v\n", name5, name6, TestPhoneticSimilarity(name5, name6))
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
	} else if strings.ContainsAny(name, "шчцжюя") { // Cyrillic characters for Russian
		return "russian"
	} else if strings.ContainsAny(name, "א") { // Hebrew characters
		return "hebrew"
	}
	// Add more languages and patterns as needed

	return "generic" // Fallback if no specific language pattern is detected
}

// ApplyLanguageSpecificRules applies phonetic transformations based on the detected language.
// This is a simplified example; it should be expanded for each supported language.
func ApplyLanguageSpecificRules(name, language string) string {
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
		return PhoneticTransform(name)
	}
}

// ApplyGermanPhoneticRules applies German-specific phonetic rules
func ApplyGermanPhoneticRules(name string) string {
	// Example transformations specific to German
	replacer := strings.NewReplacer(
		"sch", "sh",
		"ch", "k",
		"tz", "ts",
	)
	return replacer.Replace(name)
}

// ApplyPolishPhoneticRules applies Polish-specific phonetic rules
func ApplyPolishPhoneticRules(name string) string {
	// Example transformations specific to Polish
	replacer := strings.NewReplacer(
		"cz", "ch",
		"sz", "sh",
		"w", "v",
	)
	return replacer.Replace(name)
}

// ApplySpanishPhoneticRules applies Spanish-specific phonetic rules
func ApplySpanishPhoneticRules(name string) string {
	// Example transformations specific to Spanish
	replacer := strings.NewReplacer(
		"ll", "y",
		"ch", "k",
		"ñ", "n",
	)
	return replacer.Replace(name)
}

// TransliterateRussianToLatin converts Russian Cyrillic characters to Latin (simplified)
func TransliterateRussianToLatin(name string) string {
	cyrillicToLatin := map[string]string{
		"ш": "sh", "ч": "ch", "ц": "ts", "ж": "zh",
		"ю": "yu", "я": "ya", "х": "kh", "й": "y",
	}
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
	for hebrew, latin := range hebrewToLatin {
		name = strings.ReplaceAll(name, hebrew, latin)
	}
	return name
}

// ApplyPhoneticSimilarity applies the language-specific rules or generic rules
func ApplyPhoneticSimilarity(name string) string {
	// Detect the language of the name
	language := DetectLanguage(name)

	// Apply language-specific rules or fallback to generic rules
	return ApplyLanguageSpecificRules(name, language)
}

// TestPhoneticSimilarity function compares two names using phonetic similarity rules
func TestPhoneticSimilarity(name1, name2 string) bool {
	// Apply phonetic similarity processing to both names
	phonetic1 := ApplyPhoneticSimilarity(name1)
	phonetic2 := ApplyPhoneticSimilarity(name2)

	// Compare the processed names
	return phonetic1 == phonetic2
}

func main() {
	// Test examples
	name1 := "Schwarz"
	name2 := "Schwartz"

	fmt.Printf("Testing similarity between '%s' and '%s': %v\n", name1, name2, TestPhoneticSimilarity(name1, name2))

	name3 := "Juan"
	name4 := "Huan"

	fmt.Printf("Testing similarity between '%s' and '%s': %v\n", name3, name4, TestPhoneticSimilarity(name3, name4))

	name5 := "Шварц" // Cyrillic for Schwartz
	name6 := "Swarz"

	fmt.Printf("Testing similarity between '%s' and '%s': %v\n", name5, name6, TestPhoneticSimilarity(name5, name6))
}
