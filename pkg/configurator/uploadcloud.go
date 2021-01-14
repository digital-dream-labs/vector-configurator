package configurator

import (
	"fmt"
	"log"
	"os"
)

const (
	ankiPath   = "/anki/bin"
	rootHome   = "/home/root"
	backupPath = "/home/root/backups"

	cloudBin       = "vic-cloud"
	cloudOwnership = "cloud:anki"
	cloudPerms     = "755"

	gwBin       = "vic-gateway"
	gwOwnership = "net:anki"
	gwPerms     = "755"
)

// UploadCloud copies vic-cloud and vic-gateway to the robot + sets permissions, etc
func UploadCloud(host, key, path string) {

	c, err := getSSHConn(key, host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := runCmds(c, []string{"mount -o remount rw /"}); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// copy vic-cloud
	if err := scpFile(
		c,
		fmt.Sprintf("%s/%s", path, cloudBin),
		fmt.Sprintf("%s/%s", rootHome, cloudBin),
	); err != nil {
		log.Fatal(err)
	}

	// copy vic-gateway
	if err := scpFile(
		c,
		fmt.Sprintf("%s/%s", path, gwBin),
		fmt.Sprintf("%s/%s", rootHome, gwBin),
	); err != nil {
		log.Fatal(err)
	}

	if err := runCmds(
		c,
		[]string{
			fmt.Sprintf("mkdir -p %s", backupPath),
			fmt.Sprintf("cp -pf %s/%s %s", ankiPath, cloudBin, backupPath),
			fmt.Sprintf("cp -pf %s/%s %s", ankiPath, gwBin, backupPath),
			fmt.Sprintf("mv %s/%s %s/%s", rootHome, cloudBin, ankiPath, cloudBin),
			fmt.Sprintf("mv %s/%s %s/%s", rootHome, gwBin, ankiPath, gwBin),
			fmt.Sprintf("chown %s %s/%s", cloudOwnership, ankiPath, cloudBin),
			fmt.Sprintf("chown %s %s/%s", gwOwnership, ankiPath, gwBin),
			fmt.Sprintf("chmod %s %s/%s", cloudPerms, ankiPath, cloudBin),
			fmt.Sprintf("chmod %s %s/%s", gwPerms, ankiPath, gwBin),
		},
	); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rebootPrompt(c)

}
