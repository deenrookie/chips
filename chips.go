package chips

import (
	"github.com/deenrookie/chips/results"
	"github.com/deenrookie/chips/utils"
)

func SendGet(url string) results.Result {
	domain := utils.GetUrlDomain(url)
	ip := utils.GetDomainIp(domain)
	statusCode, body, _ := utils.HTTPGet(url)
	title := utils.GetTileFromBody(body)
	ret := results.Result{
		StatusCode: statusCode,
		Title:      title,
		Ip:         ip,
		Url:        url,
		Domain:     domain,
	}
	return ret
}
