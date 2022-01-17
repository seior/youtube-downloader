package utils

import "fmt"

func UrlError() {
	err := recover()

	if err != nil {
		fmt.Println("\terror occurred!, please check your connection or make sure url was correct\n")
	}
}
