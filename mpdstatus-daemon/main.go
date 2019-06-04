package main

import (
	"encoding/json"
	"fmt"
	"github.com/0xAX/notificator"
	"github.com/fhs/gompd/mpd"
	homedir "github.com/mitchellh/go-homedir"
	"log"
	"os"
	"strconv"
	"time"
)

type configFile struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Auth     bool   `json:"auth"`
	Password string `json:"password"`
}

func cfe(err error) bool {
	if err != nil {
		log.Panicln(err)
		return false
	}
	return true
}

var notify *notificator.Notificator

func main() {
	notify = notificator.New(notificator.Options{
		DefaultIcon: "audio-headphones",
		AppName:     "MPD",
	})

	homePath, err := homedir.Dir()
	cfe(err)
	configFilepath := homePath + "/.config/mpdstatus-service.json"

	_, err = os.Stat(configFilepath)

	if os.IsNotExist(err) {
		var file, err = os.Create(configFilepath)
		cfe(err)
		// open file using READ & WRITE permission
		file, err = os.OpenFile(configFilepath, os.O_RDWR, 0644)
		cfe(err)
		defer file.Close()

		_, err = file.WriteString("{\n")
		cfe(err)
		_, err = file.WriteString("	\"address\": \"localhost\",\n")
		cfe(err)
		_, err = file.WriteString("	\"port\": 6600,\n")
		cfe(err)
		_, err = file.WriteString("	\"auth\": false,\n")
		cfe(err)
		_, err = file.WriteString("	\"password\": \"passwd\"\n")
		cfe(err)
		_, err = file.WriteString("}\n")
		cfe(err)
		err = file.Sync()
		cfe(err)

		fmt.Println("The config file was created at", configFilepath)

		defer file.Close()
	}

	file, _ := os.Open(configFilepath)
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := configFile{}
	err = decoder.Decode(&config)
	cfe(err)

	var conn *mpd.Client
	port := strconv.Itoa(config.Port)
	addressplusport := config.Address + ":" + port
	if config.Auth == false {
		conn, err = mpd.Dial("tcp", addressplusport)
	} else {
		conn, err = mpd.DialAuthenticated("tcp", addressplusport, config.Password)
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	line := ""
	line1 := ""
	title := "MPD"
	// Loop printing the current status of MPD.
	for {
		status, err := conn.Status()
		if err != nil {
			log.Fatalln(err)
		}
		song, err := conn.CurrentSong()
		if err != nil {
			log.Fatalln(err)
		}
		if status["state"] == "play" {
			title = "Now playing"
			line1 = fmt.Sprintf("%s - %s", song["Artist"], song["Title"])
		} else {
			if song["Artist"] == "" {
				title = "Nothing's playing"
				line1 = ""
			} else {
				title = "Paused"
				line1 = fmt.Sprintf("%s - %s", song["Artist"], song["Title"])
			}
		}
		if line != line1 {
			line = line1
			fmt.Println(line)
			notify.Push("MPD", line1, "audio-headphones", notificator.UR_NORMAL)
		}
		time.Sleep(1e9)
	}
}
