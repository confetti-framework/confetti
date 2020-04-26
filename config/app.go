package config

import (
	"github.com/lanvard/contract/inter"
	"github.com/lanvard/support/environment"
	"golang.org/x/text/language"
	"time"
)

var _ = environment.Bootstrap(NewPath())

var App = struct {
	Name           string
	Env            string
	Debug          bool
	Url            string
	AssetUrl       string
	LineSeparator  string
	BasePath       inter.BasePath
	Timezone       *time.Location
	Locale         language.Tag
	FallbackLocale language.Tag
	FakerLocale    language.Tag
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
	Name: environment.StrEnvOr("APP_NAME", "Lanvard"),

	/*
	   |--------------------------------------------------------------------------
	   | Application IsEnvironment
	   |--------------------------------------------------------------------------
	   |
	   | This value determines the "environment" your application is currently
	   | running in. This may determine how you prefer to configure various
	   | services the application utilizes. Set this in your ".env" file.
	   |
	*/
	Env: environment.StrEnvOr("APP_ENV", "production"),

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
	Debug: environment.BoolEnvOr("APP_DEBUG", false),

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
	Url: environment.StrEnvOr("APP_URL", "http://localhost"),

	/*
	   |--------------------------------------------------------------------------
	   | Asset URL
	   |--------------------------------------------------------------------------
	   |
	   | You can configure the asset URL host, this can be useful if you host your
	   | assets on an external service like Amazon S3.
	   |
	*/
	AssetUrl: environment.StrEnvOr("ASSET_URL", "http://asset.localhost"),

	/*
	   |--------------------------------------------------------------------------
	   | Line separator
	   |--------------------------------------------------------------------------
	   |
	   | Determine what the line separator should be for your application
	   |
	*/
	LineSeparator: environment.StrEnvOr("LINE_SEPARATOR", "\n"),

	/*
	   |--------------------------------------------------------------------------
	   | Base AppPath
	   |--------------------------------------------------------------------------
	   |
	   | The base path is the fully qualified path to the project root. Feel free
	   | to adjust this so that it fits to your needs.
	   |
	*/
	BasePath: NewPath(),

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
	Timezone: time.UTC,

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
	Locale: language.AmericanEnglish,

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
	FallbackLocale: language.AmericanEnglish,

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
	FakerLocale: language.AmericanEnglish,

	/*
	   |--------------------------------------------------------------------------
	   | Encryption Key
	   |--------------------------------------------------------------------------
	   |
	   | This key is used by the Lanvard encrypter service and should be set
	   | to a random, 32 character string, otherwise these encrypted strings
	   | will not be safe. Please do this before deploying an application!
	   |
	*/
	Key: environment.StrEnv("APP_KEY"),

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
