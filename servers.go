package generic_utils

import (
	"flag"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/bruno-anjos/solution-utils/http_utils"
	log "github.com/sirupsen/logrus"
)

const (
	// LocalhostAddr contains the default interface address
	LocalhostAddr = "127.0.0.1"
)

// StartServer seeds the random generator and starts a server on the
// specified host and port serving the routes passed with a specified prefix.
func StartServer(serviceName, hostPort string, port int, prefixPath string, routes []http_utils.Route) {
	rand.Seed(time.Now().UnixNano())

	debug := flag.Bool("d", false, "add debug logs")
	listenAddr := flag.String("l", LocalhostAddr, "address to listen on")
	flag.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}

	log.Debug("starting log in debug mode")
	r := http_utils.NewRouter(prefixPath, routes)

	var listenAddrPort string
	if *listenAddr != "" {
		listenAddrPort = *listenAddr + ":" + strconv.Itoa(port)
	} else {
		listenAddrPort = hostPort
	}

	log.Infof("%s server listening at %s...\n", serviceName, listenAddrPort)
	log.Fatal(http.ListenAndServe(listenAddrPort, r))
}
