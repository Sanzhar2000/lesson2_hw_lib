package lesson2_hw_lib

import (
	"strings"
	"unicode"
)

func ChangeCaseOfCharacters(initial string) string  {

	var resArray []string //Array of strings store char by char

	//Iterate through string and replace uppercase to lowercase and vice versa
	for _, ch := range initial {
		if !unicode.IsLower(ch) && ch != ' ' {
			resArray = append(resArray, strings.ToLower(string(ch)))
		}else {
			resArray = append(resArray, strings.ToUpper(string(ch)))
		}
	}

	//Convert array of strings into one string
	var result string
	result = strings.Join(resArray, "")

	return result
}