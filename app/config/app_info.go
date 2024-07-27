package config

var AppInfo = struct {
    ApiByPathPrefix,
    Repository,
    Service string
}{
    ApiByPathPrefix: EnvString("API_BY_PATH_PREFIX"),
    Repository:      EnvString("PROJECT_REPOSITORY_NAME"),
    Service:         EnvString("APP_SERVICE"),
}
