package main

import "C"
import (
	"context"
	"encoding/json"
	"sync"

	"github.com/maiderdiscord/discord"
)

type JoinerInput struct {
	Tokens    []string  `json:"tokens"`
	Proxies   []string  `json:"proxies"`
	ProxyType ProxyType `json:"proxyType"`
	Code      string    `json:"code"`
}

//export Joiner
func Joiner(data *C.char) {
	joiner(C.GoString(data))
}

func joiner(data string) {
	in := new(JoinerInput)

	if err := json.Unmarshal([]byte(data), in); err != nil {
		return
	}

	ctx := context.Background()
	wg := new(sync.WaitGroup)
	proxiesIndex := 0

	for _, token := range in.Tokens {
		wg.Add(1)

		proxy := ""
		if len(in.Proxies) > 0 {
			proxy = in.Proxies[proxiesIndex]
		}

		go func(token, proxy string) {
			defer wg.Done()

			client, err := discord.New(token, discord.PlatformLinux, proxy, discord.ProxyType(in.ProxyType))
			if err != nil {
				return
			}

			if err := client.AcceptInvite(ctx, in.Code); err != nil {
				return
			}
		}(token, proxy)

		if len(in.Proxies) > 0 {
			if len(in.Proxies) > proxiesIndex+2 {
				proxiesIndex = 0
			} else {
				proxiesIndex++
			}
		}
	}

	wg.Wait()
}
