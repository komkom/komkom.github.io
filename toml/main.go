package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"syscall/js"

	"github.com/komkom/toml"
)

var Document = js.Global().Get("document")

func main() {

	fmt.Printf("____test\n")
	//println(`test`)

	js.Global().Set(`format`, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		format()
		return nil
	}))

	js.Global().Set(`clearScreen`, js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		clear()
		return nil
	}))

	<-make(chan struct{})
}

func format() {
	fmt.Printf("____format3\n")

	screen := Document.Call("getElementById", `json`)
	screen.Set(`innerHTML`, ``)

	errMsg := Document.Call("getElementById", `errormsg`)
	errMsg.Set(`innerHTML`, ``)
	errMsg.Set(`hidden`, true)

	tomlScreen := Document.Call("getElementById", `toml`)
	val := tomlScreen.Get(`value`).String()
	reader := bytes.NewBufferString(val)

	tomlReader := toml.New(reader)

	data, err := io.ReadAll(tomlReader)
	if err != nil {
		errMsg.Set(`hidden`, false)
		errMsg.Set(`innerHTML`, fmt.Errorf("toml-error: %w", err).Error())
		return
	}

	pjson, err := prettyJSON(data)
	if err != nil {
		errMsg.Set(`hidden`, false)
		errMsg.Set(`innerHTML`, fmt.Errorf("json-error: %w", err).Error())
		return
	}

	screen.Set(`innerHTML`, pjson)
}

func clear() {
	fmt.Printf("____clear4\n")

	screen := Document.Call("getElementById", `json`)
	screen.Set(`innerHTML`, ``)

	errMsg := Document.Call("getElementById", `errormsg`)
	errMsg.Set(`innerHTML`, ``)
	errMsg.Set(`hidden`, true)
}

func prettyJSON(jsn []byte) (string, error) {

	var pretty bytes.Buffer
	err := json.Indent(&pretty, jsn, "", "&nbsp;&nbsp;&nbsp;")
	if err != nil {
		return ``, err
	}

	return string(pretty.Bytes()), nil
}
