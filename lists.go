package redis

func (c *Client) BLPop(key string, timeout int, keys ...string) (string, error) {
	// TODO(ttacon): ✔
	return "", nil
}

func (c *Client) BRPop(key string, timeout int, keys ...string) (string, error) {
	// TODO(ttacon): ✔
	return "", nil
}

func (c *Client) BRPopLPush(source, destination string, timeout int) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) LIndex(key string, index int) (string, error) {
	// TODO(ttacon): ✔
	return "", nil
}

type BeforeAfter string

const (
	BEFORE BeforeAfter = "BEFORE"
	AFTER  BeforeAfter = "AFTER"
)

func (b BeforeAfter) String() string { return string(b) }

// TODO(ttacon): value is interface so can use int, float or string
func (c *Client) LInsert(key string, insert BeforeAfter, pivot string, value interface{}) (string, error) {
	// TODO(ttacon): ✔
	return "", nil
}

func (c *Client) LLen(key string) (int, error) {
	// TODO(ttacon): ✔
	return 0, nil
}

func (c *Client) LPop(key string) (string, error) {
	// TODO(ttacon): ✔
	return "", nil
}

func (c *Client) LPush(key, value string, values ...string) (int, error) {
	// TODO(ttacon): ✔
	return 0, nil
}

func (c *Client) LRange(key string, start, stop int) ([]string, error) {
	// TODO(ttacon): ✔
	return nil, nil
}

func (c *Client) LRem(key string, count int, value string) (int, error) {
	// TODO(ttacon): ✔
	return 0, nil
}

func (c *Client) LSet(key string, index int, value string) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) LTrim(key string, start, stop int) error {
	// TODO(ttacon): ✔
	return nil
}

func (c *Client) RPop(key string) (string, error) {
	// TODO(ttacon): ✔
	return "", nil
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
