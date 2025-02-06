package flagsparse

import (
	"flag"
	"fmt"
	"reflect"
	"strings"
)

// LoadFlags takes a struct pointer and binds its fields to command-line flags.
func LoadFlags(config any) {
	val := reflect.ValueOf(config).Elem() // get struct value
	typ := val.Type()                     // this is the type

	// itarate in the structure
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		//if there is no flag return empty but this is to get the flag value
		tag := field.Tag.Get("flag")
		if tag == "" {
			continue // this is to skip fields without flag tag
		}

		name := strings.Split(tag, ",")[0] // extract flag name
		usage := field.Tag.Get("usage")    //get usage tag value

		switch field.Type.Kind() {
		case reflect.String:
			value := val.Field(i).String()
			flag.StringVar(val.Field(i).Addr().Interface().(*string), name, value, usage)
		case reflect.Int:
			value := int(val.Field(i).Int())
			flag.IntVar(val.Field(i).Addr().Interface().(*int), name, value, usage)
		case reflect.Bool:
			value := val.Field(i).Bool()
			flag.BoolVar(val.Field(i).Addr().Interface().(*bool), name, value, usage)
		default:
			fmt.Printf("Unsupported field type: %s\n", field.Type.Kind())
		}
	}
	flag.Parse()
}

// In this example, the Config struct has three fields tagged with "flag" and "usage" values.
// When running the program with command-line arguments like --name=John --age=30 --verbose,
// the LoadFlags function will automatically bind these values to the corresponding struct fields.
// This is a real-world example of how reflection can be used to build a generic command-line argument parser that can
// work with arbitrary struct types, without needing to write boilerplate code for each struct.
/*
type Config struct {
	Name    string `flag:"name" usage:"Your name"`
	Age     int    `flag:"age" usage:"Your age"`
	Verbose bool   `flag:"verbose" usage:"Enable verbose logging"`
	Ignored string // This field will be ignored
}

func main() {
	config := &Config{}
	LoadFlags(config)

	fmt.Printf("Name: %s\n", config.Name)
	fmt.Printf("Age: %d\n", config.Age)
	fmt.Printf("Verbose: %t\n", config.Verbose)
}
*/
