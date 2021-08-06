package chips

import (
	"github.com/deenrookie/chips/utils"
)

type Result struct {
	StatusCode int
	Title      string
	Ip         string
	Url        string
	Domain     string
}

func SendGet(url string) Result {
	domain := utils.GetUrlDomain(url)
	ip := utils.GetDomainIp(domain)
	statusCode, body, _ := utils.HTTPGet(url)
	title := utils.GetTileFromBody(body)
	ret := Result{
		StatusCode: statusCode,
		Title:      title,
		Ip:         ip,
		Url:        url,
		Domain:     domain,
	}
	return ret
}
