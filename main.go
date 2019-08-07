package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Encode/decode? Enter [e/d]: \n")
	choice, _ := reader.ReadString('\n')

	if strings.TrimRight(choice, "\n") == "e" {

		fmt.Println("Enter value to encode: ")
		value, _ := reader.ReadString('\n')
		fmt.Println(morseEncode(strings.ToLower(value)))

	} else if strings.TrimRight(choice, "\n") == "d" {

		fmt.Println("Enter value to decode: ")
		value, _ := reader.ReadString('\n')
		output := morseDecode(value)
		fmt.Println(output)

	} else {
		fmt.Println("You have entered an incorrect value, please try again.")
	}
}

func morseEncode(value string) string {

	startArr := strings.Split(value, "")
	endArr := []string{}

	for i := range startArr {

		endArr = append(endArr, morse[startArr[i]]+" ")
	}

	var sb strings.Builder
	for _, str := range endArr {

		sb.WriteString(str)
	}
	return sb.String()
}

func morseDecode(value string) string {

	convertBack := reverseMap(morse)
	startArr := strings.Split(value, " ")
	endArr := []string{}

	for i := range startArr {

		endArr = append(endArr, convertBack[startArr[i]])
	}

	var sb strings.Builder
	for _, str := range endArr {

		sb.WriteString(str)
	}
	return sb.String()
}

func reverseMap(m map[string]string) map[string]string {

	inverted := make(map[string]string)
	for key, value := range m {
		inverted[value] = key
	}
	return inverted
}

var morse = map[string]string{
	"a": ".-",
	"b": "-...",
	"c": "-.-.",
	"d": "-..",
	"e": ".",
	"f": "..-.",
	"g": "--.",
	"h": "....",
	"i": "..",
	"j": ".---",
	"k": "-.-",
	"l": ".-..",
	"m": "--",
	"n": "-.",
	"o": "---",
	"p": ".--.",
	"q": "--.-",
	"r": ".-.",
	"s": "...",
	"t": "-",
	"u": "..-",
	"v": "...-",
	"w": ".--",
	"x": "-..-",
	"y": "-.--",
	"z": "--..",
	"0": "-----",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	" ": "/",
	".": ".-.-.-",
	",": "--..--",
	"?": "..--..",
}
