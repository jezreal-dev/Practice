package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

func aOrAn(nextWord string) string {
    firstLetter := strings.ToLower(string(nextWord[0]))  // get first letter
    vowelsAndH := "aeiouh"                           // what goes here?

    for _, v := range vowelsAndH {
        if string(v) == firstLetter {
            return "an" 
        }
    }
    return "a"
}

func main() {
fmt.Printf("aOrAn(\"apple\") -> %q\n", aOrAn("apple"))// fmt.Printf("aOrAn(\"horse\") -> %q\n", aOrAn("horse"))// fmt.Printf("aOrAn(\"honest\") -> %q\n", aOrAn("honest"))// }
}

// // Example - Input: words = ["this", "is", "so", "exciting"], n = 2
// // Expected Output: ["this", "is", "SO", "EXCITING"]

func uppercaseLastN(words []string, n int) []string {
	start := len(words) - n

	for i := start; i < len(words); i++ {
		words[i] = strings.ToUpper(words[i])
	}
	return words
}

// func main() {
// 	words := []string{"this", "is", "so", "exciting"}
// 	fmt.Printf("uppercaseLastN(..., 2) -> %q\n", uppercaseLastN(words, 4))
// }

// Example 1 - Input: "' awesome '" -> Expected Output: "'awesome'"
// Example 2 - Input: "' hello world '" -> Expected Output: "'hello world'"
func fixSingleQuotes(text string) string {
	re := regexp.MustCompile(`'([^']*)'`)
	match := re.FindStringSubmatch(text)

	if len(match) > 0 {
		inner := strings.TrimSpace(match[1])
		result := re.ReplaceAllString(text, "'"+inner+"'")
		return result
	}
	return text
}

// func main() {
// 	fmt.Printf("fixSingleQuotes(...) -> %q\n", fixSingleQuotes("' awesome '"))
// 	fmt.Printf("fixSingleQuotes(...) -> %q\n", fixSingleQuotes("' hello world '"))
// }


// Example 1 - Input: "," -> Expected Output: true// // Example 2 - Input: "!" -> Expected Output: true// // Example 3 - Input: "x" -> Expected Output: false// 

// func isPunctuation(s string) bool {

// 	punctuations := []string{",", "!", ".", "'"}

// 	for _, p := range punctuations {
// 		if s == p {
// 			return true
// 		}
// 	}
// 	return false
// }

// func main() {
// fmt.Printf("isPunctuation(\",\") -> %v\n", isPunctuation(","))
// fmt.Printf("isPunctuation(\"x\") -> %v\n", isPunctuation("x"))
// }


// Example - Input: ["hello", ",", "world", "!"] -> Expected Output: "hello, world!"
func joinWithPunctuation(tokens []string) string {
	result := "" // Initialize an empty string to build the result
	
	for i, token := range tokens { // Iterate over each token with its index
		if i > 0 && isPunctuation(token) {
			result += token
		} else if i > 0 && !isPunctuation(tokens[i-1]) {
			result += " " + token
		} else {
			result += token
		}
	}
	return result
}

// func main() {
//  tokens := []string{"hello", ",", "world", "!"}
// 	fmt.Printf("joinWithPunctuation(...) -> %q\n", joinWithPunctuation(tokens))
// }

func isPunctuation(s string) bool {
	punctuations := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range punctuations {
		if s == p {
			return true
		}
	}
	return false
}

// Example 1 - Input: "10" -> Expected Output: 2
// Example 2 - Input: "1010" -> Expected Output: 10
// Example 3 - Input: "11111111" -> Expected Output: 255
func binToDecimal(binStr string) (int64, error) {
	result, err := strconv.ParseInt(binStr, 2, 64)
	
	if err != nil {
		return result, err
	}
	return result, err
}

func hexToDecimal(hexStr string) (int64, error) {
	result, err := strconv.ParseInt(hexStr, 16, 64)

	if err != nil {
		return result, err
	}
	return result, err
}


// func main() {
// 	result, err := binToDecimal("10")
// 	fmt.Printf("binToDecimal(\"10\") -> %v, %v\n", result, err)
// 	result, err = binToDecimal("11111111")
// 	fmt.Printf("binToDecimal(\"11111111\") -> %v, %v\n", result, err)
// 	result, err = hexToDecimal("1E")
// 	fmt.Printf("hexToDecimal(\"1E\") -> %v, %v\n", result, err)
// 	result, err = hexToDecimal("FF")
// 	fmt.Printf("hexToDecimal(\"FF\") -> %v, %v\n", result, err)
// }