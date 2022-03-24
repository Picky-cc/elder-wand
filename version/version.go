package version

import "time"

var (
	VERSION    string
	COMMIT_ID  string
	BUILD_TIME string
	GO_VERSION string
	START_TIME string
)

func init() {
	START_TIME = time.Now().Format("2006-01-02 15:04:05.000000")
}
