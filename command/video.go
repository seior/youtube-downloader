package command

import (
	"os/exec"
	youtube "ytdl/app/youtube-dl"
	"ytdl/utils"
)

func DownloadVideo(url string, quality utils.Quality) {
	cmd := exec.Command(youtube.YOUTUBE, "-f", "bestvideo" + "[height="+ string(quality) + "]+bestaudio", url)
	defer utils.UrlError()

	runCommand(cmd)
}
