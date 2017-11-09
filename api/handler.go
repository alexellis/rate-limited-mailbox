package function

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

// Handle a serverless request
func Handle(req []byte) string {
	status := "OK"
	if string(req) == "reset" {
		ioutil.WriteFile("/tmp/rate_limit", []byte("0"), 0600)

	} else {
		if _, err := os.Stat("/tmp/rate_limit"); err != nil {
			ioutil.WriteFile("/tmp/rate_limit", []byte("0"), 0600)
		}

		max := 5
		data, readErr := ioutil.ReadFile("/tmp/rate_limit")
		if readErr != nil {
			return string(readErr.Error())
		}

		val, convErr := strconv.Atoi(string(data))
		if convErr != nil {
			return string(convErr.Error())
		}

		if val >= max {
			status = "FAIL"
		} else {
			val = val + 1
			ioutil.WriteFile("/tmp/rate_limit", []byte(fmt.Sprintf("%d", val)), 0600)
		}
	}

	return fmt.Sprintf(status)
}
