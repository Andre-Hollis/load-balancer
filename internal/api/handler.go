package api

import (
	"net/http"

	"github.com/Andre-Hollis/load-balancer/internal/model"
)

func Handler(w http.ResponseWriter, r *http.Request, lb *model.LoadBalancer, servers []*model.Server) {
	server := lb.GetNextServer(servers)
	if server == nil {
		http.Error(w, "No healthy server available", http.StatusServiceUnavailable)
		return
	}

	// adding this header just for checking from which server the request is being handled.
	// this is not recommended from security perspective as we don't want to let the client know which server is handling the request.
	w.Header().Add("X-Forwarded-Server", server.URL.String())
	server.ReverseProxy().ServeHTTP(w, r)
}
