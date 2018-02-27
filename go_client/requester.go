package tor

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/proxy"
)

type Torlet struct {
	control    *controller
	httpClient *http.Client
	Ip         string
	auth       *proxy.Auth
}

func NewTorlet(config *Config) *Torlet {
	return &Torlet{newController(config.ControlAddress, config.ControlPassword), &http.Client{}, config.SocksAddress, &proxy.Auth{}}
}

func (t *Torlet) Init() error {
	transport := &http.Transport{}
	dialer, err := proxy.SOCKS5("tcp", t.Ip, t.auth, proxy.Direct)
	if err != nil {
		return err
	}
	transport.Dial = dialer.Dial
	t.httpClient = &http.Client{Transport: transport}
	log.Println(fmt.Sprintf("Tor connected, IP - %s", t.CheckIP()))
	return nil
}

func (t *Torlet) ResetCircuit() error {
	log.Println("Reseting tor circuit")
	err := t.control.dial()
	if err != nil {
		return err
	}
	err = t.control.resetCircuit()
	if err != nil {
		return err
	}
	t.control.close()
	log.Println("Tor circuit reset, IP - %s", t.CheckIP())
	return nil
}

func (t *Torlet) Request(req *http.Request) (*http.Response, error) {
	return t.httpClient.Do(req)
}

func (t *Torlet) Close() {
	t.control.close()
}

func (t *Torlet) CheckEndpointIP() string {
	iUrl := "http://icanhazip.com/"
	req, err := http.NewRequest("GET", iUrl, nil)
	if err != nil {
		return err.Error()
	}

	resp, err := t.Request(req)
	if err != nil {
		return err.Error()
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(b)
}
