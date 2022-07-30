package main

import "C"
import (
	"context"
	"encoding/json"
	"sync"

	"github.com/maiderdiscord/discord"
)

type CheckTokensResult struct {
	Token string `json:"token"`
	Alive bool   `json:"alive"`
}

type CheckTokensInput struct {
	Tokens    []string  `json:"tokens"`
	Proxies   []string  `json:"proxies"`
	ProxyType ProxyType `json:"proxyType"`
}

type CheckTokensOutput struct {
	Success bool                `json:"success"`
	Result  []CheckTokensResult `json:"result"`
}

//export CheckTokens
func CheckTokens(data *C.char) *C.char {
	return C.CString(checkTokens(C.GoString(data)))
}

func checkTokens(data string) string {
	in := new(CheckTokensInput)
	out := new(CheckTokensOutput)

	if err := json.Unmarshal([]byte(data), in); err != nil {
		o, _ := json.Marshal(out)
		return string(o)
	}

	ctx := context.Background()
	wg := new(sync.WaitGroup)
	proxiesIndex := 0
	m := sync.Map{}

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
				m.Store(token, false)
				return
			}

			_, err := client.Me(ctx)
			m.Store(token, err != nil)
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

	m.Range(func(k, v interface{}) bool {
		out.Result = append(out.Result, CheckTokensResult{
			Token: k.(string),
			Alive: v.(bool),
		})

		return true
	})

	out.Success = true
	o, _ := json.Marshal(out)
	return string(o)
}
