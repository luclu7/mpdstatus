package main

import (
	"fmt"
	"github.com/fhs/gompd/mpd"

	"github.com/0xAX/notificator"
	"log"
)

var notify *notificator.Notificator

func main() {
	notify = notificator.New(notificator.Options{
		DefaultIcon: "audio-headphones",
		AppName:     "MPD",
	})

	conn, err := mpd.Dial("tcp", "localhost:6600")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	line := ""
	line1 := ""
	// Loop printing the current status of MPD.
	status, err := conn.Status()
	if err != nil {
		log.Fatalln(err)
	}
	song, err := conn.CurrentSong()
	if err != nil {
		log.Fatalln(err)
	}
	if status["state"] == "play" {
		line1 = fmt.Sprintf("Now playing: %s - %s", song["Artist"], song["Title"])
	} else {
		if song["Artist"] == "" {
			line1 = "Now playing: nothing"
		} else {
			line1 = fmt.Sprintf("Paused: %s - %s", song["Artist"], song["Title"])
		}
	}
	if line != line1 {
		line = line1
		fmt.Println(line)
	}

	notify.Push("MPD", line1, "audio-headphones", notificator.UR_NORMAL)
}
