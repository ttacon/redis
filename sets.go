package redis

// TODO(ttacon): make not int64 but int
func (c *Client) SAdd(key, member string, members ...string) (int64, error) {
	args := append([]string{key, member}, members...)
	resp, err := c.exec("SADD", args...)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

// TODO(ttacon): make not int64 but int
func (c *Client) SCard(key string) (int64, error) {
	// TODO(ttacon): do it
	resp, err := c.exec("SCARD", key)
	if err != nil {
		return 0, err
	}
	return c.intResp(resp)
}

func (c *Client) SDiff(key string, keys ...string) ([]string, error) {
	// TODO(ttacon): do it
	return nil, nil
}

func (c *Client) SDiffStore(destination, key string, keys ...string) (int, error) {
	// TODO(ttacon): do it
	return 0, nil
}

func (c *Client) SInter(key string, keys ...string) ([]string, error) {
	// TODO(ttacon): do it
	return nil, nil
}

func (c *Client) SInterStore(destination, key string, keys ...string) (int, error) {
	// TODO(ttacon): do it
	return 0, nil
}

func (c *Client) SIsMember(key, member string) (bool, error) {
	// TODO(ttacon): do it
	return false, nil
}

func (c *Client) SMembers(key string) ([]string, error) {
	resp, err := c.exec("SMEMBERS", key)
	if err != nil {
		return nil, err
	}
	return c.stringSlice(resp)
}

func (c *Client) SMove(source, destination, member string) (bool, error) {
	// TODO(ttacon): do it
	return false, nil
}

func (c *Client) SPop(key string) (string, error) {
	// TODO(ttacon): do it
	return "", nil
}

func (c *Client) SRandMember(key string, count int) ([]string, error) {
	// TODO(ttacon): do it
	return nil, nil
}

func (c *Client) SRem(key, member string, members ...string) (int, error) {
	// TODO(ttacon): do it
	return 0, nil
}

func (c *Client) SUnion(key string, keys ...string) ([]string, error) {
	// TODO(ttacon): do it
	return nil, nil
}

func (c *Client) SUnionStore(destination, key string, keys ...string) (int, error) {
	// TODO(ttacon): do it
	return 0, nil
}

// TODO(ttacon): amend to have correct function signature
func (c *Client) SScan(key string) ([]string, error) {
	// TODO(ttacon): do it
	return nil, nil
}
