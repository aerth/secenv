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

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-d" {
			if !seconf.Detect("config.enc") {
				fmt.Println("No config file to delete.")
				os.Exit(1)
			}
			fmt.Println("Removing Environment")
			seconf.Destroy("config.enc")
			os.Exit(1)
		}
	}
	if !seconf.Detect("config.enc") && askForConfirmation("No config file found. Would you like to create one?") {
		err := seconf.Lock("config.enc", "Secured Environment", "Enter the first environmental NAME", "Enter the first environmental VALUE", "Enter the second environmental NAME", "Enter the second environmental VALUE")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if len(os.Args) < 2 {
			os.Exit(0)
		}
	}
	if !seconf.Detect("config.enc") {
		fmt.Println("No config.")
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		fmt.Println("User error: No command supplied.")
		fmt.Println("Usage: " + os.Args[0] + " ./yourprogram or " + os.Args[0] + " yourprogram ")
		fmt.Println("Destroy: " + os.Args[0] + " -d")
		os.Exit(1)
	}

	configdecoded, err := seconf.Read("config.enc")
	if err != nil {
		fmt.Println("error:")
		fmt.Println(err)
		os.Exit(1)
	}
	configarray := strings.Split(configdecoded, "::::")
	if len(configarray) < 4 {
		fmt.Println("Bad config magic.")
		os.Exit(1)
	}
	os.Setenv(configarray[0], configarray[1])
	os.Setenv(configarray[2], configarray[3])

	cmd := strings.Join(os.Args[1:], " ")
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf(string(out))
}

// constainsString returns true if a slice contains a string.
func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

// askForConfirmation returns true if the user types one of the "okayResponses"
// https://gist.github.com/albrow/5882501
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
