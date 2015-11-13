package key

import "errors"

// The method to make a key with prefix.
//
// key the real key; prefix the prefix to the key.
func MakeKey(key, prefix *string) (*string, error) {
	if key == nil {
		return nil, errors.New("The key can't be nil.")
	}

	if prefix == nil {
		return nil, errors.New("The prefix can't be nil.")
	}

	length := len(*key) + len(*prefix)

	bs := make([]byte, length)
	bl := 0

	bl += copy(bs[bl:], *prefix)
	bl += copy(bs[bl:], *key)

	rKey := string(bs)
	return &rKey, nil
}
