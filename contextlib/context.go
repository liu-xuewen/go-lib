package contextlib

import "context"

// GetString get string or default in ctx
func GetString(ctx context.Context, key string, defaultVal string) string {
	v := ctx.Value(key)
	if v != nil {
		return v.(string)
	}

	return defaultVal
}

// GetInt get GetInt or default in ctx
func GetInt(ctx context.Context, key string, defaultVal int) int {
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
func GetUint(ctx context.Context, key string, defaultVal int) int {
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

// GetUint64 get uint64 or default in ctx
func GetUint64(ctx context.Context, key string, defaultVal uint64) uint64 {
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
