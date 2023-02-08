package cache

import (
	"errors"
)

type Data struct {
	data map[string]interface{}
}

func New() *Data {
	return &Data{
		data: make(map[string]interface{}),
	}
}
func (d *Data) Set(key string, value interface{}) {
	d.data[key] = value
}

func (d Data) Get(key string) (interface{}, error) {
	value, ok := d.data[key]
	if !ok {
		return nil, errors.New("value didn't found")
	}
	return value, nil
}

func (d Data) Delete(key string) {
	delete(d.data, key)
}
