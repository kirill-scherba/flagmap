// Copyright 2023 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Flagmap package contains helper function to simplify adding and parsing
// applications flags. Flags definition added to go map, parsed with Parse
// mesod and get result with Get method.
package flagmap

import (
	"flag"
	"fmt"
)

// Application flags map by flag name
type Flags map[string]*FlagData

// Application flags map data structure
type FlagData struct {
	Value interface{}
	Usage string
}

// Parse Application flags
func (f Flags) Parse() Flags {

	for flagName, flg := range f {
		var value interface{}

		switch defaultValue := flg.Value.(type) {
		case int:
			value = flag.Int(flagName, defaultValue, flg.Usage)
		case bool:
			value = flag.Bool(flagName, defaultValue, flg.Usage)
		case string:
			value = flag.String(flagName, defaultValue, flg.Usage)
		case float64:
			value = flag.Float64(flagName, defaultValue, flg.Usage)
		}

		f[flagName].Value = value

	}
	flag.Parse()

	return f
}

// Stringify flags
func (f Flags) String() (str string) {
	var strIf = func(flagName string, value interface{}) (str string) {
		const format = "%s: %v\n"
		var sprintf = func(name string, value interface{}) string {
			return fmt.Sprintf(format, name, value)
		}
		switch v := value.(type) {
		case *int:
			str += sprintf(flagName, *v)
		case *bool:
			str += sprintf(flagName, *v)
		case *string:
			str += sprintf(flagName, *v)
		case *float64:
			str += sprintf(flagName, *v)
		default:
			str += fmt.Sprintf("%s: %v\n", flagName, v)
		}
		return
	}
	// str += fmt.Sprintln("got flags:")
	for flagName, flg := range f {
		str += strIf(flagName, flg.Value)
	}
	return str
}

// Get flag value by name
func (f Flags) Get(name string) interface{} {
	value := f[name].Value
	switch v := value.(type) {
	case *int:
		return *v
	case *bool:
		return *v
	case *string:
		return *v
	case *float64:
		return *v
	}
	return nil
}

// Get flag usage by name
func (f Flags) Usage(name string) string {
	return f[name].Usage
}
