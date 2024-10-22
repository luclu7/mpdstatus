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
	"password": "passwd"
}
```

## Daemon
There's a daemonisable version available, in the `mpdstatus-daemon` folder. The configuration file is the same, but it's stored at `~/.config/mpdstatus-service.json`.

You can use this systemd Unit:
```
[Unit]
Description=MPD status daemon

[Service]
Type=simple

ExecStart=$HOME/bins/mpdstatus-daemon

[Install]
WantedBy=default.target
```
Dont forget to edit the binary's path.

## Why
I had my [bash script](https://github.com/Luclu7/dotfiles/blob/master/i3/.config/i3/nowplaying.sh) which did the exact same thing, but eh, I was bored so I just redid it in Go (ok, copy pasted 80%).
