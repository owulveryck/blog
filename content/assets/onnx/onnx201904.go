// +build wasm

package main

import (
	"fmt"
	"log"
	"syscall/js"
	"time"

	"github.com/vincent-petithory/dataurl"
)

func main() {
	files := js.Global().Get("document").Call("getElementById", "knowledgeFile").Get("files")
	fmt.Println("file", files)
	fmt.Println("Length", files.Length())
	if files.Length() == 1 {
		fmt.Println("Reading from uploaded file")
		reader := js.Global().Get("FileReader").New()
		reader.Call("readAsDataURL", files.Index(0))
		for reader.Get("readyState").Int() != 2 {
			fmt.Println("Waiting for the file to be uploaded")
			time.Sleep(1 * time.Second)
		}
		content := reader.Get("result").String()
		dataURL, err := dataurl.DecodeString(content)
		if err != nil {
			log.Println(err)
			return
		}
		// modelonnx = dataURL.Data
	}
	// Declare callback
	cb := js.NewEventCallback(js.PreventDefault, func(ev js.Value) {
		// handle event
		js.Global().Get("document").
			Call("getElementById", "guess").
			Set("value", "hello wasm")
	})
	// Hook it up with a DOM event
	js.Global().Get("document").
		Call("getElementById", "btnSubmit").
		Call("addEventListener", "click", cb)
}
