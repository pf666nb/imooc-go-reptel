package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.75.14 (KHTML, like Gecko) Version/7.0.3 Safari/537.75.14") //刚才找到的浏览器中User-agent

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	//defer resp.Body.Close()
	//
	//resp, err := http.Get(url)
	//if err != nil {
	//	return nil, err
	//}
	//defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong state code : %d",
			resp.StatusCode)

	}
	return ioutil.ReadAll(resp.Body)

}
