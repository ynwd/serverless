package internal

type containter map[string]interface{}

func New() containter {
	var c = make(map[string]interface{})
	c["db"] = func(){}
	return c
}
