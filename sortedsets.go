package redis

import "strconv"

func (c *Client) ZAdd(key, score, member string, values ...string) (int64, error) {
	resp, err := c.exec("ZADD", append([]string{key, score, member}, values...)...)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

func (c *Client) ZCard(key string) (int64, error) {
	resp, err := c.exec("ZCARD", key)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

func (c *Client) ZCount(key, min, max string) (int64, error) {
	resp, err := c.exec("ZCOUNT", key, min, max)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

func (c *Client) ZIncryBy(key, increment, member string) (string, error) {
	resp, err := c.exec("ZINCRBY", key, increment, member)
	if err != nil {
		return "", err
	}
	return c.bulkString(resp)
}

// TODO(ttacon): put in correct/intuitive function signature
func (c *Client) ZInterStore(key string) (int, error) {
	// TODO(ttacon): do it
	return 0, nil
}

func (c *Client) ZLexCount(key string, min, max int) (int64, error) {
	resp, err := c.exec("ZLexCount", key, strconv.Itoa(min), strconv.Itoa(max))
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}
