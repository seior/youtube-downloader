package view

import (
	"fmt"
	"ytdl/utils"
)

func Banner() {
	logo := `
 __     __   ____    _    _   _______   _    _   ____    ______     _____    _      
 \ \   / /  / __ \  | |  | | |__   __| | |  | | |  _ \  |  ____|   |  __ \  | |     
  \ \_/ /  | |  | | | |  | |    | |    | |  | | | |_) | | |__      | |  | | | |     
   \   /   | |  | | | |  | |    | |    | |  | | |  _ <  |  __|     | |  | | | |     
    | |    | |__| | | |__| |    | |    | |__| | | |_) | | |____    | |__| | | |____ 
    |_|     \____/   \____/     |_|     \____/  |____/  |______|   |_____/  |______|
`

	fmt.Println(logo)
	fmt.Println("\tcreated by : seior")
}

func Option() string {
	fmt.Print("\n\toption : \n" +
		"\t1. video downloader\n" +
		"\t2. audio downloader\n" +
		"\t3. about\n" +
		"\t4. exit\n",
	)

	option := GetOption(1, 4)

	fmt.Print("\n")

	return option
}

func Url() string {
	fmt.Print("\tURL : ")
	utils.Input.Scan()

	url := utils.Input.Text()

	fmt.Print("\n")

	return url
}

func About() {
	fmt.Println("\tAbout :\n" +
		"\tyang buat ganteng, terima kasih\n")
}

func GetOption(start uint, end uint) string {
	fmt.Printf("\n\tOption [%d - %d] : ", start, end)

	utils.Input.Scan()

	result := utils.Input.Text()

	return result
}
