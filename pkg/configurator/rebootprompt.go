package configurator

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
)

func rebootPrompt(c *ssh.Client) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("configuration complete.  Would you like to reboot your robot? (y/n)")

	for {
		fmt.Print("-> ")
		r, _ := reader.ReadString('\n')
		r = strings.ReplaceAll(r, "\n", "")
		fmt.Println(r)
		if r == "y" {
			if err := reboot(c); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			os.Exit(0)
		} else {
			os.Exit(0)
		}
	}

}
