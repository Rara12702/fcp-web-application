package client

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func GetClientWithCookie(token string, cookies ...*http.Cookie) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	cookies = append(cookies, &http.Cookie{
		Name:  "session_token",
		Value: token,
	})

	jar.SetCookies(&url.URL{
		Scheme: "https",
		Host:   "fcp-web-application-production-c9fc.up.railway.app",
	}, cookies)

	c := &http.Client{
		Jar: jar,
	}

	return c, nil
}
