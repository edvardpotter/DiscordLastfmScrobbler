// DiscordLastfmScrobbler project main.go
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-ini/ini"
	"github.com/shkh/lastfm-go/lastfm"
)

func Print(text string) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(currentTime, " - ", text)
}

func scrobbler() error {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println(err)
		return err
	}

	token := cfg.Section("discord").Key("token").String()
	apiKey := cfg.Section("lastfm").Key("api_key").String()
	username := cfg.Section("lastfm").Key("username").String()
	configInterval, err := cfg.Section("lastfm").Key("interval").Int()

	if err != nil {
		fmt.Println(err)
		return err
	}

	api := lastfm.New(apiKey, "")

	Print("Settings loaded: config.ini")

	dg, err := discordgo.New(token)
	if err != nil {
		fmt.Println("Discord error: ", err)
		return err
	}
	Print("Authorized to Discord")
	if err := dg.Open(); err != nil {
		fmt.Println("Discord error: ", err)
		return err
	}
	Print("Connected to Discord")

	interval := time.Duration(configInterval*1000) * time.Millisecond
	ticker := time.NewTicker(interval)
	var prevTrack = ""

	for {
		select {
		case <-ticker.C:
			result, err := api.User.GetRecentTracks(lastfm.P{"limit": "1", "user": username})
			if err != nil {
				fmt.Println("LastFM error: ", err)
			} else {
				if len(result.Tracks) > 0 {
					currentTrack := result.Tracks[0]
					isNowPlaying, _ := strconv.ParseBool(currentTrack.NowPlaying)
					trackName := currentTrack.Artist.Name + " - " + currentTrack.Name
					if isNowPlaying && prevTrack != trackName {
						prevTrack = trackName
						statusData := discordgo.UpdateStatusData{
							Game: &discordgo.Game{
								Name:    prevTrack,
								Type:    discordgo.GameTypeListening,
								Details: "LAST.FM",
								State:   "DiscordLastfmScrobbler",
							},
							AFK:    false,
							Status: "online",
						}
						if err := dg.UpdateStatusComplex(statusData); err != nil {
							fmt.Println("Discord error: ", err)
							return err
						}
						Print("Now playing: " + trackName)
					}
				}
			}
		}
	}
}

func main() {
	_ = scrobbler()
	fmt.Println("Press the Enter Key to terminate the console screen!")
	_, _ = fmt.Scanln()
}
