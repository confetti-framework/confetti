package config

import (
    "path"
)

var AppInfo = struct {
    ApiByPathPrefix,
    Service string
}{
    ApiByPathPrefix: EnvString("API_BY_PATH_PREFIX"),
    Service:         "/" + path.Join(EnvString("PROJECT_REPOSITORY_NAME"), EnvString("APP_SERVICE")),
}
