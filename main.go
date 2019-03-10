package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Printf(`
Changes ports of IPFS configurations so they don't overlap. Prints out a list of Peer IDs.

Usage:
	%s {number of sybil nodes}
`, os.Args[0])
	}

	n, err := strconv.Atoi(os.Args[1])
	check(err)

	for i := 1; i <= n; i++ {
		file, err := ioutil.ReadFile(".ipfs" + strconv.Itoa(i) + "/config")
		check(err)

		lines := strings.Split(string(file), "\n")

		for j, line := range lines {
			number1 := strconv.Itoa(4001 + i)
			number2 := strconv.Itoa(5001 + i)
			number3 := strconv.Itoa(8080 + i)

			lines[j] = strings.Replace(line, "/ip4/0.0.0.0/tcp/4001", "/ip4/0.0.0.0/tcp/"+number1, -1)
			lines[j] = strings.Replace(lines[j], "/ip6/::/tcp/4001", "/ip6/::/tcp/"+number1, -1)
			lines[j] = strings.Replace(lines[j], "/ip4/127.0.0.1/tcp/5001", "/ip4/127.0.0.1/tcp/"+number2, -1)
			lines[j] = strings.Replace(lines[j], "/ip4/127.0.0.1/tcp/8080", "/ip4/127.0.0.1/tcp/"+number3, -1)
		}

		output := strings.Join(lines, "\n")

		err = ioutil.WriteFile(".ipfs"+strconv.Itoa(i)+"/config", []byte(output), 0644)
		check(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
