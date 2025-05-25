package out

import "io"

type QuickChartApi interface {
	GetWordCloud(text string) (io.Reader, int64, error)
}
