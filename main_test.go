package main

import (
	"testing"
)

type Data struct {
	name   string
	region int
	gender int
}

func TestMagic(t *testing.T) {
	dataset := []Data{
		{"尼格买提热合曼", 2, 1},
		{"张三", 1, 1},
		{"张三疯", 1, 1},
		{"李四", 2, 1},
		{"李四小次郎", 2, 1},
		{"玫瑰", 1, 2},
		{"红玫瑰", 1, 2},
		{"蔷薇", 2, 2},
		{"蔷薇漂亮", 2, 2},
	}
	for _, data := range dataset {
		if !magic(data.name, data.region, data.gender) {
			t.Error(data)
		}
	}
}
