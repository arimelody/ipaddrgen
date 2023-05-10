package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/TwiN/go-color"
)

// === internal vars (no need to edit these!) ===

var running bool
var i_last_method string
var i_methods []string

//go:embed methods.txt
var i_methods_file embed.FS

// === timings ===

// minimum wait for an update (ms).
var min_wait int64 = 0

// maximum wait for an update (ms).
var max_wait int64 = 50

// === chances ===
// (defaults, can be configured via console arguments!)

// chance (0.0 - 1.0) for the operation to complete each update.
var complete_rate = 0.1

// chance (0.0 - 1.0) for the random IP to be IPv6.
var ipv6_chance = 0.05

// chance (0.0 - 1.0) for the random IP to have a port.
var port_chance = 0.05

func main() {

	if argIndex("h") >= 0 || argIndex("help") >= 0 {
		showHelp()
		return
	}

	parseArgs()

	// read the methods file!
	i_methods = readMethods()
	if i_methods == nil {
		return
	}

	rand.Seed(time.Now().UnixNano())

	var waitdelta = max_wait - min_wait

	running = true
	for running {

		// === create a new IP to do the funny to ===

		// what are we doing (method)
		var method = randomString(i_methods, i_last_method)
		i_last_method = method

		// where are we doing it (address)
		var address string
		if method != "opening all ports on" && method != "closing all ports on" {
			address = newAddress(ipv6_chance, port_chance)
		} else {
			address = newAddress(ipv6_chance, 0.0) // address with no ports
		}

		// === print the damn thing! ===

		// > [doing something to] 127.0.0.1...
		if method != " " {
			fmt.Print(method, " ")
		}
		fmt.Print(address)

		// and now, the slow, painful wait begins...
		// if complete check fails, and we have over 3 wait periods...
		// print periods at random intervals, wasting the user's time
		var progress = 0
		for rand.Float64() > complete_rate || progress < 3 {
			// > ..........
			fmt.Print(".")
			progress++

			// time (ms) after which the progress bar will update
			var waitUntil = time.Now().UnixMilli() +
				int64(min_wait) +
				rand.Int63n(waitdelta)

			// wait until flag time reached
			for time.Now().UnixMilli() < waitUntil {
			}
		}

		// [ √ ]
		fmt.Print("[ " + color.Ize(color.Green, "√") + " ]\n")

	}
}

func showHelp() {
	// content, err := i_help_file.ReadFile("help.txt")

	// if err != nil {
	// 	fmt.Println("help machine broke")
	// }

	// fmt.Print("\n" + string(content) + "\n\n")
	var args = map[string]string{
		"-min=0":   "the minimum time (ms) spent between progress dots.",
		"-max=50":  "the maximum time (ms) spent between progress dots.",
		"-rate=10": "the likelihood (%) the \"operation\" will actually \"complete\".",
		"-ipv6=5":  "the chance (%) of a generated IP address being IPv6.",
		"-port=5":  "the chance (%) of a generated IP address having a port.",
	}

	fmt.Println()
	fmt.Println("\x1b[1m### IPADDRGEN ###\x1b[0m")
	fmt.Println("a silly gag tool made by mellodoot!")
	fmt.Println()
	fmt.Println("https://github.com/mellodoot/ipaddrgen")
	fmt.Println()
	fmt.Println("all command-line arguments are optional, and purely intended to customise the experience per user :)")
	fmt.Println()
	fmt.Println("\x1b[1mUsage:\x1b[0m")
	fmt.Println("\t./ipaddrgen -min=1000 -max=4000 -rate=15")
	fmt.Println()
	fmt.Println("\x1b[1mArguments:\x1b[0m")
	for argument, description := range args {
		fmt.Println("\t" + argument)
		fmt.Println("\t\t" + description)
	}
	fmt.Println()
}

// parseArgs parses arguments.
//
// not really gonna document this one, it just saves space
// moving this code to its own function.
func parseArgs() {
	var arg_min_wait int64 = intFromArg("min") // -min=0
	if arg_min_wait >= 0 {
		min_wait = arg_min_wait
	}

	var arg_max_wait int64 = intFromArg("max") // -max=50
	if arg_max_wait >= 0 {
		max_wait = arg_max_wait
	}

	var arg_rate int64 = intFromArg("rate") // -rate=10
	if arg_rate >= 0 {
		complete_rate = float64(arg_rate) / 100
	}

	var arg_ipv6_rate int64 = intFromArg("ipv6") // -ipv6=5
	if arg_ipv6_rate >= 0 {
		ipv6_chance = float64(arg_ipv6_rate) / 100
	}

	var arg_port_rate int64 = intFromArg("port") // -port=5
	if arg_port_rate >= 0 {
		port_chance = float64(arg_port_rate) / 100
	}
}

// argIndex returns the index of a given argument,
// or 0 if it doesn't exist.
func argIndex(arg string) int {
	// for each argument in arguments
	for index := range os.Args[1:] {
		// get current argument from list
		var current_arg string = os.Args[1+index]

		if string(current_arg[0]) != "-" {
			// not an argument- skip!
			continue
		}

		// split argument into key:value pair
		var arg_split = strings.Split(current_arg[1:], "=")

		if arg_split[0] != arg {
			// this arg isn't what we're looking for- skip!
			continue
		}

		// argument exists! return the index
		return index + 1
	}
	return -1
}

// stringFromArg parses an argument, specified by `arg`,
// and if it exists, passes the argument's value to
// the provided variable `v`.
func stringFromArg(arg string) string {
	var index = argIndex(arg)
	if index == -1 {
		return "none"
	}

	var argument = os.Args[index]

	var arg_split = strings.Split(argument[1:], "=")

	if len(arg_split) > 1 {
		// return value of the argument!
		return arg_split[1]
	}

	// argument exists but no value exists!
	return "true"
}

// intFromArg parses argument `arg`, returning
// its value as an int, or -1 if it doesn't exist.
func intFromArg(arg string) int64 {
	var value = stringFromArg(arg)
	if value == "none" {
		return -1
	} else if value == "true" {
		return 1
	} else if value == "false" {
		return 0
	}

	response, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		// not a valid number, but the value *does* exist... passing 1 for now?
		return 1
	}
	return response
}

func readMethods() []string {
	var methods []string

	content, err := i_methods_file.Open("methods.txt")
	if err != nil {
		// fuck!!!
		log.Fatal(err)
		// return nil
	}

	// line scanner for the methods file!
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		// append new lines to the methods list!
		methods = append(methods, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		// fuck!!!
		log.Fatal(err)
		// return nil
	}

	// close the file afterwards!
	content.Close()

	return methods
}

// randomString returns a randomly-selected string from
// a provided list of `strings`, avoiding duplicates
// given a provided `last` string.
func randomString(strings []string, last string) string {
	var method string = last
	for method == last {
		method = color.Ize(color.White, strings[rand.Intn(len(strings))])
	}
	return method
}

// newAddress returns a randomly-generated IP address, IPv4 or IPv6, sometimes with a port.
//
// the chance of an IPv6 being returned is determined by `ipv6_chance`.
// the chance of a port being assigned is determined by `port_chance` and `no_ports`.
func newAddress(ipv6_chance float64, port_chance float64) string {
	var addr string

	// random chance to be IPv6 (ffe8::...)
	if rand.Float64() <= ipv6_chance {
		addr = color.Ize(color.Yellow, segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6())
	} else {
		addr = color.Ize(color.Yellow, segv4()+"."+segv4()+"."+segv4()+"."+segv4())
	}

	// random chance to have a port
	// "what are the chances of this firing with ipv6?"
	//     - loudar, seconds before getting an ipv6 with a port
	if rand.Float64() <= port_chance {
		addr += ":" + color.Ize(color.Cyan, strconv.FormatInt(rand.Int63n(65535), 10))
	}

	return addr
}

// segv4 returns an IPv4 segment, a random 8-bit value, in decimal.
//
// i.e. 0, 69, 88, 123, 127, 255...
func segv4() string {
	return strconv.FormatInt(rand.Int63n(255), 10)
}

// segv6 returns an IPv6 segment, a random 16-bit value, in hexadecimal.
//
// i.e. ff80, e396, 4d87, c388, c5f6...
func segv6() string {
	return strconv.FormatInt(rand.Int63n(65535), 16)
}
