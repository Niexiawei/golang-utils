//go:build !prod && !test

package build

func init() {
	VersionType = DEV
}
