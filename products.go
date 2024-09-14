// Copyright (c) 2020-2024 Andrew Stormont
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"context"
	"net"
	"strings"
	"sync"

	"github.com/grandcat/zeroconf"
)

type ProductDetails struct {
	IPs  []net.IP
	Name string
	Jid  string
}

func discoverProducts(ctx context.Context) (map[string]*ProductDetails, error) {
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	products := make(map[string]*ProductDetails)

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		defer wg.Done()

		for entry := range results {
			m := make(map[string]string)
			for _, t := range entry.Text {
				if s := strings.SplitN(t, "=", 2); len(s) == 2 {
					m[s[0]] = s[1]
				}
			}

			products[m["jid"]] = &ProductDetails{
				IPs:  entry.AddrIPv4,
				Name: m["name"],
				Jid:  m["jid"],
			}
		}
	}(entries)

	err = resolver.Browse(ctx, "_beoremote._tcp", "local", entries)
	if err != nil {
		return nil, err
	}

	wg.Wait()

	return products, nil
}
