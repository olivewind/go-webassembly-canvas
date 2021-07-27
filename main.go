package main

import (
	"math/rand"
	"syscall/js"
	"time"
)

const (
	width = 400
	height = 400
)

func getRandomNum() float32 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := float32(rand.Intn(10000))
	return  n / 10000.0
}

func render() {
	var canvas js.Value = js.
		Global().
		Get("document").
		Call("getElementById", "canvas")

	var context js.Value = canvas.Call("getContext", "2d")

	// reset
	canvas.Set("height", height)
	canvas.Set("width", width)
	context.Call("clearRect", 0, 0, width, height)

	for i := 0; i < 50; i ++ {
		context.Call("beginPath")
		context.Call("moveTo", getRandomNum() * width, getRandomNum() * height)
		context.Call("lineTo", getRandomNum() * width, getRandomNum() * height)
		context.Call("stroke")
	}
}

func addEventListener()  {
	// see https://tip.golang.org/pkg/syscall/js/?GOOS=js&GOARCH=wasm#NewCallback
	done := make(chan struct{})
	var cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			println("click")
			render()
		}()
		return nil
	})
	js.
		Global().
		Get("document").
		Call("getElementById", "canvas").
		Call("addEventListener", "click", cb)
	<-done

}


func bootstrapApp() {
	render()
	addEventListener()
}

func main() {
	println("wasm app works")
	// bootstrap app
	bootstrapApp()
}
