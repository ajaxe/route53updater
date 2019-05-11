package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ajaxe/route53updater/pkg/logging"
	"github.com/ajaxe/route53updater/pkg/shared"
)

var ipCheckerURL = "http://checkip.amazonaws.com/"

func main() {
	fs := flag.NewFlagSet("updater", flag.ExitOnError)
	lambdaURL := fs.String("url", "", "lambda url to invoke")
	psk := fs.String("psk", "", "pre shared key used to identify the caller to the lambda")
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: route53updater [options]

Invokes updater lambda and forwards the public IP of the host.

Options are:

`)
		fs.PrintDefaults()
		os.Exit(1)
	}
	fs.Parse(os.Args[1:])
	if fs.NArg() != 0 {
		fs.Usage()
		os.Exit(1)
	}

	cfg := &config{
		URL: *lambdaURL,
		PSK: *psk,
	}

	if fs.Parsed() {
		if len(cfg.URL) == 0 {
			logging.DBGLogger.Print("URL parameter is required")
			return
		}
		if len(cfg.PSK) == 0 {
			logging.DBGLogger.Print("PSK parameter is required")
			return
		}
	}
	ip, _ := getPublicIP()

	invokeUpdateLambda(cfg, ip)
}

func invokeUpdateLambda(cfg *config, IP string) error {
	n := getNonce()
	p := shared.Payload{
		IP:      IP,
		Nonce:   n,
		HashKey: getHash(n, cfg),
	}
	logging.DBGLogger.Println("ip:", IP)
	b, _ := json.MarshalIndent(&p, "", "  ")
	logging.DBGLogger.Println(string(b))
	_, err := http.Post(cfg.URL, "application/json", bytes.NewReader(b))
	return err
}

func getPublicIP() (string, error) {
	resp, err := http.Get(ipCheckerURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	ip := string(b)
	if len(ip) == 0 {
		return "", fmt.Errorf("unable to parse ip string")
	}
	return strings.Trim(ip, "\n"), nil
}

func getNonce() string {
	return fmt.Sprintf("%d", time.Now().Unix())
}

func getHash(nonce string, cfg *config) string {
	sum := sha256.Sum256([]byte(fmt.Sprintf("%s%s", nonce, cfg.PSK)))
	return fmt.Sprintf("%x", sum)
}
