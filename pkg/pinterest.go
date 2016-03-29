package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Pinterest returns a Platform specific to Pinterest
func Pinterest() Platform {
	return Platform{
		enabled:  true,
		name:     "pinterest",
		statsURL: "http://api.pinterest.com/v1/urls/count.json?callback=call&url=%s",
		format:   "jsonp",
		parseWith: func(r *http.Response) (Stat, error) {
			body, error := ioutil.ReadAll(r.Body)
			if error != nil {
				return Stat{}, error
			}

			jsonBody, error := parseJSONP(body)
			if error != nil {
				return Stat{}, error
			}

			var jsonBlob map[string]interface{}
			if err := json.Unmarshal([]byte(jsonBody), &jsonBlob); err != nil {
				return Stat{}, err
			}

			return Stat{
				data: map[string]interface{}{"count": jsonBlob["count"]},
			}, nil
		}}
}
