package command

import (
	"os/exec"
	youtube "ytdl/app/youtube-dl"
	"ytdl/utils"
)

func DownloadAudio(url string) {
	cmd := exec.Command(youtube.YOUTUBE, "-x", "--audio-format", "mp3", url)
	defer utils.UrlError()

	runCommand(cmd)
}
