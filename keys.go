package redis

// Key commands http://redis.io/commands#generic

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

func (c *Client) PExpire(key string, milliseconds int64) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) PExpireAt(key string, timestamp int64) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) PTTL(key string) (int64, error) {
	// TODO(ttacon): ✔
	return 0, nil
}

func (c *Client) RandomKey() (string, error) {
	resp, err := c.exec("RANDOMKEY")
	if err != nil {
		return "", err
	}
	return c.bulkString(resp)
}

func (c *Client) Rename(key, newkey string) error {
	// TODO(ttacon): ✔
	return nil
}

// TODO(ttacon): add correct function signature
func (c *Client) Restore(key, newkey string) error {
	// TODO(ttacon): ✔
	return nil
}

// TODO(ttacon): add correct function signature
func (c *Client) Sort(key, newkey string) error {
	// TODO(ttacon): ✔
	return nil
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

// TODO(ttacon): add correct function signature
func (c *Client) Scan(key, newkey string) error {
	// TODO(ttacon): ✔
	return nil
}
