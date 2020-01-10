package util

func ErrState(v interface{}) error {
	if es, ok := v.(interface{ ErrState() error }); ok {
		return es.ErrState()
	}
	return nil
}
