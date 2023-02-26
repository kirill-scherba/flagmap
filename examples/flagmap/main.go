// Copyright 2023 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Flags map helper sample Apllication. This application defines flags map,
// parse application flags, print flags and its values
package main

import (
	"fmt"

	"github.com/kirill-scherba/flagmap"
)

const (
	appName    = "Flags map helper sample Apllication"
	appVersion = "0.0.1"
)

// Flags name
const (
	FLAG1 = "flag1"
	FLAG2 = "flag2"
	FLAG3 = "flag3"
	FLAG4 = "flag4"
	FLAG5 = "flag5"
	FLAG6 = "flag6"
)

func main() {
	// Application logo
	fmt.Println(appName + " ver. " + appVersion)

	// Fill flags map and Parse Application flags
	flags := flagmap.Flags{
		FLAG1: {Value: "", Usage: "flag1 string value"},
		FLAG2: {Value: "", Usage: "flag2 string value"},
		FLAG3: {Value: "", Usage: "flag3 string value"},
		FLAG4: {Value: 0, Usage: "flag4 int value"},
		FLAG5: {Value: false, Usage: "flag5 bool value"},
		FLAG6: {Value: 0.0, Usage: "flag6 float value"},
	}.Parse()

	// Strigify and Print Application flags
	fmt.Println("\nApplication flags:")
	fmt.Print(flags)

	// Get flags map Value
	fmt.Printf("\n%s %f\n", FLAG6, flags.Get(FLAG6))

	// Get flags map Usage
	fmt.Printf("%s -> %s\n", FLAG6, flags.Usage(FLAG6))
}
