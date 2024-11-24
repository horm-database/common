package url

import (
	"net/url"
)

func ParseQuery(s string) (map[string]string, error) {
	ret := map[string]string{}

	if s == "" {
		return ret, nil
	}

	tmp, err := url.ParseQuery(s)
	if err != nil {
		return nil, err
	}

	for k, v := range tmp {
		ret[k] = v[0]
	}

	return ret, nil
}

func ParamEncode(params map[string]string) string {
	value := url.Values{}
	for k, v := range params {
		value[k] = []string{v}
	}

	return value.Encode()
}
