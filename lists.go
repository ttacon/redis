package redis

import (
	"errors"
	"fmt"
	"strconv"
)

func (c *Client) BLPop(key string, timeout int, keys ...string) ([]string, error) {
	var vals = []string{key}
	vals = append(vals, keys...)
	vals = append(vals, strconv.Itoa(timeout))
	resp, err := c.exec("BLPOP", vals...)
	if err != nil {
		return nil, err
	}
	return c.stringSlice(resp)
}

func (c *Client) BRPop(key string, timeout int, keys ...string) ([]string, error) {
	var vals = []string{key}
	vals = append(vals, keys...)
	vals = append(vals, strconv.Itoa(timeout))
	resp, err := c.exec("BRPOP", vals...)
	if err != nil {
		return nil, err
	}
	return c.stringSlice(resp)
}

func (c *Client) BRPopLPush(source, destination string, timeout int) (*string, error) {
	resp, err := c.exec("BRPOPLPUSH", source, destination, strconv.Itoa(timeout))
	if err != nil {
		return nil, err
	}
	return c.nillableBulkString(resp)
}

func (c *Client) LIndex(key string, index int) (*string, error) {
	resp, err := c.exec("LINDEX", key, strconv.Itoa(index))
	if err != nil {
		return nil, err
	}
	return c.nillableBulkString(resp)
}

type BeforeAfter string

const (
	BEFORE BeforeAfter = "BEFORE"
	AFTER  BeforeAfter = "AFTER"
)

func (b BeforeAfter) String() string { return string(b) }

// TODO(ttacon): value is interface so can use int, float or string
func (c *Client) LInsert(
	key string,
	insert BeforeAfter,
	pivot string,
	value interface{}) (int64, error) {

	var strVal string
	if s, ok := value.(string); ok {
		strVal = s
	} else if i, ok := value.(int); ok {
		strVal = strconv.Itoa(i)
	} else if f, ok := value.(float64); ok {
		strVal = strconv.FormatFloat(f, byte('f'), -1, 64)
	} else {
		return -1, errors.New(
			fmt.Sprintf("value must be a string, int or float64, got: %v", value))
	}

	resp, err := c.exec("LINSERT", key, insert.String(), pivot, strVal)
	if err != nil {
		return -1, err
	}
	return c.intResp(resp)
}

func (c *Client) LLen(key string) (int64, error) {
	resp, err := c.exec("LLEN", key)
	if err != nil {
		return -1, err
	}
	return c.intResp(resp)
}

func (c *Client) LPop(key string) (*string, error) {
	resp, err := c.exec("LPOP", key)
	if err != nil {
		return nil, err
	}
	return c.nillableBulkString(resp)
}

func (c *Client) LPush(key, value string, values ...string) (int64, error) {
	resp, err := c.exec("LPUSH", append([]string{key, value}, values...)...)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

func (c *Client) LPushx(key, value string) (int64, error) {
	resp, err := c.exec("LPUSHX", key, value)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

func (c *Client) LRange(key string, start, stop int) ([]string, error) {
	resp, err := c.exec("LRANGE", key, strconv.Itoa(start), strconv.Itoa(stop))
	if err != nil {
		return nil, err
	}
	return c.stringSlice(resp)
}

func (c *Client) LRem(key string, count int, value string) (int64, error) {
	resp, err := c.exec("LREM", key, strconv.Itoa(count), value)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

func (c *Client) LSet(key string, index int, value string) (bool, error) {
	resp, err := c.exec("LSET", key, strconv.Itoa(index), value)
	if err != nil {
		return false, err
	}
	val, err := c.stringResp(resp)
	return val == "OK", err
}

func (c *Client) LTrim(key string, start, stop int) (bool, error) {
	resp, err := c.exec("LTRIM", key, strconv.Itoa(start), strconv.Itoa(stop))
	if err != nil {
		return false, err
	}
	val, err := c.stringResp(resp)
	return val == "OK", err
}

func (c *Client) RPop(key string) (*string, error) {
	resp, err := c.exec("RPOP", key)
	if err != nil {
		return nil, err
	}
	return c.nillableBulkString(resp)
}

func (c *Client) RPopLPush(source, destination string) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) RPush(key, value string, values ...string) (int, error) {
	// TODO(ttacon): ✔
	return 0, nil
}

func (c *Client) RPushX(key, value string, values ...string) (int, error) {
	// TODO(ttacon): ✔
	return 0, nil
}
