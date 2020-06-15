package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
)

//var rateLimiter = time.Tick(500 * time.Millisecond)
var rateLimiter = time.Tick(time.Second)

func Fetch(urlStr string) ([]byte, error) {
	cookieStr := `bid=aTD1thdDJXc; ll="108090"; __gads=ID=ce70b0fdecc7c0a2:T=1586919054:S=ALNI_MboWkfgEQpwjCrvzyeMVryQqbAofQ; _vwo_uuid_v2=D1BC8929403C771AEABED2939F5C4E2F3|7926727588e25dc1edfb7fdf66e88c71; douban-fav-remind=1; __yadk_uid=4DCfA2VucFudftl0pciyemjQFFikfGZh; gr_user_id=c13abce1-2b3d-4fe8-bcd8-12d1a82f6f59; __utmz=30149280.1591345466.10.7.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; __utmc=30149280; ct=y; push_noty_num=0; push_doumail_num=0; __utmv=30149280.18037; refer_url=https://read.douban.com/reader/ebook/52836811/; _pk_ref.100001.8cb4=%5B%22%22%2C%22%22%2C1592184111%2C%22https%3A%2F%2Fwww.baidu.com%2Flink%3Furl%3DPtexJ3GNN3alytCe8RXAZEyC3rbjKNmksxDggT1hj68zHAmzHAVfsm_tBrCtKZDu%26wd%3D%26eqid%3Df9eef6d20000d2d2000000065ee5da8b%22%5D; _pk_ses.100001.8cb4=*; ap_v=0,6.0; __utma=30149280.1326003882.1586919054.1592126973.1592184129.28; __utmt=1; _pk_id.100001.8cb4=8c757f4462c2f084.1587005294.26.1592184178.1592126953.; amplitude_id_476d446d2d57aa9a7764732a26466a97douban.com=eyJkZXZpY2VJZCI6ImQ2MmUzZTgwLTg0ZDctNGY4MC04YTk2LWQ4MzlhMjM5ZTBhOVIiLCJ1c2VySWQiOiI4MDliYmU1Ny1hYThjLTQ4NGQtOWM5NS1lNTc4NjRhY2UxNWEiLCJvcHRPdXQiOmZhbHNlLCJzZXNzaW9uSWQiOjE1OTIxODQwMzQ4MzgsImxhc3RFdmVudFRpbWUiOjE1OTIxODQxNzc4NjMsImV2ZW50SWQiOjEwMywiaWRlbnRpZnlJZCI6NTU1LCJzZXF1ZW5jZU51bWJlciI6NjU4fQ==; __utmb=30149280.37.4.1592184178265`
	<-rateLimiter
	req, err := http.NewRequest(http.MethodGet, urlStr, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-agent", browser.Random())
	req.AddCookie(&http.Cookie{Name: cookieStr})
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
