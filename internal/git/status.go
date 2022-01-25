package git

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

func checkGitStatus(path string) (exitCode int) {

	cmd := exec.Command("git", "status")
	cmd.Dir = path
	err := cmd.Run()
	var (
		ee *exec.ExitError
		pe *os.PathError
	)

	if errors.As(err, &ee) {
		// ran, but non-zero exit code
		exitCode = ee.ExitCode()

	} else if errors.As(err, &pe) {
		// "no such file ...", "permission denied" etc.
		log.Printf("os.PathError: %v", pe)
		exitCode = 1

	} else if err != nil {
		// something really bad happened!
		log.Printf("general error: %v", err)
		exitCode = 1

	} else {
		// ran without error (exit code zero)
		exitCode = 0
	}

	return

}
