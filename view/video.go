package view

import (
	"fmt"
	"strconv"
	"ytdl/utils"
)

func Video() (string, utils.Quality) {
	quality := []utils.Quality{utils.SD240, utils.SD360, utils.SD, utils.HD, utils.FHD}

	url := Url()

	for {
		fmt.Println("\tSelect Video Quality :")
		for i, q := range quality {
			fmt.Printf("\t%d. "+string(q)+"\n", i+1)
		}

		fmt.Printf("\t%d. Cancel\n", len(quality)+1)
		fmt.Println("\t*if video quality not available, will automatically select the quality below it")

		option := GetOption(1, uint(len(quality)+1))

		result, err := strconv.Atoi(option)
		if err != nil {
			fmt.Println("\n\tPlease Input 1 - " + strconv.Itoa(len(quality)+1) + "\n")
		} else {
			if result == 6 {
				return url, utils.NIL
			}

			return url, quality[result - 1]
		}
	}
}
