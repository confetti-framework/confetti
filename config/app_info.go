package config

var AppInfo = struct {
	Permission,
	Repository,
	Service,
	ServiceUriPrefix string
}{
	Permission:       EnvString("APP_PERMISSION"),
	Repository:       EnvString("PROJECT_REPOSITORY_NAME"),
	Service:          EnvString("APP_SERVICE"),
	ServiceUriPrefix: EnvString("SERVICE_URI_PREFIX"),
}
