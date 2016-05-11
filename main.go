// The MIT License (MIT)
//
// Copyright (c) 2016 aerth
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
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/aerth/seconf"
)

var usageinfo = `
			
[seconf version 0.8]

	usage
		
		Delete Config: secenv -d			
		Create Config: secenv
		Usage: secenv env
			   secenv yourprogram
			   secenv 'yourprogram has -flags /etc'
`
var configarray []string

// Check Arguments
func checkflags() {

	// Since running secenv is config creation, not checking for 'not enough args' until further down.
	if len(os.Args) > 1 {

		// Show version and info
		if os.Args[1] == "-h" || os.Args[1] == "-v" || os.Args[1] == "help" {
			fmt.Println(usageinfo)
			os.Exit(1)
		}

		// Delete ~/.secenv if possible
		if os.Args[1] == "-d" {
			if !seconf.Detect("secenv") {
				fmt.Println("No config file to delete.")
				os.Exit(1)
			}
			fmt.Println("Removing Environment")
			seconf.Destroy("secenv")
			os.Exit(1)
		}
	}
}

// Load configuration
func doconf() []string {
	if !seconf.Detect("secenv") && askForConfirmation("Welcome to secenv. No config file found at ~/.secenv, would you like to create one?") {
		err := seconf.Lock("secenv", // name of config file ( in $HOME )
			"Secured Environment", // Title (for display only)
			"Enter the first environmental NAME",
			"Enter the first environmental VALUE",
			"Enter the second environmental NAME",
			"Enter the second environmental VALUE") // We could keep going but secenv is just for pair of key:value
		if err != nil {
			fmt.Println(err) // Print error from seconf library
			os.Exit(1)
		}
		// Config file exists and is unloaded, but there is nothing to run.
		if len(os.Args) < 2 {
			os.Exit(0)
		}
	}

	if !seconf.Detect("secenv") {
		fmt.Println("No config.")
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		fmt.Println("User error: No command supplied.")
		fmt.Println(usageinfo)
		os.Exit(1)
	}

	configdecoded, err := seconf.Read("secenv")
	if err != nil {
		fmt.Println("error:")
		fmt.Println(err)
		os.Exit(1)
	}

	configarray := strings.Split(configdecoded, "::::")
	return configarray
}

func main() {

	// delete and exit if -d is set
	checkflags()

	// otherwise operate as normal, create or decode config.
	configarray = doconf()

	// corrupt or wrong version config file
	if len(configarray) < 4 {
		fmt.Println("Bad config magic.")
		os.Exit(1)
	}

	// Set first variable
	os.Setenv(configarray[0], configarray[1])
	// Set second variable
	os.Setenv(configarray[2], configarray[3])

	// Run command with the env under bash
	cmd := strings.Join(os.Args[1:], " ")
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err)
	}

	// Print output
	fmt.Printf(string(out))

}

// constainsString returns true if a slice contains a string.
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

// askForConfirmation returns true if the user types one of the "okayResponses"
// adapted from https://gist.github.com/albrow/5882501
func askForConfirmation(p string) bool {
	fmt.Println(p)
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		fmt.Println(err)
		return false
	}
	okayResponses := []string{"y", "Y", "yes", "Yes", "YES"}
	nokayResponses := []string{"n", "N", "no", "No", "NO", "\n"}
	quitResponses := []string{"q", "Q", "exit", "quit"}
	if containsString(okayResponses, response) {
		return true
	} else if containsString(nokayResponses, response) {
		return false
	} else if containsString(quitResponses, response) {
		return false
	} else {
		fmt.Println("\nNot valid answer, try again. [y/n] [yes/no]")
		return askForConfirmation(p)
	}
}

// posString returns the first index of element in slice.
// If slice does not contain element, returns -1.
func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}
