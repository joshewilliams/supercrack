package ciphers

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/joshewilliams/supercrack/util"
)

type XorParameters struct {
	Key []byte
}

func (x *XorParameters) String() string {
	return fmt.Sprintf("\nKey:\t%x\n", string(x.Key))
}

func xor(src, key []byte) []byte {
	out := make([]byte, len(src), len(src))
	if len(key) <= 2 {
		for i := range src {
			out[i] = src[i] & key[0]
		}
	} else {
		for i := range src {
			out[i] = src[i] & key[i]
		}
	}
	return out
}

// Xor isn't working correctly for some reason. Isn't returning results correctly.
func Xor(p *util.Parameters) []byte {
	x := XorParameters{}
	XorOptionsMenu(&x)
	return xor(p.Ciphertext, x.Key)
}

func XorOptionsCompleter(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		//Xor Parameters
		{Text: "key=", Description: "Key to XOR against"},
		{Text: "hex", Description: "Key is hex encoded"},
		{Text: "b64", Description: "Key is b64 encoded"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func XorOptionsMenu(x *XorParameters) {
	fmt.Printf("\nXOR Options Menu\n\n")
menuLoop:
	for {
		cmd := prompt.Input("supercrack > ", XorOptionsCompleter, prompt.OptionPrefixTextColor(prompt.Red))
		splitCmd := strings.SplitN(cmd, "=", 2)

		switch splitCmd[0] {
		case "key":
			x.Key = []byte(splitCmd[1])
		case "hex":
			var err error
			len, err := hex.Decode(x.Key, x.Key)
			if err != nil {
				fmt.Println("Input not hex encoded")
				continue menuLoop
			}
			x.Key = x.Key[:len]
		case "b64":
			var err error
			len, err := base64.StdEncoding.Decode(x.Key, x.Key)
			if err != nil {
				fmt.Println("Input not b64 encoded")
				continue menuLoop
			}
			x.Key = x.Key[:len]
		case "run":
			break menuLoop
		case "info":
			fmt.Println(x)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("Bad option selected")
		}
	}
}
