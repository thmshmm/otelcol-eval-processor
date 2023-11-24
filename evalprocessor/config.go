package evalprocessor

type Config struct {
	URL            string `mapstructure:"url"`
	TimeoutSeconds int32  `mapstructure:"timeout_seconds"`
}
