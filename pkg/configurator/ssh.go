package configurator

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/tmc/scp"
	"golang.org/x/crypto/ssh"
)

// getSSHConn connects to the robot and returns an ssh client
func getSSHConn(key, host string) (*ssh.Client, error) {
	b, err := ioutil.ReadFile(filepath.Clean(key))
	if err != nil {
		return nil, err
	}

	k, err := ssh.ParsePrivateKey(b)
	if err != nil {
		return nil, err
	}

	return ssh.Dial(
		"tcp",
		fmt.Sprintf("%s:22", host),
		&ssh.ClientConfig{
			User: "root",
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(k),
			},
			// nolint: gosec -- I have no idea how this would be handled on windows, so I expect an issue
			// or PR about this at some point -bd
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
	)
}

func runCmds(c *ssh.Client, cmds []string) error {
	for i := range cmds {
		s, err := c.NewSession()
		if err != nil {
			return err
		}
		defer func() {
			_ = s.Close()
		}()

		if err := s.Run(cmds[i]); err != nil {
			return fmt.Errorf(
				"error running command %s: %v",
				cmds[i],
				err,
			)
		}

	}
	return nil
}

func reboot(c *ssh.Client) error {
	s, err := c.NewSession()
	if err != nil {
		return err
	}
	defer func() {
		_ = s.Close()
	}()

	return s.Run("reboot")
}

func scpFile(c *ssh.Client, file, dest string) error {
	s, err := c.NewSession()
	if err != nil {
		return err
	}
	defer func() {
		_ = s.Close()
	}()

	return scp.CopyPath(file, dest, s)
}
