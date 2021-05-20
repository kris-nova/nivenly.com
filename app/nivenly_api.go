package nivenly

import (
	"net/http"

	"github.com/kris-nova/logger"
)

// NivenlyAPI is a request specific API for the website.
//
// All user input that comes into the website comes in
// the form of an HTTP request, therefor regardless of
// the state of the backend, we can get everything we
// need via an HTTP request.
//
// Here is the API.
//
// Add fields carefully!
//
// This is concretely what is referenced in the "params"
// in the config.toml
//
// If you can get data here - you can interpolate into the website.
type NivenlyAPI struct {

	// ClientAddr
	// 1 if by proxy
	// 2 if by request
	//
	// Here is where we define exactly "who" our friendly visitor is.
	//
	// This is the most critical string as it also serves as the "key"
	// in our cache in memory.
	ClientAddr string

	// Client is the nested *Client which is the same data structure
	// that can be found at the /client endpoint.
	Client Client

	// R is the *http.Request for EACH request to the server.
	R *http.Request
}

// GetAPI is the method that is called for every HTTP request to the website.
//
// Break out by client here.
func (v *Nivenly) GetAPI(r *http.Request) *NivenlyAPI {
	// Every request MUST have a client!
	// We always have to be able to get here or else
	// we are putting something VERY wrong in our code path.
	//
	// Think about it - every HTTP request MUST have a client.
	// Our code should be resilient enough to support whatever
	// we get from the server.
	client := v.clientHandler.GetClient(r)

	// Log client for every interpolation
	logger.Info("Client: %s", client.Addr)

	// Build our API for the site
	api := &NivenlyAPI{
		Client:     client,
		ClientAddr: client.Addr,
		R:          r,
	}
	return api
}