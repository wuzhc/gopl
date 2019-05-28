// 演示了sync.Once的使用,实际代码运行不了的
package main

import (
	"image"
	"fmt"
)

var icons map[string]image.Image

// 并发不安全
func loadIcons() {
	icons=map[string]image.Image{
		"one.png": loadIcon("one.png"),
		"two.png": loadIcon("two.png"),
		"three.png": loadIcon("three.png")
	}
}

func loadIcon(name string) image.Image{
	return nil
}

func icon(name string) image.Image {
	// icons不为nil,不代表已经全部初始化
	if icons==nil{
		icons=loadIcons()
	}
	return icons[name]
}


// 解决方法如下:
var loadIconsOnce sync.Once
var icon map[string]image.Image

func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIconsOnce)
	return icon[name]
}