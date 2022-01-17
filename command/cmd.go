package command

import (
	"fmt"
	"os/exec"
	"ytdl/utils"
)

func runCommand(cmd *exec.Cmd) {
	err := cmd.Start()
	utils.PanicIfErr(err)

	fmt.Println("\n\tStart Downloading...")

	output, _ := cmd.Output()
	fmt.Println(string(output))

	err = cmd.Wait()
	utils.PanicIfErr(err)

	fmt.Println("\tDone...\n")
}
