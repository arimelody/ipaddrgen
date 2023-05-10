require('colors');

const methods = [
    'trolling',
    'nuking',
    'DDoSing',
    'colour-coding',
    'evaluating',
    'pissing on',
    'installing apache2 on',
    'crypto mining on',
    'teleporting to',
    'griefing',
    'fetishising',
    'publically endorsing',
    'crowdfunding',
    'noscoping',
    'load balancing',
    'expiring SSL certs on',
    'scamming',
    'http.POST(\'dickbutt\')ing on',
    'torrenting Bee Movie from',
    'defragmenting',
    'debugging',
    'lorem ipsuming',
    'telling your mother about',
    'randomly generating',
    '[meta joke]ing',
    'making amazon purchases on',
    'minting',
    'investing in',
    'starting an NFT collection using',
    'buying',
    'parking',
    'truncating',
    'deleting',
    'disconnecting',
    'blacklisting "pornhub.com" on',
    'reporting',
    'surveilling',
    'webcrawling',
    'scanning traffic from',
    '',
    'i forgot',
    'applying thermal paste to',
    'committing tax fraud with',
    'torrenting disney movies on',
    'surfing the web with',
    'LOL\'ing',
    'shitposting on',
    'wasting time on',
    'setting while loop flags for', // :(
    'livestreaming',
    'introducing the in-laws to',
    '404\'ing',
    'leaking',
    'connecting via bluetooth to',
    'obtaining the wi-fi password for',
    'you were never actually going to visit',
    'decrypting',
    'encrypting',
    'firewalling',
    'installing',
    'hacking',
    'inspecting the elements of',
    'tunelling via',
    'downloading from',
    'uploading to',
    'throttling',
    'nuclear-powering',
    'rat-infesting',
    'installing Windows Server 2008 on',
    'connecting my raspberry pi to',
    'git committing',
    'sudo rm -rf /\'ing',
    'doing your mom on',
    'balling',
    'pirating on',
    'downloading R2R software from',
    'water-cooling',
    'ejecting',
    'formatting',
    'formatting system partition on',
    'vaccinating',
    'medicating',
    'injecting',
    'pouring milk on',
    'microwaving',
    'deepfrying',
    'randomizing chance on',
    'exposing dream\'s cheating scandal from',
    'hosting dream SMP on',
    'wow look funny numbers!',
    'brapping',
    ':)',
    'jerking off',
    'rendering',
    'saving to',
    'screenshotting',
    'pinging',
    'FATAL ERROR: cannot connect to',
    'ok bud 👍',
    'streaming \'Big Mouth\' from',
    'casting to',
    'synchronising',
    'closing all ports on',
    'finding your father on',
    'SSHing to',
    'printscreening',
    'sending ominous countdown to',
    'sending pizzas to',
    'sending doordash to',
    'transcoding',
    'killing',
    'pending',
    'buffering',
    'loading',
    'hard-wiring',
    'fucking',
    'sending nudes to',
    'installing a VPN on',
    'stealing nudes from',
    'getting critical alerts from microsoft on',
    ':trollface:ing',
    'staring at',
    'PWNing',
    'downloading RAM from',
    'hey google, connect to',
    'stealing',
    'generating',
    'banning',
    'proxying',
    'racially profiling',
    'grossly offending',
    'nullnullnull'.zalgo
]

let running = true
let progress = 0
let wasIPv6 = false
let wasPorted = false
let lastMethod = ""
let method = ""

while(running) {
    
    // create ip segments
    function segv4() { return Math.round(Math.random()*255) }
    function segv6() { return Math.round(Math.random()*16**4).toString(16) }
    
    // what are we doing
    while(method == lastMethod) {
        method = methods[Math.floor(Math.random()*methods.length)]

        // check if method is '' because funny joke
        if (method == '') method = null
    }
    
    // where are we doing it
    let address = `${segv4()}.${segv4()}.${segv4()}.${segv4()}`.yellow
    
    // random chance to be IPv6
    if (Math.random() < .05 && !wasIPv6) {
        address = `${segv6()}:${segv6()}:${segv6()}:${segv6()}:${segv6()}:${segv6()}:${segv6()}:${segv6()}`.yellow
        wasIPv6 = true
    } else {
        wasIPv6 = false
    }
    
    // random chance to have a port
        // "what are the chances of this firing with ipv6?" - loudar, seconds before getting an ipv6 with a port
    if (Math.random() < .05 && !wasPorted && method != 'closing all ports on' || true) {
        address += `:${Math.floor(Math.random()*65535)}`.cyan
        wasPorted = true
    } else {
        wasPorted = false
    }

    // if (method) process.stdout.write(`${method} ${address}`)
    if (method) output(`${method} ${address}`)
    
    let complete = false
    while(!complete) {
        if (Math.random() < .9 || progress < 3) {
            // process.stdout.write('.')
            output('.')
            progress++
            
            // set complete time
            let waitUntil = new Date().getTime() + 250 + Math.round(Math.min(2**(Math.random()*12), 5000+Math.random()*1000))
            
            // wait until time has arrived
            while (new Date().getTime() < waitUntil) {}
        } else {
            // console.log('['+' ✔️ '.green+']')
            output('['+' ✔️ '.green+']\n')
            progress = 0
            complete = true
        }
    }

    lastMethod = method
   
}

/**
 * output content to...whatever you like. this is meant for printing methods, addresses and progress dots!
 * @param {string} content 
 */
function output(content) {
    process.stdout.write(content)
}