package utils

import "net/url"

func Valid(rawURL string) bool {
	u, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	// must have a scheme (http/https) and a host (example.com)
	if u.Scheme == "" || u.Host == "" {
		return false
	}

	if u.Scheme != "http" && u.Scheme != "https" {
		return false
	}

	return true
}
