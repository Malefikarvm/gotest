package main

/**
 * main class project
 */
import (
	"fmt"
	"strconv"
	"bufio"
	"os"
	"goTest/cntrllrs"
	"goTest/cmpnnts"
)

func main() {
	mdls.Generate("root","","local")
	obj := cntrllrs.NewDecToBin()
	fmt.Print("Ingrese un número, lo convertiremos en binario: ")
	var n int
	_, err := fmt.Scanf("%d", &n)
	_ = err
	obj.Dec, obj.Bin = Dectobin(n)
	fmt.Println(obj.Dec, ":", obj.Bin)
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func Dectobin(n int) (int, string) {
	number := n
	var residual int = 0
	binary := ""
	for n != 0 {
		residual = n % 2
		n /= 2
		binary += strconv.Itoa(residual)
	}
	return number, Reverse(binary)
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r) - 1; i < len(r) / 2; i, j = i + 1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
