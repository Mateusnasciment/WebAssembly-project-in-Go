package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")

	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", "Hello, WebAssembly!")
	document.Get("body").Call("appendChild", p)
	p.Set("className", "block")

	styles := document.Call("createElement", "style")
	styles.Set("innerHTML", `
		.block {
			background-color: #f00;
			color: #fff;
			padding: 1em;
		}
	`)
	document.Get("head").Call("appendChild", styles)
	document.Get("body").Call("appendChild", p)
}
