/*
 * Copyright (c) 2022 NetLOX Inc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at:
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"os"
	"os/exec"

	"github.com/jessevdk/go-flags"
	ln "github.com/loxilb-io/loxilb/loxinet"
	opts "github.com/loxilb-io/loxilb/options"
)

// utility variables
const (
	MkfsScript     = "/usr/local/sbin/mkllb_bpffs"
	BpfFsCheckFile = "/opt/loxilb/dp/bpf/intf_map"
)

func fileExists(fname string) bool {
	info, err := os.Stat(fname)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func fileCreate(fname string) int {
	file, e := os.Create(fname)
	if e != nil {
		return -1
	}
	file.Close()
	return 0
}

// loxiSyncMain - NSync subsystem init
func loxiSyncMain() {
	rpcObj := new(ln.NSync)
	rpc.Register(rpcObj)
	rpc.HandleHTTP()

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		io.WriteString(res, "loxilb-nsync\n")
	})

	listener := fmt.Sprintf(":%d", ln.NSyncPort)
	http.ListenAndServe(listener, nil)
}

var version string = "0.7.9"
var buildInfo string = ""

func main() {
	fmt.Printf("loxilb start\n")

	// Parse command-line arguments
	_, err := flags.Parse(&opts.Opts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if opts.Opts.Version {
		fmt.Printf("loxilb version: %s %s\n", version, buildInfo)
		os.Exit(0)
	}

	// It is important to make sure loxilb's eBPF filesystem
	// is in place and mounted to make sure maps are pinned properly
	if fileExists(BpfFsCheckFile) == false {
		if fileExists(MkfsScript) {
			_, err := exec.Command("/bin/bash", MkfsScript).Output()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	go loxiSyncMain()
	ln.Main()
}
