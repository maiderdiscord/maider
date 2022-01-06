package main

import "C"
import (
	"encoding/json"
	"sync"
	"time"

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

//export CheckProxy
func CheckProxy(data *C.char) *C.char {
	in := new(CheckProxyInput)
	out := new(CheckProxyOutput)

	if err := json.Unmarshal([]byte(C.GoString(data)), in); err != nil {
		o, _ := json.Marshal(out)
		return C.CString(string(o))
	}

	proxyType := proxychecker.TypeHTTP

	switch in.Type {
	case ProxyTypeSOCKS5:
		proxyType = proxychecker.TypeSOCKS5
	}

	wg := new(sync.WaitGroup)
	m := sync.Map{}

	for i, proxy := range in.Proxies {
		wg.Add(1)
		go func(proxy string) {
			defer wg.Done()

			alive, _ := proxychecker.Check(proxy, proxyType)
			m.Store(proxy, alive)
		}(proxy)

		if (i+1)%50 == 0 {
			time.Sleep(time.Second * 5)
		}
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
	return C.CString(string(o))
}
