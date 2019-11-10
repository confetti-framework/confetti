package support

import (
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"runtime"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(filepath.Dir(filename))

	file := basePath + string(os.PathSeparator) + ".env"
	err := godotenv.Load(file)
	if err != nil {
		println(err)
		panic("Error loading .env file in directory " + file)
	}
}

func Env(search string) string {
	env, OK := os.LookupEnv(search)
	if !OK {
		panic("Enviroment '" + search + "' not found")
	}

	return env
}

func EnvOr(search string, def string) string {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	return env
}

func EnvOrBool(search string, def bool) bool {
	env, OK := os.LookupEnv(search)
	if !OK {
		return def
	}

	return env == "true"
}
