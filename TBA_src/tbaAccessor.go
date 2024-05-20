package TBA_src

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func AccessTBA(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	req.Header.Add("X-TBA-Auth-Key", "euhCCAEE2wTJnK9bRqP2Ng1bOAqfvQ1xXUrPdhV651N9rg1AuZbGjQXjfGQ73yaJ")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return body, nil
}
