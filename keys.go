package redis

import

// Key commands http://redis.io/commands#generic
"strconv"

func (c *Client) Del(key string, otherKeys ...string) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) Dump(key string) (bool, error) {
	// TODO(ttacon): ✔
	return false, nil
}

func (c *Client) Expire(key string, seconds int) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) ExpireAt(key string, timestamp int64) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) Keys(pattern string) ([]string, error) {
	// TODO(ttacon): ✔
	return nil, nil
}

// TODO(ttacon): add correct function signature
func (c *Client) Migrate(key string) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) Move(key string, db int) error {
	// TODO(ttacon): ✔
	return nil
}

// TODO(ttacon): add correct function signature
func (c *Client) Object(key string) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) Persist(key string) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) PExpire(key string, milliseconds int64) (bool, error) {
	resp, err := c.exec("PEXPIRE", key, strconv.FormatInt(milliseconds, 10))
	if err != nil {
		return false, err
	}
	val, err := c.intResp(resp)
	return val == 1, err
}

func (c *Client) PExpireAt(key string, timestamp int64) (bool, error) {
	resp, err := c.exec("PEXPIREAT", key, strconv.FormatInt(timestamp, 10))
	if err != nil {
		return false, err
	}
	v, err := c.intResp(resp)
	return v == 1, err
}

func (c *Client) PTTL(key string) (int64, error) {
	resp, err := c.exec("PTTL", key)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

func (c *Client) RandomKey() (string, error) {
	resp, err := c.exec("RANDOMKEY")
	if err != nil {
		return "", err
	}
	return c.bulkString(resp)
}

func (c *Client) Rename(key, newkey string) error {
	_, err := c.exec("RENAME", key, newkey)
	return err
}

func (c *Client) Restore(key string, ttl int, serialized string) error {
	_, err := c.exec("RESTORE", key, strconv.Itoa(ttl), serialized)
	return err
}

func (c *Client) Sort(key string, modifiers ...string) ([]string, error) {
	resp, err := c.exec("SORT", append([]string{key}, modifiers...)...)
	if err != nil {
		return nil, err
	}
	return c.stringSlice(resp)
}

func (c *Client) TTL(key string) (int64, error) {
	resp, err := c.exec("TTL", key)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

// TODO(ttacon): add type enum and change return types
func (c *Client) Type(key string) (string, error) {
	resp, err := c.exec("TYPE", key)
	if err != nil {
		return "", err
	}
	return c.stringResp(resp)
}

func (c *Client) Scan(key string, cursor int) ([]string, int, error) {
	resp, err := c.exec("SCAN", strconv.Itoa(cursor))
	if err != nil {
		return nil, 0, err
	}
	vals, err := c.stringSlice(resp)

	// the first val in vals (if no error) is the cursor val
	if err != nil {
		return nil, 0, err
	}

	// TODO(ttacon): this should probably be 64, shouldn't it?
	val, err := strconv.ParseInt(vals[0], 10, 32)
	return vals[1:], int(val), err
}
