package builddeploy

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Build takes the name of the repo. It executes the bash script located at the ./scripts folder, that has to be named as the repo with the "sh" extension. Example: for test, test.sh.
func Build(repo string) error {
    script := "./scripts/" + repo + ".sh"
    if _, err := os.Stat(script); os.IsNotExist(err) {
		fmt.Println(err.Error())
        return err
	}
	cmd := exec.Command("bash" + script)
	stdout, err := cmd.Output()
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	fmt.Println(string(stdout))
	return err
}
