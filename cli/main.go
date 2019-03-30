package main

import (
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ajaxe/route53updater/cli/shared"
)

var ipCheckerURL = "http://checkip.amazonaws.com/"

func main() {
	fs := flag.NewFlagSet("updater", flag.ExitOnError)
	lambdaURL := fs.String("url", "", "lambda url to invoke")
	psk := fs.String("psk", "", "pre shared key used to identify the caller to the lambda")
	clientID := fs.String("client-id", "", "client id")
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
	}

	cfg := &config{
		URL:      *lambdaURL,
		PSK:      *psk,
		ClientID: *clientID,
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
	fmt.Println("ip:", IP)
	b, _ := json.MarshalIndent(&p, "", "  ")
	fmt.Println(string(b))
	return nil
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
