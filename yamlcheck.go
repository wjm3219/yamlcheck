/*
 * COPYRIGHT Terry.Wei 2016, all rights reserved.
 * Author: Terry.Wei
 * Email: wjm.terry@gmail.com
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-yaml/yaml"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: yamlcheck <file>\n")
		flag.PrintDefaults()
	}

	flag.Parse()
}

func main() {

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	bytes, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file %s: %v\n", flag.Arg(0), err)
		os.Exit(2)
	}

	var output interface{}
	err = yaml.Unmarshal(bytes, &output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Yamlcheck err %v \n", err)
		os.Exit(3)
	}

	fmt.Fprintf(os.Stderr, "Yamlcheck success.\n")
}
