package ip2location

import (
	"encoding/json"
	"testing"

	"github.com/ax-i-om/bitcrook/internal/config"
	"github.com/ax-i-om/bitcrook/internal/http"
)

func TestDomainLookup(t *testing.T) {
	elConf, err := config.LoadConfig("../../../keyconfig.json")
	if err != nil {
		t.Error(err)
	}
	resp, err := http.GetReq("https://api.ip2whois.com/v2?key=" + elConf.Ip2LocationKey + "&domain=bitcrook.tech&format=json")
	if err != nil {
		t.Error(err)
	}
	response := new(DomainData)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		t.Error(err)
	}
	if response.Domain != "bitcrook.tech" {
		t.Error()
	}
}

func TestIpLookup(t *testing.T) {
	elConf, err := config.LoadConfig("../../../keyconfig.json")
	if err != nil {
		t.Error(err)
	}
	resp, err := http.GetReq("https://api.ip2location.io/?key=" + elConf.Ip2LocationKey + "&ip=1.1.1.1&format=json")
	if err != nil {
		t.Error(err)
	}
	response := new(IpData)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		t.Error(err)
	}
}
