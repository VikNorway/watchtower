package meta

var (
	// Version is the compile-time set version of Watchtower
	Version = "v1.7.1.1vn"

	// UserAgent is the http client identifier derived from Version
	UserAgent string
)

func init() {
	UserAgent = "Watchtower/" + Version
}
