package main

import (
	wh "palhook/internal/webhook"
)

func main() {
	if false {
		client := wh.NewClient()

		client.OverrideIdentity(wh.GetIdentityByPreset(wh.SneakyDetector))
	}
}
