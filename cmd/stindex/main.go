// Copyright (C) 2014 The Syncthing Authors.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {
	var mode string
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	flag.StringVar(&mode, "mode", "dump", "Mode of operation: dump, dumpsize")

	flag.Parse()

	path := flag.Arg(0)
	if path == "" {
		path = filepath.Join(defaultConfigDir(), "index-v0.11.0.db")
	}

	fmt.Println("Path:", path)

	ldb, err := leveldb.OpenFile(path, &opt.Options{
		ErrorIfMissing:         true,
		Strict:                 opt.StrictAll,
		OpenFilesCacheCapacity: 100,
	})
	if err != nil {
		log.Fatal(err)
	}

	if mode == "dump" {
		dump(ldb)
	} else if mode == "dumpsize" {
		dumpsize(ldb)
	} else {
		fmt.Println("Unknown mode")
	}
}
