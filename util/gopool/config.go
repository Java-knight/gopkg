package gopool

const (
	DEFAULT_SCALA_THRESHOLD = 1 // 默认 Scala 阈值
)

// Config pool 的配置
type Config struct {

	// 如果 len(task chan) > ScaleThreshold，则创建新的 goroutine。默认为 DEFAULT_SCALA_THRESHOLD
	ScaleThreshold int32
}

// NewConfig 创建默认的 Config
func NewConfig() *Config {
	c := &Config{
		ScaleThreshold: DEFAULT_SCALA_THRESHOLD,
	}
	return c
}
