package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"os"

	"github.com/cyrildever/feistel"
)

func encryptFromStdin() {
	input, _ := ioutil.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}
	secretkey, err := ioutil.ReadFile(flag.Arg(0))
	rounds := flag.Arg(1)
	i, err := strconv.Atoi(rounds)
	if err != nil {
		panic(err)
	}
	cipher := feistel.NewCipher(string(secretkey), i)
        obfuscated, err := cipher.Encrypt(string(input))
	fmt.Print(string(obfuscated))
}

func decryptFromStdin() {
	input, _ := ioutil.ReadAll(os.Stdin)
	if len(input) == 0 {
		return
	}
	secretkey, err := ioutil.ReadFile(flag.Arg(0))
	rounds := flag.Arg(1)
	i, err := strconv.Atoi(rounds)
	if err != nil {
		panic(err)
	}
	cipher := feistel.NewCipher(string(secretkey), i)
        deciphered, err := cipher.Decrypt(input)
	fmt.Print(string(deciphered))
}

func main() {

	decryptFlag := flag.Bool("d", false, "Decrypt: feistel -d keyfile 10 < infile > outfile")

	flag.Parse()
	
	if decryptFlag != nil && *decryptFlag {
		decryptFromStdin()
	} else {
		encryptFromStdin()
	}
}
