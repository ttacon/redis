package redis

func (c *Client) PFAdd(key, element string, elements ...string) (bool, error) {
	// TODO(ttacon): do it
	return false, nil
}

func (c *Client) PFCount(key string, keys ...string) (int, error) {
	// TODO(ttacon): do it
	return 0, nil
}

// TODO(ttacon): is there a point in having any thing other than error here for
// connection issues as the command documentation says it always returns "OK"
func (c *Client) PFMerge(destination, source string, sources ...string) (bool, error) {
	// TODO(ttacon): do it
	return false, nil
}
