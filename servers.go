package generic_utils

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/bruno-anjos/edge-deployment-utils/http_utils"
	log "github.com/sirupsen/logrus"
)

const (
	// DefaultHost contains the default interface address
	DefaultHost = "0.0.0.0"
)

// StartServer seeds the random generator and starts a server on the
// specified host and port serving the routes passed.
func StartServer(serviceName, host string, port int, routes []http_utils.Route) {
	rand.Seed(time.Now().UnixNano())
	addr := fmt.Sprintf("%s:%d", host, port)
	r := http_utils.NewRouter(routes)
	log.Infof("starting %s server in port %d...\n", serviceName, port)
	log.Fatal(http.ListenAndServe(addr, r))
}
