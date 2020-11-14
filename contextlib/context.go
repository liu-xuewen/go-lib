package contextlib

import "context"

// CxtKey cxt Key
type CxtKey string

// GetString get string or default in ctx
func GetString(ctx context.Context, key CxtKey, defaultVal string) string {
	v := ctx.Value(key)
	if v != nil {
		return v.(string)
	}

	return defaultVal
}

// GetInt get GetInt or default in ctx
func GetInt(ctx context.Context, key CxtKey, defaultVal int) int {
	v := ctx.Value(key)
	if v != nil {
		value, ok := v.(int)
		if ok {
			return value
		}
		return defaultVal
	}

	return defaultVal
}

// GetUint get GetUint or default in ctx
func GetUint(ctx context.Context, key CxtKey, defaultVal uint) uint {
	v := ctx.Value(key)
	if v != nil {
		value, ok := v.(uint)
		if ok {
			return value
		}
		return defaultVal
	}

	return defaultVal
}

// GetUint64 get uint64 or default in ctx
func GetUint64(ctx context.Context, key CxtKey, defaultVal uint64) uint64 {
	v := ctx.Value(key)
	if v != nil {
		value, ok := v.(uint64)
		if ok {
			return value
		}
		return defaultVal
	}

	return defaultVal
}
