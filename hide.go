package hide

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"unicode/utf8"
)

var base64ToTag = map[rune]rune{
	// Uppercase A–Z
	'A': 0xE0041, 'B': 0xE0042, 'C': 0xE0043, 'D': 0xE0044, 'E': 0xE0045, 'F': 0xE0046,
	'G': 0xE0047, 'H': 0xE0048, 'I': 0xE0049, 'J': 0xE004A, 'K': 0xE004B, 'L': 0xE004C,
	'M': 0xE004D, 'N': 0xE004E, 'O': 0xE004F, 'P': 0xE0050, 'Q': 0xE0051, 'R': 0xE0052,
	'S': 0xE0053, 'T': 0xE0054, 'U': 0xE0055, 'V': 0xE0056, 'W': 0xE0057, 'X': 0xE0058,
	'Y': 0xE0059, 'Z': 0xE005A,

	// Lowercase a–z
	'a': 0xE0061, 'b': 0xE0062, 'c': 0xE0063, 'd': 0xE0064, 'e': 0xE0065, 'f': 0xE0066,
	'g': 0xE0067, 'h': 0xE0068, 'i': 0xE0069, 'j': 0xE006A, 'k': 0xE006B, 'l': 0xE006C,
	'm': 0xE006D, 'n': 0xE006E, 'o': 0xE006F, 'p': 0xE0070, 'q': 0xE0071, 'r': 0xE0072,
	's': 0xE0073, 't': 0xE0074, 'u': 0xE0075, 'v': 0xE0076, 'w': 0xE0077, 'x': 0xE0078,
	'y': 0xE0079, 'z': 0xE007A,

	// Digits
	'0': 0xE0030, '1': 0xE0031, '2': 0xE0032, '3': 0xE0033, '4': 0xE0034, '5': 0xE0035,
	'6': 0xE0036, '7': 0xE0037, '8': 0xE0038, '9': 0xE0039,

	// Symbols
	'+': 0xE002B, '/': 0xE002F, '=': 0xE003D,
}

var tagToBase64 = make(map[rune]rune)

func init() {
	for k, v := range base64ToTag {
		tagToBase64[v] = k
	}
}

type concealer struct{}

func New() concealer {
	return concealer{}
}

// Hide encodes the input string to a tag-encoded string.
func (h concealer) Hide(input string) (string, error) {
	base64Encoded := base64.StdEncoding.EncodeToString([]byte(input))
	var buf bytes.Buffer

	for _, c := range base64Encoded {
		tagCode, ok := base64ToTag[c]
		if !ok {
			return "", fmt.Errorf("unsupported Base64 char: '%c'", c)
		}
		buf.WriteRune(tagCode)
	}

	return buf.String(), nil
}

// Unhide decodes the tag-encoded string back to the original string.
func (h concealer) Unhide(tagEncoded string) (string, error) {
	var base64Builder bytes.Buffer

	for i := 0; i < len(tagEncoded); {
		r, size := utf8.DecodeRuneInString(tagEncoded[i:])
		if r == utf8.RuneError {
			return "", fmt.Errorf("invalid UTF-8 at position %d", i)
		}
		base64Char, ok := tagToBase64[r]
		if !ok {
			return "", fmt.Errorf("unsupported tag character: U+%04X", r)
		}
		base64Builder.WriteRune(base64Char)
		i += size
	}

	decoded, err := base64.StdEncoding.DecodeString(base64Builder.String())
	if err != nil {
		return "", fmt.Errorf("base64 decode failed: %w", err)
	}

	return string(decoded), nil
}
