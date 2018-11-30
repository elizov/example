package handlers

import (
	"fmt"
	"net"
	"net/http"

	"app/entities"
	"app/repositories/geo"

	"github.com/mholt/binding"
	UserAgent "github.com/mssola/user_agent"
)

// Handler - main handler
func Handler(resp http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		user := new(entities.User)
		if err := binding.Bind(req, user); err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		// User agent
		ua := UserAgent.New(req.UserAgent())
		browser, version := ua.Browser()

		// Location
		geoIP := geo.GetGeoIP()
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}
		location, err := geoIP.Country(net.ParseIP(ip))
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(resp, "ID:    %d\n", user.ID)
		fmt.Fprintf(resp, "Email: %s\n", user.Email)
		fmt.Fprintf(resp, "Name:  %s\n", user.Name)
		fmt.Fprintf(resp, "OS:  %s\n", ua.OS())
		fmt.Fprintf(resp, "Browser:  %s (%s)\n", browser, version)
		fmt.Fprintf(resp, "Is Mobile:  %t\n", ua.Mobile())
		fmt.Fprintf(resp, "Country:  %v\n", location.Country.Names["en"])
	}
}
