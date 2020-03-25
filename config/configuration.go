package config

import (
	"encoding/json"
)

type Configuration struct {
	Restapi     Restapi     `json:"restapi,omitempty"`
	Streamapi   Streamapi   `json:"streamapi,omitempty"`
	Exploitsapi Exploitsapi `json:"exploitsapi,omitempty"`
}

func LoadConfig() Configuration {
	var config Configuration
	data := `{
		"restapi": {
		  "search": {
			"HostInformationURL": "https://api.shodan.io/shodan/host/{ip}",
			"HostCountURL": "https://api.shodan.io/shodan/host/count",
			"HostSearchURL": "https://api.shodan.io/shodan/host/search",
			"HostSearchTokensURL": "https://api.shodan.io/shodan/host/search/tokens",
			"PortsURL": "https://api.shodan.io/shodan/ports"
		  },
		  "scan": {
			"ProtocolsURL": "https://api.shodan.io/shodan/protocols",
			"ScanURL": "https://api.shodan.io/shodan/scan?key={key}",
			"ScanInternet": "https://api.shodan.io/shodan/scan/internet",
			"ScanStatusURL": "https://api.shodan.io/shodan/scan/{id}"
		  },
		  "networkAlerts": {
			"AlertURL": "https://api.shodan.io/shodan/alert",
			"AlertIdInfoURL": "https://api.shodan.io/shodan/alert/{id}/info",
			"DeleteAlertURL": "https://api.shodan.io/shodan/alert/{id}",
			"AlertInfoURL": "https://api.shodan.io/shodan/alert/info",
			"AlertTriggersURL": "https://api.shodan.io/shodan/alert/triggers",
			"EnableTriggerURL": "https://api.shodan.io/shodan/alert/{id}/trigger/{trigger}",
			"DisableTriggerURL": "https://api.shodan.io/shodan/alert/{id}/trigger/{trigger}",
			"AddTriggerToWhitelistURL": "https://api.shodan.io/shodan/alert/{id}/trigger/{trigger}/ignore/{service}",
			"RemoveFromWhitelistURL": "https://api.shodan.io/shodan/alert/{id}/trigger/{trigger}/ignore/{service}"
		  },
		  "directory": {
			"QueryURL": "https://api.shodan.io/shodan/query",
			"QuerySearchURL": "https://api.shodan.io/shodan/query/search",
			"QueryTagsURL": "https://api.shodan.io/shodan/query/tags"
		  },
		  "data": {
			"DataURL": "https://api.shodan.io/shodan/data",
			"DatasetURL": "https://api.shodan.io/shodan/data/raw-daily"
		  },
		  "organization": {
			"InfoURL": "https://api.shodan.io/org",
			"AddNewMemberURL": "https://api.shodan.io/org/member/{user}",
			"RemoveMemberURL": "https://api.shodan.io/org/member/{user}"
		  },
		  "account": {
			"AccountProfileURL": "https://api.shodan.io/account/profile"
		  },
		  "dns": {
			"DomainInfoURL": "https://api.shodan.io/dns/domain/{domain}",
			"DNSLookupURL": "https://api.shodan.io/dns/resolve",
			"ReverseDNSLookupURL": "https://api.shodan.io/dns/reverse"
		  },
		  "utility": {
			"HTTPHeadersURL": "https://api.shodan.io/tools/httpheaders",
			"MyIPAddressURL": "https://api.shodan.io/tools/myip"
		  },
		  "apistatus": {
			"APIPlanInformationURL": "https://api.shodan.io/api-info"
		  },
		  "experimental": {
			"HoneyscoreURL": "https://api.shodan.io/labs/honeyscore/{ip}"
		  }
		},
		"streamapi": {
		  "streamData": {
			"BannersURL": "https://stream.shodan.io/shodan/banners",
			"FiltereByASNURL":"https://stream.shodan.io/shodan/asn/{asn}",
			"FiltereByCountry":"https://stream.shodan.io/shodan/countries/{countries}",
			"FiltereByPorts":"https://stream.shodan.io/shodan/ports/{ports}"
		  },
		  "streamNetworkAlerts": {
			"AllNetworkAlertsURL":"https://stream.shodan.io/shodan/alert",
			"FiltereByAlertIDURL":"https://stream.shodan.io/shodan/alert/{id}"
		  }
		},
		"exploitsapi": {
		  "SearchURL":"https://exploits.shodan.io/api/search",
		  "CountURL":"https://exploits.shodan.io/api/count"
		}
	  }`

	json.Unmarshal([]byte(data), &config)

	return config
}

type Exploitsapi struct {
	SearchURL string `json:"SearchURL,omitempty"`
	CountURL  string `json:"CountURL,omitempty"`
}

type Restapi struct {
	Search        Search        `json:"search,omitempty"`
	Scan          Scan          `json:"scan,omitempty"`
	NetworkAlerts NetworkAlerts `json:"networkAlerts,omitempty"`
	Directory     Directory     `json:"directory,omitempty"`
	Data          Data          `json:"data,omitempty"`
	Organization  Organization  `json:"organization,omitempty"`
	Account       Account       `json:"account,omitempty"`
	DNS           DNS           `json:"dns,omitempty"`
	Utility       Utility       `json:"utility,omitempty"`
	Apistatus     Apistatus     `json:"apistatus,omitempty"`
	Experimental  Experimental  `json:"experimental,omitempty"`
}

type Account struct {
	AccountProfileURL string `json:"AccountProfileURL,omitempty"`
}

type Apistatus struct {
	APIPlanInformationURL string `json:"APIPlanInformationURL,omitempty"`
}

type DNS struct {
	DomainInfoURL       string `json:"DomainInfoURL,omitempty"`
	DNSLookupURL        string `json:"DNSLookupURL,omitempty"`
	ReverseDNSLookupURL string `json:"ReverseDNSLookupURL,omitempty"`
}

type Data struct {
	DataURL    string `json:"DataURL,omitempty"`
	DatasetURL string `json:"DatasetURL,omitempty"`
}

type Directory struct {
	QueryURL       string `json:"QueryURL,omitempty"`
	QuerySearchURL string `json:"QuerySearchURL,omitempty"`
	QueryTagsURL   string `json:"QueryTagsURL,omitempty"`
}

type Experimental struct {
	HoneyscoreURL string `json:"HoneyscoreURL,omitempty"`
}

type NetworkAlerts struct {
	AlertURL                 string `json:"AlertURL,omitempty"`
	AlertIDInfoURL           string `json:"AlertIdInfoURL,omitempty"`
	DeleteAlertURL           string `json:"DeleteAlertURL,omitempty"`
	AlertInfoURL             string `json:"AlertInfoURL,omitempty"`
	AlertTriggersURL         string `json:"AlertTriggersURL,omitempty"`
	EnableTriggerURL         string `json:"EnableTriggerURL,omitempty"`
	DisableTriggerURL        string `json:"DisableTriggerURL,omitempty"`
	AddTriggerToWhitelistURL string `json:"AddTriggerToWhitelistURL,omitempty"`
	RemoveFromWhitelistURL   string `json:"RemoveFromWhitelistURL,omitempty"`
}

type Organization struct {
	InfoURL         string `json:"InfoURL,omitempty"`
	AddNewMemberURL string `json:"AddNewMemberURL,omitempty"`
	RemoveMemberURL string `json:"RemoveMemberURL,omitempty"`
}

type Scan struct {
	ProtocolsURL  string `json:"ProtocolsURL,omitempty"`
	ScanURL       string `json:"ScanURL,omitempty"`
	ScanInternet  string `json:"ScanInternet,omitempty"`
	ScanStatusURL string `json:"ScanStatusURL,omitempty"`
}

type Search struct {
	HostInformationURL  string `json:"HostInformationURL,omitempty"`
	HostCountURL        string `json:"HostCountURL,omitempty"`
	HostSearchURL       string `json:"HostSearchURL,omitempty"`
	HostSearchTokensURL string `json:"HostSearchTokensURL,omitempty"`
	PortsURL            string `json:"PortsURL,omitempty"`
}

type Utility struct {
	HTTPHeadersURL string `json:"HTTPHeadersURL,omitempty"`
	MyIPAddressURL string `json:"MyIPAddressURL,omitempty"`
}

type Streamapi struct {
	StreamData          StreamData          `json:"streamData,omitempty"`
	StreamNetworkAlerts StreamNetworkAlerts `json:"streamNetworkAlerts,omitempty"`
}

type StreamData struct {
	BannersURL          string `json:"BannersURL,omitempty"`
	FiltereByASNURL     string `json:"FiltereByASNURL,omitempty"`
	FiltereByCountryURL string `json:"FiltereByCountry,omitempty"`
	FiltereByPortsURL   string `json:"FiltereByPorts,omitempty"`
}

type StreamNetworkAlerts struct {
	AllNetworkAlertsURL string `json:"AllNetworkAlertsURL,omitempty"`
	FiltereByAlertIDURL string `json:"FiltereByAlertIDURL,omitempty"`
}
