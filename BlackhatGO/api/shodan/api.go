package shodan

type APIInfo struct {
	queryCredits int 	`json:"query_credits"`
	ScanCredits  int 	`json:"scan_credits"`
	Telnet 	     bool 	`json:"telnet"`
	Plan 		 string	`json:"plan"`
	HTTPS 		 bool	`json:"https"`
	Unlocked 	 bool 	`json:"unlocked"`
}

func (s *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err 
	}
	defer res.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err 
	}
	return &ret, nil
}