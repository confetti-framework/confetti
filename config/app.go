package config

import (
	. "lanvard/support"
	"path/filepath"
	"runtime"
)

var App = struct {
	Name           string
	Env            string
	Debug          bool
	Url            string
	AssetUrl       string
	BasePath       string
	Timezone       string
	Locale         string
	FallbackLocale string
	FakerLocale    string
	Key            string
	Cipher         string
}{

	/*
	   |--------------------------------------------------------------------------
	   | Application Name
	   |--------------------------------------------------------------------------
	   |
	   | This value is the name of your application. This value is used when the
	   | framework needs to place the application's name in a notification or
	   | any other location as required by the application or its packages.
	   |
	*/
	Name: EnvOr("APP_NAME", "Lanvard"),

	/*
	   |--------------------------------------------------------------------------
	   | Application Environment
	   |--------------------------------------------------------------------------
	   |
	   | This value determines the "environment" your application is currently
	   | running in. This may determine how you prefer to configure various
	   | services the application utilizes. Set this in your ".env" file.
	   |
	*/
	Env: EnvOr("APP_ENV", "production"),

	/*
	   |--------------------------------------------------------------------------
	   | Application Debug Mode
	   |--------------------------------------------------------------------------
	   |
	   | When your application is in debug mode, detailed error messages with
	   | stack traces will be shown on every error that occurs within your
	   | application. If disabled, a simple generic error page is shown.
	   |
	*/
	Debug: EnvOrBool("APP_DEBUG", false),

	/*
	   |--------------------------------------------------------------------------
	   | Application URL
	   |--------------------------------------------------------------------------
	   |
	   | This URL is used by the console to properly generate URLs when using
	   | the Artisan command line tool. You should set this to the root of
	   | your application so that it is used when running Artisan tasks.
	   |
	*/
	Url: EnvOr("APP_URL", "http://localhost"),

	/*
	   |--------------------------------------------------------------------------
	   | Asset URL
	   |--------------------------------------------------------------------------
	   |
	   | You can configure the asset URL host, this can be useful if you host your
	   | assets on an external service like Amazon S3.
	   |
	*/
	AssetUrl: EnvOr("ASSET_URL", "http://asset.localhost"),

	/*
	   |--------------------------------------------------------------------------
	   | Base Path
	   |--------------------------------------------------------------------------
	   |
	   | The base path is the fully qualified path to the project root.
	   |
	*/
	BasePath: getBasePath(),

	/*
	   |--------------------------------------------------------------------------
	   | Application Timezone
	   |--------------------------------------------------------------------------
	   |
	   | Here you may specify the default timezone for your application, which
	   | will be used by the PHP date and date-time functions. We have gone
	   | ahead and set this to a sensible default for you out of the box.
	   |
	*/
	Timezone: "UTC",

	/*
	   |--------------------------------------------------------------------------
	   | Application Locale Configuration
	   |--------------------------------------------------------------------------
	   |
	   | The application locale determines the default locale that will be used
	   | by the translation service provider. You are free to set this value
	   | to any of the locales which will be supported by the application.
	   |
	*/
	Locale: "en",

	/*
	   |--------------------------------------------------------------------------
	   | Application Fallback Locale
	   |--------------------------------------------------------------------------
	   |
	   | The fallback locale determines the locale to use when the current one
	   | is not available. You may change the value to correspond to any of
	   | the language folders that are provided through your application.
	   |
	*/
	FallbackLocale: "en",

	/*
	   |--------------------------------------------------------------------------
	   | Faker Locale
	   |--------------------------------------------------------------------------
	   |
	   | This locale will be used by the Faker PHP library when generating fake
	   | data for your database seeds. For example, this will be used to get
	   | localized telephone numbers, street address information and more.
	   |
	*/
	FakerLocale: "en_US",

	/*
	   |--------------------------------------------------------------------------
	   | Encryption Key
	   |--------------------------------------------------------------------------
	   |
	   | This key is used by the Illuminate encrypter service and should be set
	   | to a random, 32 character string, otherwise these encrypted strings
	   | will not be safe. Please do this before deploying an application!
	   |
	*/
	Key: Env("APP_KEY"),

	/*
	   |--------------------------------------------------------------------------
	   | Cipher type
	   |--------------------------------------------------------------------------
	   |
	   | All encrypted values are encrypted using OpenSSL and this
	   | cipher. Furthermore, all encrypted values are signed with a message
	   | authentication code (MAC) to detect any modifications to the encrypted
	   | string.
	   |
	*/
	Cipher: "AES-256-CBC",
}

func getBasePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filepath.Dir(filename))
}
