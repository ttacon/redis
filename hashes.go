package redis

import (
	"errors"
	"fmt"
	"strconv"
)

func (c *Client) Hdel(key, field string, fields ...string) (int, error) {
	resp, err := c.exec("HDEL", append([]string{key, field}, fields...)...)
	if err != nil {
		return 0, err
	}
	val, err := c.intResp(resp)
	return int(val), err
}

func (c *Client) Hexists(key, field string) (bool, error) {
	resp, err := c.exec("HEXISTS", key, field)
	if err != nil {
		return false, err
	}
	return c.boolResp(resp)
}

func (c *Client) Hget(key, field string) (string, error) {
	resp, err := c.exec("HGET", key, field)
	if err != nil {
		return "", err
	}
	return c.stringResp(resp)
}

func (c *Client) Hgetall(key string) ([]string, error) {
	resp, err := c.exec("HGETALL", key)
	if err != nil {
		return nil, err
	}
	return c.stringSlice(resp)
}

func (c *Client) Hincrby(key, field string, increment int) (int64, error) {
	resp, err := c.exec("HINCRBY", key, field, fmt.Sprintf("%d", increment))
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

func (c *Client) Hincrbyfloat(key, field string, increment float64) (float64, error) {
	resp, err := c.exec(
		"HINCRBYFLOAT",
		key,
		field,
		strconv.FormatFloat(increment, "f", 64))
	if err != nil {
		return 0, err
	}
	val, err := c.stringResp(resp)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(val, 64)
}

func (c *Client) Hkeys(key string) ([]string, error) {
	resp, err := c.exec("HKEYS", key)
	if err != nil {
		return nil, err
	}
	return c.stringSlice(resp)
}

func (c *Client) Hlen(key string) (int64, error) {
	resp, err := c.exec("HLEN", key)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

func (c *Client) Hmget(key, field string, fields ...string) ([]string, error) {
	resp, err := c.exec("HMGET", append([]string{field}, fields...)...)
	if err != nil {
		return nil, err
	}
	return c.stringSlice(resp)
}

func (c *Client) Hmset(key string, values map[string]string) error {
	_, err := c.exec("HMSET", append([]string{key}, mapToStrSlice(values)...)...)
	return err
}

func mapToStrSlice(mp map[string]string) []string {
	var (
		vals = make([]string, len(mp)*2)
		i    = 0
	)

	for k, v := range mp {
		vals[i], vals[i+1] = k, v
		i += 2
	}
	return vals
}

func (c *Client) Hset(key, field string, value interface{}) (int, error) {
	var strVal string
	if str, ok := value.(string); ok {
		strVal = str
	} else if i, ok := value.(int); ok {
		strVal = strconv.FormatInt(int64(i), 64)
	} else {
		return -1, errors.New(
			fmt.Sprintf(
				"HSET only takes int or string as it's value, was given: %v",
				value))
	}

	resp, err := c.exec("HSET", field, strVal)
	if err != nil {
		return -1, err
	}
	return c.intResp(resp)
}

func (c *Client) Hsetnx(key, field string, value interface{}) (int, error) {
	var strVal string
	if str, ok := value.(string); ok {
		strVal = str
	} else if i, ok := value.(int); ok {
		strVal = strconv.FormatInt(int64(i), 64)
	} else {
		return -1, errors.New(
			fmt.Sprintf(
				"HSET only takes int or string as it's value, was given: %v",
				value))
	}

	resp, err := c.exec("HSETNX", field, strVal)
	if err != nil {
		return -1, err
	}
	return c.intResp(resp)
}

func (c *Client) Hvals(key string) ([]string, error) {
	resp, err := c.exec("HVALS", key)
	if err != nil {
		return nil, err
	}
	return c.stringSlice(resp)
}

// TODO(ttacon): set correct function signature
func (c *Client) Hscan(key string, cursor int64) (int64, []string, error) {
	resp, err := c.sxec("HSCAN", key, strconv.FormatInt(cursor, 10))
	if err != nil {
		return -1, nil, err
	}

	vals, err := c.stringSlice(resp)
	if err != nil {
		return -1, nil, err
	}

	// for now we'll assume if there was no error len(vals) >= 0
	val, err := strconv.ParseInt(vals[0], 10, 64)
	return vals[1:], val, err
}
