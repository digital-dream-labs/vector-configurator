package configurator

import (
	"fmt"
	"os"
)

const (
	escapepod = `
{
    "jdocs": "escapepod.local:8084",
    "tms": "escapepod.local:8084",
    "chipper": "escapepod.local:8084",
    "check": "escapepod.local/ok",
    "logfiles": "s3://anki-device-logs-prod/victor",
    "appkey": "oDoa0quieSeir6goowai7f"
}`
	prod = `
{
    "jdocs": "jdocs.api.anki.com:443",
    "tms": "token.api.anki.com:443",
    "chipper": "chipper.api.anki.com:443",
    "check": "conncheck.global.anki-services.com/ok",
    "logfiles": "s3://anki-device-logs-prod/victor",
    "appkey": "oDoa0quieSeir6goowai7f"
}`

	cfpath = "/anki/data/assets/cozmo_resources/config/server_config.json"
)

// SetEnvironment updates a vectors environment
func SetEnvironment(host, key, env string) {

	var cf string

	switch env {
	case "escapepod":
		cf = escapepod
	case "production":
		cf = prod
	default:
		fmt.Println(`valid environments are:
		escapepod
		production`)
		os.Exit(1)
	}

	c, err := getSSHConn(key, host)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := runCmds(
		c,
		[]string{
			"mount -o remount rw /",
			fmt.Sprintf("cp %s %s.bak", cfpath, cfpath),
			fmt.Sprintf(`cat > %s << EOF
%s
EOF`,
				cfpath,
				cf,
			),
		},
	); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	rebootPrompt(c)
}
