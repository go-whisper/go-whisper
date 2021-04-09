package model

type Option map[string]interface{}

func (opt Option) Set(k string, v interface{}) {
	opt[k] = v
}

func (opt Option) MustGet(k string) interface{} {
	if v, has := opt[k]; has {
		return v
	}
	return nil
}

func (opt Option) GetString(k string) string {
	if v, ok := opt.MustGet(k).(string); ok {
		return v
	}
	return ""
}

func (opt Option) GetInt(k string) int {
	if v, ok := opt.MustGet(k).(int); ok {
		return v
	}
	return 0
}

func (opt Option) GetUint(k string) uint {
	if v, ok := opt.MustGet(k).(uint); ok {
		return v
	}
	return 0
}
