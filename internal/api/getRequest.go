package api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/mamoss-oss/pokedexcli/internal/pokecache"
)

func GetRequest(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 { // everything starting from code 300 is not succesful
		err = fmt.Errorf("get request received status code %d for url:%s", res.StatusCode, url)
	}
	return body, err
}

func CacheOrGet(url string, cache *pokecache.Cache) ([]byte, error) {
	data, res := cache.Get(url)
	if res {
		return data, nil
	} else {
		data, err := GetRequest(url)
		if err != nil {
			return []byte{}, err
		}
		cache.Add(url, data)
		return data, err
	}
}
