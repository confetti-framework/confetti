package config

var AppInfo = struct {
    ServiceUriPrefix,
    Repository,
    Service string
}{
    ServiceUriPrefix: EnvString("SERVICE_URI_PREFIX"),
    Repository:       EnvString("PROJECT_REPOSITORY_NAME"),
    Service:          EnvString("APP_SERVICE"),
}
