package build

const (
	DEV  = "dev"
	PROD = "prod"
	Test = "test"
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

func IsTest() bool {
	return VersionType == Test
}
