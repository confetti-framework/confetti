package config

import (
    "path"
)

var AppInfo = struct {
    Service string
}{
    Service: "/" + path.Join(EnvString("PROJECT_REPOSITORY_NAME"), EnvString("APP_SERVICE")),
}
