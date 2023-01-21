package asciiart

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func AsciiArt(text, theme string) (string, error) {
	x := ("./pkg/ascii-art/" + theme + ".txt")
	content, err := os.ReadFile(x)
	if err != nil {
		return "", err
	}

	for _, r := range text {
		if !(r >= 32 && r <= 126) {
			return "", errors.New("Wrong characters")
		}
	}

	theme = string(content)
	theme = strings.ReplaceAll(theme, "\r", "")
	lines := strings.Split(theme, "\n")
	text = strings.ReplaceAll(text, `\n`, "\n")
	words := strings.Split(text, "\n")
	result := ""

	for j, word := range words {
		if word != "" {
			for i := 1; i <= 8; i++ {
				for _, r := range word {
					result += lines[int(r-32)*9+i]
				}
				result += "\n"
			}
		} else {
			if j != len(words)-1 {
				result += "\n"
			}
		}
	}

	return result, nil
}

func checkGraphic(text string) error {
	for _, r := range text {
		if !(byte(r) >= 32 && byte(r) <= 126) {
			return fmt.Errorf("Error: Invalid symbol has been inserted!!!")
		}
	}
	return nil
}