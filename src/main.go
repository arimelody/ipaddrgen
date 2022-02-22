package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/TwiN/go-color"
)

var methods = []string{
	"trolling",
	"nuking",
	"DDoSing",
	"colour-coding",
	"evaluating",
	"pissing on",
	"installing apache2 on",
	"crypto mining on",
	"teleporting to",
	"griefing",
	"fetishising",
	"publically endorsing",
	"crowdfunding",
	"noscoping",
	"load balancing",
	"expiring SSL certs on",
	"scamming",
	"http.POST('dickbutt')ing on",
	"torrenting Bee Movie from",
	"defragmenting",
	"debugging",
	"lorem ipsuming",
	"telling your mother about",
	"randomly generating",
	"[meta joke]ing",
	"making amazon purchases on",
	"minting",
	"investing in",
	"starting an NFT collection using",
	"buying",
	"parking",
	"truncating",
	"deleting",
	"disconnecting",
	"blacklisting \"pornhub.com\" on",
	"reporting",
	"surveilling",
	"webcrawling",
	"scanning traffic from",
	"",
	"i forgot",
	"applying thermal paste to",
	"committing tax fraud with",
	"torrenting disney movies on",
	"surfing the web with",
	"LOL'ing",
	"shitposting on",
	"wasting time on",
	"setting while loop flags for', // ",
	"livestreaming",
	"introducing the in-laws to",
	"404'ing",
	"leaking",
	"connecting via bluetooth to",
	"obtaining the wi-fi password for",
	"you were never actually going to visit",
	"decrypting",
	"encrypting",
	"firewalling",
	"installing",
	"hacking",
	"inspecting the elements of",
	"tunelling via",
	"downloading from",
	"uploading to",
	"throttling",
	"nuclear-powering",
	"rat-infesting",
	"installing Windows Server 2008 on",
	"connecting my raspberry pi to",
	"git committing",
	"sudo rm -rf /'ing",
	"doing your mom on",
	"balling",
	"pirating on",
	"downloading R2R software from",
	"water-cooling",
	"ejecting",
	"formatting",
	"formatting system partition on",
	"vaccinating",
	"medicating",
	"injecting",
	"pouring milk on",
	"microwaving",
	"deepfrying",
	"randomizing chance on",
	"exposing dream's cheating scandal from",
	"hosting dream SMP on",
	"wow look funny numbers!",
	"brapping",
	":)",
	"jerking off",
	"rendering",
	"saving to",
	"screenshotting",
	"pinging",
	"FATAL ERROR: cannot connect to",
	"ok bud 👍",
	"streaming 'Big Mouth' from",
	"casting to",
	"synchronising",
	"closing all ports on",
	"finding your father on",
	"SSHing to",
	"printscreening",
	"sending ominous countdown to",
	"sending pizzas to",
	"sending doordash to",
	"transcoding",
	"killing",
	"pending",
	"buffering",
	"loading",
	"hard-wiring",
	"fucking",
	"sending nudes to",
	"installing a VPN on",
	"stealing nudes from",
	"getting critical alerts from microsoft on",
	":trollface:ing",
	"staring at",
	"PWNing",
	"downloading RAM from",
	"hey google, connect to",
	"stealing",
	"generating",
	"banning",
	"proxying",
	"racially profiling",
	"grossly offending",
	"nullnullnull",
}

var running = true
var progress = 0
var wasIPv6 = false
var wasPorted = false
var lastMethod string
var method string

func segv4() string {
	return strconv.FormatInt(rand.Int63n(255), 10)
}
func segv6() string {
	return strconv.FormatInt(rand.Int63n(65535), 16)
}

func main() {

	rand.Seed(time.Now().UnixNano())

	for running {

		// what are we doing
		for method == lastMethod {
			method = color.Ize(color.White, methods[rand.Intn(len(methods))])
		}

		// where are we doing it
		var address string = color.Ize(color.Yellow, segv4()+"."+segv4()+"."+segv4()+"."+segv4())

		// random chance to be IPv6
		if rand.Intn(100) < 5 && !wasIPv6 {
			address = color.Ize(color.Yellow, segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6()+":"+segv6())
			wasIPv6 = true
		} else {
			wasIPv6 = false
		}

		// random chance to have a port
		// "what are the chances of this firing with ipv6?" - loudar, seconds before getting an ipv6 with a port
		if rand.Intn(100) < 5 && !wasPorted && method != "closing all ports on" || true {
			address += ":" + color.Ize(color.Cyan, strconv.FormatInt(rand.Int63n(65535), 10))
			wasPorted = true
		} else {
			wasPorted = false
		}

		// print ip address with method(if applicable)
		if method != " " {
			fmt.Print(method, " ")
		}
		fmt.Print(address)

		// ...........................................[ √ ]
		var complete = false
		for !complete {
			if rand.Intn(10) < 9 || progress < 3 {
				// process.stdout.write('.')
				fmt.Print(".")
				progress++

				// set complete time
				// var waitUntil = new Date().getTime() + 250 + math.round(math.min(2**(random()*12), 5000+random()*1000))
				var waitUntil = time.Now().UnixMilli() + 250 + int64(math.Min(math.Pow(2, float64(rand.Intn(12))), float64(5000+rand.Intn(1000))))

				// wait until time has arrived
				for time.Now().UnixMilli() < waitUntil {
					// do nothing
				}
			} else {
				// console.log('['+' ✔️ '.green+']')
				fmt.Print("[ " + color.Ize(color.Green, "√") + " ]\n")
				progress = 0
				complete = true
			}
		}

		lastMethod = method

	}
}
