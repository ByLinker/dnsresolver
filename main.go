package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"

	"github.com/miekg/dns"
)

const (
	root = "198.41.0.4"
)

// response types
const (
	ANSWER int = iota
	NS
	EXTRA
)

var (
	host = flag.String("host", "google.com", "domain name to resolve")
)

func main() {
	flag.Parse()

	_, err := url.Parse(*host)
	if err != nil {
		log.Fatal("-host is not a real url")
	}

	rr, err := resolve(*host, root)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("We found answers %s\n", *host)
	fmt.Println(rr)

}

func resolve(name, ns string) ([]dns.RR, error) {
	c := new(dns.Client)

loop:
	for {
		m := new(dns.Msg)
		m.SetQuestion(dns.Fqdn(name), dns.TypeA)

		resp, _, err := c.Exchange(m, fmt.Sprintf("%s:53", ns))
		if err != nil {
			log.Fatal(err)
		}

		switch respType(resp) {
		case ANSWER:
			return resp.Answer, nil
		case NS:
			nms := getNS(resp.Ns)
			n, err := resolve(nms, root)
			if err != nil {
				return nil, err
			}
			ns = getA(n)
			continue
		case EXTRA:
			ns = getA(resp.Extra)
			continue
		default:
			break loop
		}

	}
	return nil, errors.New("unable to find answers")
}

func respType(resp *dns.Msg) int {
	if len(resp.Answer) > 0 {
		return ANSWER
	}
	if len(resp.Extra) > 0 {
		return EXTRA
	}
	if len(resp.Ns) > 0 {
		return NS
	}
	return -1
}

func getA(rrs []dns.RR) string {
	for _, rr := range rrs {
		record, ok := rr.(*dns.A)
		if ok {
			return record.A.String()
		}
	}
	return ""
}

func getNS(rrs []dns.RR) string {
	for _, rr := range rrs {
		record, ok := rr.(*dns.NS)
		if ok {
			return record.Ns
		}
	}
	return ""
}
