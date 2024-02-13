package main

import (
	"testing"
)

type Data struct {
	card   string
	name   string
	region int
	gender int
	debug  bool
}

func TestMagic(t *testing.T) {
	dataset := []Data{
		{"1,2,3,4", "尼格买提热合曼", 2, 1, false},
		{"1,2,3,4", "张三", 1, 1, false},
		{"1,2,3,4", "张三疯", 1, 1, false},
		{"1,2,3,4", "李四", 2, 1, false},
		{"1,2,3,4", "李四小次郎", 2, 1, false},
		{"1,2,3,4", "玫瑰", 1, 2, false},
		{"1,2,3,4", "红玫瑰", 1, 2, false},
		{"1,2,3,4", "蔷薇", 2, 2, false},
		{"1,2,3,4", "蔷薇漂亮", 2, 2, false},
		{"5,6,7,8", "罗兰", 3, 1, false},
		{"5,6,7,8", "紫罗兰", 3, 2, false},
	}
	for _, data := range dataset {
		if !magic(data.card, data.name, data.region, data.gender, data.debug) {
			t.Error(data)
		}
	}
}
