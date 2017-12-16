package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
)

func main() {
	body, _ := ioutil.ReadAll(os.Stdin)

	pattern := regexp.MustCompile(`[A-Z]`)
	m := pattern.FindAll(body, -1)

	ans := 1
	if m != nil {
		ans += len(m)
	}

	fmt.Print(ans)
}
