redis
=====

A redis client in pure go that doesn't return []byte.

Most other redis client libraries in go are prety good, but most
return a []byte which I feel it a little painful. There was one that
returned strings instead, but it was a wrapper for the C bindings -
I was looking for something that was just go.