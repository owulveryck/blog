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
	//
	var player *js.Object
	js.Global.Get("window").Set("onYouTubeIframeAPIReady", func() {
		player = js.Global.Get("YT").Get("Player").New("player")
		player.Call("loadVideoById", "A0yQ0dPhkOg", 5, "large")
		log.Println(player)
	})
}

func onPlayerReady() {
	log.Println("hello")
}
