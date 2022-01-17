package main

import (
	"fmt"
	"ytdl/command"
	"ytdl/utils"
	"ytdl/view"
)

func main() {
	// show banner
	view.Banner()
	status := true

	for status {
		option := view.Option()

		switch option {
		case "1":
			url, quality := view.Video()

			if quality == utils.NIL {
				break
			}

			command.DownloadVideo(url, quality)
		case "2":
			url := view.Url()
			command.DownloadAudio(url)
		case "3":
			view.About()
		case "4":
			status = false
		default:
			fmt.Println("\n\tPlease input 1 - 4")
		}
	}

	fmt.Println("\tThanks <3")
}
