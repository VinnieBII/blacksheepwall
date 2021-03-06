package bsw

import (
	"os"
	"strings"
	"testing"
)

func TestViewDNSInfoAPI(t *testing.T) {
	key := os.Getenv("VIEWDNS_API_KEY")
	if key == "" {
		t.Fatal("Can not test ViewDNSInfoAPI with out api key in evironment variable VIEWDNS_API_KEY")
	}
	tsk, results, err := ViewDNSInfoAPI("104.131.56.170", key)
	if tsk != "viewdns.info API" {
		t.Error("task for ViewDNSInfoAPI not viewdns.info API")
	}
	if err != nil {
		t.Error("error returned from ViewDNSInfoAPI")
		t.Log(err)
	}
	if len(results) < 1 {
		t.Error("no results returned from ViewDNSInfoAPI")
	}

	found := false
	for _, r := range results {
		if strings.Contains(r.Hostname, "stacktitan.com") {
			found = true
		}
	}
	if !found {
		t.Error("no results were correct for ViewDNSInfoAPI")
	}
}
