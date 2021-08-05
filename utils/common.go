package utils

import (
	"compress/gzip"
	"context"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

func GetDomainIp(domain string) string {
	addr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		return ""
	}
	return addr.String()
}

func GetUrlDomain(targetUrl string) string {
	u, err := url.Parse(targetUrl)
	if err != nil {
		return ""
	}
	return u.Hostname()
}

func HTTPGet(remoteUrl string) (statusCode int, body []byte, err error) {
	client := &http.Client{}
	body = nil
	var response *http.Response
	uri, err := url.Parse(remoteUrl)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	request, err := http.NewRequest("GET", remoteUrl, nil)

	if request != nil {
		request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
		request.Header.Add("Accept-Encoding", "gzip, deflate")
		request.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3")
		request.Header.Add("Connection", "keep-alive")
		request.Header.Add("Host", uri.Host)
		request.Header.Add("Referer", uri.String())
		request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0")
		response, err = client.Do(request.WithContext(ctx))
	}

	if response != nil {
		defer func() {
			_ = response.Body.Close()
		}()
	} else {
		return
	}

	if err != nil {
		return
	}

	statusCode = response.StatusCode

	if response.StatusCode == 200 {
		switch response.Header.Get("Content-Encoding") {
		case "gzip":
			reader, _ := gzip.NewReader(response.Body)
			for {
				buf := make([]byte, 1024)
				n, newErr := reader.Read(buf)

				if newErr != nil && newErr != io.EOF {
					return
				}

				if n == 0 {
					break
				}
				body = append(body, buf...)
			}
		default:
			body, _ = ioutil.ReadAll(response.Body)
		}
	}
	return
}

func GetTileFromBody(body []byte) string {
	regex, _ := regexp.Compile("<title>([\\s\\S]*?)</title>")
	title := string(regex.Find(body))
	re, _ := regexp.Compile("<[\\S\\s]+?>")
	title = re.ReplaceAllString(title, "")
	return title
}
