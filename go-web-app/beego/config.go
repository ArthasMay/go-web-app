package beego

var (
	BConfig *Config
)

type Config struct {
	AppName string
	WebConfig WebConfig
}

type WebConfig struct {
	AutoRender bool
	StaticDir  map[string]string
}

func init() {
	BConfig = newConfig()
}

func newConfig() *Config {
	return &Config {
		AppName: "beego",
		WebConfig: WebConfig {
			StaticDir: map[string]string{"/static": "static"},
			AutoRender: true,
		},
	}
}