package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		t := strings.HasPrefix(arg, "http://")
		x := ""
		if !t {
			x = "http://" + arg
		} else {
			x = arg
		}
		// 13:19 ex 1.8
		resp, err := http.Get(x)
		if err != nil {
			fmt.Fprintf(os.Stderr, " fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := io.ReadAll(resp.Body)
		// resp.Body.Close()
		// if err != nil {
		// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", arg, err)
		// 	os.Exit(1)
		// }
		// fmt.Printf("%s", b)
		c, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", arg, err)
			os.Exit(1)
		}
		fmt.Println("status", c)
		//24:29 ex 1.7
		fmt.Println(resp.Status)
		// ex.1.9

	}
}
