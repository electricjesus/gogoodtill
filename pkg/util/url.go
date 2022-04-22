package util

import "strings"

type URL struct {
	Base string // Base URL no trailing slash
}

func (u *URL) Get(params ...string) string {
	parts := append([]string{u.Base}, params...)
	return strings.Join(parts, "/")
}
