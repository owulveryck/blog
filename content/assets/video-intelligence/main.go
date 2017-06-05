package main

// go:generate gopherjs build main.go -o app.js -m
// +build ignore

import (
	"github.com/gopherjs/gopherjs/js"
	"log"
	"time"
)

func main() {

	// var tag = document.createElement("script");
	tag := js.Global.Get("document").Call("createElement", "script")
	// tag.src = "//www.youtube.com/iframe_api";
	tag.Set("src", "//www.youtube.com/iframe_api")
	// var firstScriptTag = document.getElementsByTagName("script")[0];
	// firstScriptTag.parentNode.insertBefore(tag, firstScriptTag);
	scriptTags := js.Global.Get("document").Call("getElementsByTagName", "script")
	scriptTags.Index(0).Get("parentNode").Call("insertBefore", tag, scriptTags.Index(0))
	// // This function creates an <iframe> (and YouTube player)
	// // after the API code downloads.
	// var player;
	// window.onYouTubeIframeAPIReady = function() {
	//   player = new YT.Player("player", {
	//       "height": "315",
	//       "width": "560",
	//       "videoId": "A0yQ0dPhkOg",
	//       "events": {
	//       "onReady": onPlayerReady,
	//       "onStateChange": onPlayerStateChange
	//       }
	//       });
	// }
	// Create one configuration object that will be transpiled in json Object
	// and passed to the constructor of the player
	// See https://github.com/gopherjs/gopherjs/wiki/JavaScript-Tips-and-Gotchas
	type ytConfig struct {
		*js.Object        // so far so good
		Height     string `js:"height"`
		Width      string `js:"width"`
		VideoID    string `js:"videoId"`
	}
	// Create the configuration
	config := &ytConfig{Object: js.Global.Get("Object").New()}
	config.Height = "315"
	config.Width = "560"
	config.VideoID = "A0yQ0dPhkOg"
	js.Global.Get("window").Set("onYouTubeIframeAPIReady", func() {
		// Then create a new Player instance called "player", actually creating an iFrame "player" instead of the
		// Div identified by "player"
		var player *ytPlayer
		player = &ytPlayer{*(js.Global.Get("YT").Get("Player").New("player", config)), make(chan string)}
		player.Call("addEventListener", "onReady", player.onPlayerReady)
		player.Call("addEventListener", "onStateChange", player.onPlayerStateChange)
	})
}

type ytPlayer struct {
	js.Object
	state chan string
}

func (yt *ytPlayer) onPlayerReady(event *js.Object) {
	// Trigger the goroutine that will display the current time of the video
	go func() {
		var state string
		for {
			select {
			case state = <-yt.state:
			default:
			}
			if state == "1" {
				log.Println(yt.getCurrentTime())
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()
}
func (yt *ytPlayer) onPlayerStateChange(event *js.Object) {
	go func() {
		yt.state <- event.Get("data").String()
	}()
}

func (yt *ytPlayer) getCurrentTime() (time.Duration, error) {
	return time.ParseDuration(yt.Call("getCurrentTime").String() + "us")
}
