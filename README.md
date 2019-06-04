# MPDStatus

Just a small notification to inform you what is currently playing.

![Screenshot](https://raw.githubusercontent.com/Luclu7/mpdstatus/master/screenshot.png)

## Install
It's just a normal Go program:
```
go get -u github.com/luclu7/mpdstatus
```

## Configuration
The configuration file is stored at `~/.config/mpdstatus.json`. It is automatically generated at the initial launch of MPDStatus. It's as simple that:
```json
{
	"address": "localhost",
	"port": 6600,
	"auth": false,
	"password:" "passwd"
}
```


## Why
I had my [bash script](https://github.com/Luclu7/dotfiles/blob/master/i3/.config/i3/nowplaying.sh) which did the exact same thing, but eh, I was bored so I just redid it in Go (ok, copy pasted 80%).
