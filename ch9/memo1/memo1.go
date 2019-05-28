// 并发非阻塞缓存系统
// 第一获取url之后,会缓存结果
// 第二次获取同个url时,会直接从结果返回
package memo1

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var cache = make(map[string]interface{})

func GetUrl(url string) (interface{}, error) {
	if v, ok := cache[url]; ok {
		fmt.Println("from cache")
		return v, nil
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
	}

	nbyte, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		cache[url] = nbyte
	}

	return nbyte, err
}
