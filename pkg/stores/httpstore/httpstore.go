package httpstore

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/brendoncarroll/webfs/pkg/stores"
)

const MaxBlobSize = 1 << 16

type HttpStore struct {
	endpoint    string
	prefix      string
	maxBlobSize int
}

func New(endpoint string, prefix string) *HttpStore {
	if len(endpoint) > 0 && endpoint[len(endpoint)-1] != '/' {
		endpoint += "/"
	}
	return &HttpStore{
		endpoint:    endpoint,
		maxBlobSize: MaxBlobSize,
		prefix:      prefix,
	}
}

func (hs *HttpStore) Get(ctx context.Context, key string) ([]byte, error) {
	key2, err := hs.removePrefix(key)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(hs.endpoint + key2)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()
	if resp.StatusCode == http.StatusNotFound {
		return nil, stores.ErrNotFound
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got non-200 status: %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (hs *HttpStore) Check(ctx context.Context, key string) (bool, error) {
	// TODO: http HEAD
	return false, nil
}

func (hs *HttpStore) Post(ctx context.Context, prefix string, data []byte) (string, error) {
	if len(data) > hs.maxBlobSize {
		return "", stores.ErrMaxSizeExceeded
	}
	_, err := hs.removePrefix(prefix)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(data)
	resp, err := http.Post(hs.endpoint, `application/data`, buf)
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println(err)
		}
	}()

	idBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	idStr := base64.URLEncoding.EncodeToString(idBytes)
	key := hs.prefix + idStr
	return key, nil
}

func (hs *HttpStore) MaxBlobSize() int {
	return hs.maxBlobSize
}

func (hs *HttpStore) removePrefix(x string) (string, error) {
	if !strings.HasPrefix(x, hs.prefix) {
		return "", errors.New("Wrong key: " + x)
	}
	y := x[len(hs.prefix):]
	return y, nil
}