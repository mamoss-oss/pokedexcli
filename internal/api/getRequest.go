package api

import (
	"io"
	"log"
	"net/http"

	"github.com/mamoss-oss/pokedexcli/internal/pokecache"
)

func GetRequest(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
		return []byte{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 { // everything starting from code 300 is not succesful
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return body, err
}

func CacheOrGet(url string, cache *pokecache.Cache) ([]byte, error) {
	data, res := cache.Get(url)
	if res {
		return data, nil
	} else {
		data, err := GetRequest(url)
		cache.Add(url, data)
		return data, err
	}
}
