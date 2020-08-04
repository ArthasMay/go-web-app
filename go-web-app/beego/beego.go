package beego

const (
	VERSION = "1.0.0"
	DEV = "dev"
	PROD = "prod"
)

func Run(params ...string) {
	BeeApp.Run()
}