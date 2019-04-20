package providers

import (
	"errors"
	"net/http"
	"strconv"
)

func HandleStatus(res *http.Response) error {
	return errors.New("status code error:" + strconv.Itoa(res.StatusCode) + " " + res.Status)
}
