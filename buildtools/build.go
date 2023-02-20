package build

const (
	DEV  = "dev"
	PROD = "prod"
)

var (
	VersionType string
)

func IsProd() bool {
	return VersionType == PROD
}

func IsDev() bool {
	return VersionType == DEV
}
