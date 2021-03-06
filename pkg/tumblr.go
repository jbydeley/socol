package collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Tumblr returns a Platform specific to Tumblr
func Tumblr() Platform {
	return Platform{
		enabled:  true,
		name:     "tumblr",
		format:   "json",
		statsURL: "http://api.tumblr.com/v2/share/stats?url=%s",
		parseWith: func(r *http.Response) (Stat, error) {
			body, error := ioutil.ReadAll(r.Body)
			if error != nil {
				return Stat{}, error
			}

			count := 0.0
			var jsonBlob map[string]interface{}
			if err := json.Unmarshal(body, &jsonBlob); err != nil {
				return Stat{}, err
			}
			count = jsonBlob["response"].(map[string]interface{})["note_count"].(float64)

			return Stat{
				data: map[string]interface{}{"count": count},
			}, nil
		},
	}
}
