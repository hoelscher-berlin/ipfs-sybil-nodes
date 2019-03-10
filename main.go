package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	if (len(os.Args) < 1) || (os.Args[1] != "init") || (os.Args[1] != "start") {
		fmt.Printf(`
Changes ports of IPFS configurations so they don't overlap. Prints out a list of Peer IDs.
Available commands:

	init - initializes {number of sybil nodes} sybil nodes
	start - starts {number of sybil nodes} daemons

Usage:
	%s {cmd} {number of sybil nodes}
`, os.Args[0])
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	check(err)

	for i := 1; i <= n; i++ {
		os.Setenv("IPFS_PATH", "~/.ipfsSybil"+strconv.Itoa(i))

		out, _ := exec.Command("ipfs", "init").Output()

		fmt.Printf("Initialising node %d...\n", i)

		fmt.Printf("%s", out)

		home, err := dirWindows()

		if err != nil {
			fmt.Printf("Can't find home directory of current user: %s\n", err)
		}

		configPath := home + "/.ipfsSybil" + strconv.Itoa(i) + "/config"

		file, err := ioutil.ReadFile(configPath)
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

		err = ioutil.WriteFile(configPath, []byte(output), 0644)
		check(err)

		fmt.Printf("Successfully initialised node %d and changed config.\n", i)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dirWindows() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}
