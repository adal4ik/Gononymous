package utils

import "fmt"

// Help Flag
func HelpFlag() {
	fmt.Print("Simple Storage Service.\n\n**Usage:**\n	triple-s [-port <N>] [-dir <S>]\n	triple-s --help\n\n**Options:**\n- --help     Show this screen.\n- --port N   Port number\n- --dir S    Path to the directory\n")
}
