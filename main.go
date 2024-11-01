package main

import (
	ascii "Terminal-Art/Ascii"
	ft "Terminal-Art/Format"
	terminal "Terminal-Art/Tml"
	"fmt"
	"os"
	"strings"
)

//declare the common error message for the expected format required as input
func errorText() {
	fmt.Println("Usage: go run . --align=[OPTION] [STRING] [BANNER]")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 4 {
		errorText()
	}
	var (
		alignFormat string = os.Args[1]
		input       string = os.Args[2]
		banner      string = os.Args[3]
	)

	if strings.HasPrefix(alignFormat, "--align=") {
		alignFormat = strings.TrimPrefix(alignFormat, "--align=")
	} else {
		errorText()
	}

	if input == "" {
		errorText()
	}

	// Handle extension errors...
	if strings.Contains(banner, ".") {
		if !strings.HasSuffix(banner, ".txt") {
			fmt.Println("Error: The allowed format is .txt for textual files")
			return
		}
	} else {
		banner = fmt.Sprintf("%s.txt", banner)
	}
	banner = strings.ToLower(banner)
	bannerFile := fmt.Sprintf("Banner_files/%s", banner)
	data, err := os.ReadFile(bannerFile)
	if err != nil {
		fmt.Println("Error reading from file:", err)
		return
	}

	input = strings.ReplaceAll(input, "\\n", "\n")
	input = strings.ReplaceAll(input, "\\t", "    ")
	if strings.Contains(input, "\\b") {
		input = strings.ReplaceAll(input, "\\b", "\b")

		//handle the backspace carefully
		for i := 0; i < len(input); i++ {
			b := strings.Index(input, "\b")
			if b > 0 {
				input = input[:b-1] + input[b+1:]
			}
		}
	}
	words := strings.Split(input, "\n")

	Art := FileStat(words, string(data), bannerFile)

	w, _ := terminal.TerminalDimensions()

	var formatedText string

	switch alignFormat {
	case "center":
		formatedText = ft.FormatCenter(Art, w)
	case "justify":
		formatedText = ft.FormatJustify(input, banner, w)
	case "left":
		formatedText = ft.FormatLeft(Art, w)
	case "right":
		formatedText = ft.FormatRight(Art, w)
	default:
		errorText()
	}

	fileText := DefineBanner(formatedText, banner)
	fmt.Println(fileText)
}

func DefineBanner(text, bannerType string) string {
	var temp string
	switch bannerType {
	case "standard":
		temp = standard(text)
		return temp
	case "shadow":
		temp = shadow(text)
		return temp
	default:
		return text
	}
}

func standard(text string) string {
	return text
}

func shadow(text string) string {
	return text
}

func FileStat(words []string, data, banner string) string {

	fileInfo, err := os.Stat(banner)
	if err != nil {
		fmt.Println("Error reading from the specified file: ", err)
	}

	fileSize := fileInfo.Size()

	var Art string

	if fileSize == 7463 || fileSize == 6623 || fileSize == 4703 {//Windows: || fileSize == 4496  for thinkertoy
		content := strings.Split(string(data), "\n")

		Art = ascii.AsciiArt(words, content)
	} else {
		fmt.Println("Error! ", fileSize)
		os.Exit(1)
	}
	return Art
}
