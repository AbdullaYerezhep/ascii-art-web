package asciiart

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
)
var hash = map[string]string{
	"standard":"2d177cae56dba2451b87afc493b74e3e",
	"shadow": "c6246987948ea88a66b07939c613f916",
	"thinkertoy": "ed256fdc3317714d3a526aa56e39194d"
}

func AsciiArt(text, theme string) (string, error) {
	x := ("./pkg/ascii-art/" + theme + ".txt")
	content, err := os.ReadFile(x)
	if err != nil {
		return "", err
	}

	if GetMDHash(string(content)) != hash[theme] {
		return "", errors.New("Style file has been changed!")
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

func GetMDHash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}