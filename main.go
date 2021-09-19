// DiscordLastfmScrobbler project main.go
package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-ini/ini"
	"github.com/shkh/lastfm-go/lastfm"
)

// Print prints tracks
func Print(text string) {
	log.Println(" - ", text)
}

func scrobbler() error {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Println(err)
		return err
	}

	token := cfg.Section("discord").Key("token").String()
	apiKey := cfg.Section("lastfm").Key("api_key").String()
	username := cfg.Section("lastfm").Key("username").String()
	title := cfg.Section("app").Key("title").String()
	endlessMode, err := strconv.ParseBool(cfg.Section("app").Key("endless_mode").String())
	configInterval, err := cfg.Section("lastfm").Key("interval").Int()

	if err != nil {
		log.Println(err)
		return err
	}

	api := lastfm.New(apiKey, "")

	Print("Settings loaded: config.ini")

	if endlessMode {
		Print("Endless mode! Ctrl+C to exit")
	}
	dg, err := discordgo.New(token)
	if err != nil {
		log.Println("Discord error: ", err)
		return err
	}
	Print("Authorized to Discord")
	if err := dg.Open(); err != nil {
		log.Println("Discord error: ", err)
		return err
	}
	Print("Connected to Discord")

	interval := time.Duration(configInterval*1000) * time.Millisecond
	ticker := time.NewTicker(interval)
	var prevTrack = ""

	defer func() {
		defer dg.UpdateStatusComplex(discordgo.UpdateStatusData{
			Game:   nil,
			Status: "offline",
		})
		dg.Close()
	}()

	for {
		select {
		case <-ticker.C:
			result, err := api.User.GetRecentTracks(lastfm.P{"limit": "1", "user": username})
			if err != nil {
				log.Println("LastFM error: ", err)
			} else {
				if len(result.Tracks) > 0 {
					currentTrack := result.Tracks[0]
					isNowPlaying, _ := strconv.ParseBool(currentTrack.NowPlaying)
					trackName := currentTrack.Artist.Name + " - " + currentTrack.Name
					if isNowPlaying && prevTrack != trackName {
						prevTrack = trackName
						statusData := discordgo.UpdateStatusData{
							Game: &discordgo.Game{
								Name:    title,
								Type:    discordgo.GameTypeListening,
								Details: prevTrack,
								State:   "DiscordLastfmScrobbler",
							},
							AFK:    false,
							Status: "online",
						}
						if err := dg.UpdateStatusComplex(statusData); err != nil {
							log.Println("Discord error: ", err)
							if !endlessMode {
								return err
							}
						}
						Print("Now playing: " + trackName)
					} else if !isNowPlaying {
						log.Println("!")
						statusData := discordgo.UpdateStatusData{
							Game:   nil,
							Status: "offline",
						}
						if err := dg.UpdateStatusComplex(statusData); err != nil {
							log.Println("Discord error: ", err)
							if !endlessMode {
								return err
							}
						}
					}
				}
			}
		}
	}
}

func main() {
	_ = scrobbler()
	log.Println("Press the Enter Key to terminate the console screen!")
	_, _ = fmt.Scanln()
}
