package go_tricks

func In(val interface{}, args ...interface{}) (ok bool) {
	switch t := val.(type) {
	case string:
		for _, v := range args {
			ok = t == v
			if ok {
				return
			}
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		for _, v := range args {
			ok = t == v
			if ok {
				return
			}
		}
	case float32, float64:
		for _, v := range args {
			ok = t == v
			if ok {
				return
			}
		}
	case error:
		for _, v := range args {
			ok = errors.Is(t, v.(error))
			if ok {
				return
			}
		}
	case bool:
		for _, v := range args {
			ok = t == v
			if ok {
				return
			}
		}
	}
	return
}
