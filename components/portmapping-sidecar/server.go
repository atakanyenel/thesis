package portmapper

import (
	"fmt"
	"net/http"
)

type individualServer struct {
	Mux                   *http.ServeMux
	portRange             []int
	registeredServerPorts []int
}

var portMap map[string]individualServer

/*
3000- > 3100 - 3200
4000 -> 4100 - 4200
5000 -> 5100 - 5200
*/
func init() {
	portMap = map[string]individualServer{}
	portStart, portEnd := 3100, 3200
	for _, port := range []string{":3000", ":4000", ":5000"} {
		portMap[port] = individualServer{
			Mux:                   http.NewServeMux(),
			portRange:             []int{portStart, portEnd},
			registeredServerPorts: []int{},
		}
		portStart += 1000
		portEnd += 1000
	}
}

//Start starts the portmapper server
func Start() {

	for p, server := range portMap {

		server.Mux.HandleFunc("/", portRangeHandler(server.portRange))

		fmt.Printf("Starting server with port %s\n", p)
		go http.ListenAndServe(p, server.Mux)
	}

}

func portRangeHandler(portRange []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Forwarding requests between %v", portRange)
		RandInt
	}
}
