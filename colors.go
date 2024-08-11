package main

import "fmt"

func Red(text string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", text)
}

func Green(text string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", text)
}

func Yellow(text string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", text)
}

func Blue(text string) string {
	return fmt.Sprintf("\033[34m%s\033[0m", text)
}

func PrintMessage(message string, color func(string) string) {
	fmt.Printf("%s\r\n", color(message))
}
