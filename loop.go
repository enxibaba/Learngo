package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	if n <= 0 {
		return fmt.Sprintf("can't convert numbers: %d\n", n)
	}
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(9),
		convertToBin(12312312),
		convertToBin(-4),
	)

	scannerFile()
}

func scannerFile() {
	filename := "abc.txt"

	if contents, err := os.Open(filename); err != nil {
		fmt.Printf("open file error: %g", err)
	} else {
		scanner := bufio.NewScanner(contents)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			fmt.Printf("@ %s @\n", scanner.Text())
		}
	}

}
