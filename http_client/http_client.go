package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/william22913/common/metrics"
)

func NewAPIConnector(
	metrics metrics.Metrics,
) APIConnector {
	return &apiConnector{
		metrics: metrics,
	}
}

type apiConnector struct {
	metrics metrics.Metrics
}

func (a apiConnector) HitAPI(
	method string,
	host string,
	path string,
	header map[string]string,
	body interface{},
	result interface{},
) (
	int,
	error,
) {
	var status int
	now := time.Now()

	defer func() {
		a.metrics.GetDefaultMetric().APIConnectorHist.WithLabelValues(
			host,
			path,
			method,
			strconv.Itoa(status),
		).Observe(float64(time.Since(now).Seconds()))
	}()
	url := fmt.Sprintf("%s%s", host, path)

	bodyByte, _ := json.Marshal(body)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyByte))
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	status = resp.StatusCode
	defer resp.Body.Close()
	bodyResponse, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyResponse, &result)
	if err != nil {
		return 0, err
	}

	return resp.StatusCode, nil
}
