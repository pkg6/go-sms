package gosms

import "github.com/pkg6/go-requests"

var Client *requests.Client

func init() {
	Client = requests.New()
}
