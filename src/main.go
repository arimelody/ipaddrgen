package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/TwiN/go-color"
)

// === internal vars (no need to edit these!) ===

var running bool
var i_was_IPv6 = false
var i_was_ported = false
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

// chance (0.0 - 1.0) for the operation to complete each update.
var complete_rate = 0.1

// chance (0.0 - 1.0) for the random IP to be IPv6.
var ipv6_chance = 0.05

// chance (0.0 - 1.0) for the random IP to have a port.
var port_chance = 0.05

func main() {

	// read the methods file!
	content, err := i_methods_file.Open("methods.txt")
	if err != nil {
		// fuck!!!
		log.Fatal(err)
		return
	}

	// line scanner for the methods file!
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		// append new lines to the methods list!
		i_methods = append(i_methods, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		// fuck!!!
		log.Fatal(err)
		return
	}

	// close the file afterwards!
	content.Close()

	rand.Seed(time.Now().UnixNano())

	var waitdelta = max_wait - min_wait

	running = true
	for running {

		// === create a new IP to do the funny to ===

		// what are we doing (method)
		var method = newMethod()

		// where are we doing it (address)
		var address string = newAddress(
			ipv6_chance, port_chance,
			method != "opening all ports on" &&
				method != "closing all ports on")

		// === print the damn thing! ===

		// > [doing something to] 127.0.0.1...
		if method != " " {
			fmt.Print(method, " ")
		}
		fmt.Print(address)

		// and now, the slow, painful wait begins...
		var progress = 0
		for {
			// if complete check fails, and we have over 3 wait periods...
			// print periods at random intervals, wasting the user's time
			if rand.Float64() > complete_rate || progress < 3 {
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

				continue
			}

			// [ √ ]
			fmt.Print("[ " + color.Ize(color.Green, "√") + " ]\n")

			break
		}

	}
}

func newMethod() string {
	var method string = i_last_method
	for method == i_last_method {
		method = color.Ize(color.White, i_methods[rand.Intn(len(i_methods))])
	}
	i_last_method = method
	return method
}

// newAddress returns a randomly-generated IP address, IPv4 or IPv6, sometimes with a port.
//
// the chance of an IPv6 being returned is determined by ipv6_chance.
// the chance of a port being assigned is determined by port_chance and no_ports.
func newAddress(ipv6_chance float64, port_chance float64, no_ports bool) string {
	var addr string

	// random chance to be IPv6 (ffe8::...)
	if rand.Float64() > ipv6_chance {
		addr = color.Ize(color.Yellow, segv4()+"."+segv4()+"."+segv4()+"."+segv4())
	} else {
		addr = color.Ize(color.Yellow, segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6())
	}

	// random chance to have a port
	// "what are the chances of this firing with ipv6?"
	//     - loudar, seconds before getting an ipv6 with a port
	if rand.Float64() < port_chance && !no_ports {
		addr += ":" + color.Ize(color.Cyan, strconv.FormatInt(rand.Int63n(65535), 10))
	}

	return addr
}

// returns an IPv4 segment, a random 8-bit value, in decimal.
//
// i.e. 0, 69, 88, 123, 127, 255...
func segv4() string {
	return strconv.FormatInt(rand.Int63n(255), 10)
}

// returns an IPv6 segment, a random 16-bit value, in hexadecimal.
//
// i.e. ff80, e396, 4d87, c388, c5f6...
func segv6() string {
	return strconv.FormatInt(rand.Int63n(65535), 16)
}
