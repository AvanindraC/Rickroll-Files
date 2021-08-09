package main

import "os/exec"

func main() {

	exec.Command("rundll32", "url.dll,FileProtocolHandler", "https://www.youtube.com/watch?v=dQw4w9WgXcQ").Start()

}
