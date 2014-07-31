package redis

func (c *Client) Hdel(key, field string, fields ...string) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) Hexists(key, field string) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) Hget(key, field string) (string, error) {
	// TODO(ttacon): ✔
	return "", nil
}

func (c *Client) Hgetall(key string) ([]string, error) {
	// TODO(ttacon): ✔
	return nil, nil
}

func (c *Client) Hincrby(key, field string, increment int) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) Hincrbyfloat(key, field string, increment float) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) Hkeys(key string) ([]string, error) {
	// TODO(ttacon): ✔
	return nil, nil
}

func (c *Client) Hlen(key string) (int, error) {
	// TODO(ttacon): ✔
	return 0, nil
}

func (c *Client) Hmget(key, field string, fields ...string) ([]string, error) {
	// TODO(ttacon): ✔
	return nil, nil
}

func (c *Client) Hmset(key string, values map[string]interface{}) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) Hset(key, field string, value interface{}) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) HsetNx(key, field string, value interface{}) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) Hvals(key string) ([]string, error) {
	// TODO(ttacon): ✔
	return nil, nil
}

// TODO(ttacon): set correct function signature
func (c *Client) Hscan(key string) error {
	// TODO(ttacon): ✔
	return nil
}
