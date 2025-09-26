package main

import (
	webview "github.com/webview/webview_go"
)

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("Basic Example")
	w.SetSize(480, 320, webview.HintNone)
	w.SetUserAgent("Dalvik/2.1.0 (Linux; U; Android 6.0.1; Lenovo YT3-850M Build/MMB29M)")
	w.Navigate("https://www.whatismybrowser.com/detect/what-is-my-user-agent/")
	w.Run()
}
