package redis

func (c *Client) Eval(script string, numkeys int, keysAndArgs ...string) ([]string, error) {
	// TODO(ttacon): do it
	return nil, nil
}

func (c *Client) EvalSHA(sha1 string, numkeys int, keysAndArgs ...string) ([]string, error) {
	// TODO(ttacon): do it
	return nil, nil
}

func (c *Client) ScriptExists(script string, scripts ...string) ([]bool, error) {
	// TODO(ttacon): do it
	return nil, nil
}

func (c *Client) ScriptFlush() error {
	// TODO(ttacon): do it
	return nil
}

func (c *Client) ScriptKill() error {
	// TODO(ttacon): do it
	return nil
}

func (c *Client) ScriptLoad(script string) (string, error) {
	// TODO(ttacon): do it
	return "", nil
}
