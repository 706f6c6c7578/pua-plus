package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/golang/freetype"
)

var letterToPUA = make(map[rune]rune)
var puaToLetter = make(map[rune]rune)

func init() {
	puaMap := map[string]string{
		"A": "U+F0147",
		"B": "U+F00AF",
		"C": "U+F0365",
		"D": "U+F0081",
		"E": "U+F02E9",
		"F": "U+F0073",
		"G": "U+F042D",
		"H": "U+F00C5",
		"I": "U+F00F7",
		"J": "U+F0100",
		"K": "U+F04E6",
		"L": "U+F0068",
		"M": "U+F02A5",
		"N": "U+F0097",
		"O": "U+F0252",
		"P": "U+F0110",
		"Q": "U+F0202",
		"R": "U+F01D0",
		"S": "U+F03D4",
		"T": "U+F0240",
		"U": "U+F0156",
		"V": "U+F0540",
		"W": "U+F02D7",
		"X": "U+F0218",
		"Y": "U+F01A2",
		"Z": "U+F0551",
		"a": "U+F0352",
		"b": "U+F04B5",
		"c": "U+F016D",
		"d": "U+F0537",
		"e": "U+F0465",
		"f": "U+F04A5",
		"g": "U+F0229",
		"h": "U+F027C",
		"i": "U+F055A",
		"j": "U+F01BA",
		"k": "U+F038C",
		"l": "U+F0106",
		"m": "U+F0476",
		"n": "U+F018C",
		"o": "U+F04EF",
		"p": "U+F01C6",
		"q": "U+F0500",
		"r": "U+F011E",
		"s": "U+F03FB",
		"t": "U+F01EA",
		"u": "U+F00BD",
		"v": "U+F0311",
		"w": "U+F0586",
		"x": "U+F0239",
		"y": "U+F0401",
		"z": "U+F04DA",
		"0": "U+F034C",
		"1": "U+F00EA",
		"2": "U+F0319",
		"3": "U+F03E6",
		"4": "U+F0549",
		"5": "U+F0122",
		"6": "U+F0452",
		"7": "U+F0293",
		"8": "U+F056E",
		"9": "U+F0267",
		"+": "U+F039A",
		"/": "U+F0596",
		"=": "U+F0416",
	}

	for letter, puaStr := range puaMap {
		pua, _ := strconv.ParseInt(strings.TrimPrefix(puaStr, "U+"), 16, 32)
		letterToPUA[rune(letter[0])] = rune(pua)
		puaToLetter[rune(pua)] = rune(letter[0])
	}
}

func encode(input string) string {
	var result []rune
	for _, letter := range input {
		if pua, ok := letterToPUA[letter]; ok {
			result = append(result, pua)
		} else {
			result = append(result, letter)
		}
	}
	return string(result)
}

func decode(input string) string {
	var result []rune
	for _, pua := range input {
		if letter, ok := puaToLetter[pua]; ok {
			result = append(result, letter)
		} else {
			result = append(result, pua)
		}
	}
	return string(result)
}

func main() {
	decodePtr := flag.Bool("d", false, "Set this flag to decode")
	fontPath := flag.String("f", "Path/to/font.ttf", "Path to the .ttf font file")
	flag.Parse()

	// Read the .ttf file
	fontBytes, err := ioutil.ReadFile(*fontPath)
	if err != nil {
		fmt.Println("Error reading the font file:", err)
		return
	}

	// Load the font
	_, err = freetype.ParseFont(fontBytes)
	if err != nil {
		fmt.Println("Error parsing the font:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if *decodePtr {
			fmt.Println(decode(input))
		} else {
			fmt.Println(encode(input))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
}
