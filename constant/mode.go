//go:build linux
// +build linux

package constant

const (
	Perm0755 = 0755 // 用户具有读/写/执行权限，组用户和其它用户具有读写权限；
	Perm0644 = 0644 // 用户具有读写权限，组用户和其它用户具有只读权限；
)
