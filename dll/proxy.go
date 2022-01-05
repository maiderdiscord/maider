package main

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/maiderdiscord/proxychecker"
)

type ProxyType int

const (
	ProxyTypeHTTP ProxyType = iota
	ProxyTypeSOCKS5
)

type CheckProxyResult struct {
	Proxy string `json:"proxy"`
	Alive bool   `json:"alive"`
}

type CheckProxyInput struct {
	Type    ProxyType `json:"type"`
	Proxies []string  `json:"proxies"`
}

type CheckProxyOutput struct {
	Success bool               `json:"success"`
	Result  []CheckProxyResult `json:"result"`
}

func CheckProxy(data string) string {
	in := new(CheckProxyInput)
	out := new(CheckProxyOutput)

	if err := json.Unmarshal([]byte(data), in); err != nil {
		o, _ := json.Marshal(out)
		return string(o)
	}

	proxyType := proxychecker.TypeHTTP

	switch in.Type {
	case ProxyTypeSOCKS5:
		proxyType = proxychecker.TypeSOCKS5
	}

	ctx := context.Background()
	wg := new(sync.WaitGroup)
	m := sync.Map{}

	for _, proxy := range in.Proxies {
		wg.Add(1)
		go func(proxy string) {
			defer wg.Done()

			alive, _ := proxychecker.Check(ctx, proxy, proxyType)
			m.Store(proxy, alive)
		}(proxy)
	}

	wg.Wait()

	m.Range(func(k interface{}, v interface{}) bool {
		result := CheckProxyResult{
			Proxy: k.(string),
			Alive: v.(bool),
		}
		out.Result = append(out.Result, result)
		return true
	})

	out.Success = true
	o, _ := json.Marshal(out)
	return string(o)
}
