package main

import "os/exec"

func main() {

	exec.Command("xdg-open", "https://www.youtube.com/watch?v=dQw4w9WgXcQ").Start()

}
