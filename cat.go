package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n bool

func init() {
	flag.Bool("n", false, "行数を出力")
}

func main() {
	for i := 0; i < len(os.Args)-1; i++ {
		sf, err := os.Open(os.Args[i+1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "ファイル読み込みに失敗しました：", err)
		}
		defer sf.Close()

		line := 0
		scanner := bufio.NewScanner(sf)
		for scanner.Scan() {
			fmt.Println(n)
			// if n == true {
			// ここもっとスマートに書けそう
			fmt.Print(line)
			fmt.Print(":")
			// }
			fmt.Println(scanner.Text())
			line++
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "読み込みに失敗しました：", err)
		}
	}

	//ファイルパス系は後でやる
	// path := filepath.Join("dir", "cat.go")
	// fmt.Println(filepath.Ext(path))
	// fmt.Println(filepath.Base(path))
	// fmt.Println(filepath.Dir(path))

}
