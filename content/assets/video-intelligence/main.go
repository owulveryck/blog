package main

// go:generate gopherjs build main.go -o app.js -m
// +build ignore

import (
	"github.com/gopherjs/gopherjs/js"
	"log"
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
		player = &ytPlayer{*(js.Global.Get("YT").Get("Player").New("player", config))}
		player.Call("addEventListener", "onReady", player.onPlayerReady)
		player.Call("addEventListener", "onStateChange", player.onPlayerStateChange)
	})
}

type ytPlayer struct {
	js.Object
}

func (yt *ytPlayer) onPlayerReady(event *js.Object) {
	log.Println("hello")
}
func (yt *ytPlayer) onPlayerStateChange(event *js.Object) {
	if event.Get("data").String() == "1" {
		time := yt.Call("getCurrentTime").String()
		log.Println(time)
	}
}
