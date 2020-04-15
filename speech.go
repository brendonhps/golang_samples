package main

import "github.com/hegedustibor/htgo-tts"

func main() {
speech := htgotts.Speech{Folder: "audio", Language: "pt"}
speech.Speak("Flying to the moon")

}
	