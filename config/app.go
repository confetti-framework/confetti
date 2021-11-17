package config

import (
	"github.com/confetti-framework/support/env"
	"golang.org/x/text/language"
	"os"
	"time"
)

// App contains all the configuration for your application.
var App = struct {
	Name,
	Url,
	AssetUrl,
	LineSeparator,
	Key,
	Env string
	OsArgs   []string
	Port     int
	Host     string
	Cipher   string
	Debug    bool
	Timezone *time.Location
	Locale,
	FallbackLocale,
	FakerLocale language.Tag
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
	Name: env.StringOr("APP_NAME", "Confetti"),

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
	Url: env.StringOr("APP_URL", "http://localhost"),

	/*
		|--------------------------------------------------------------------------
		| Asset URL
		|--------------------------------------------------------------------------
		|
		| You can configure the asset URL host, this can be useful if you host your
		| assets on an external service like Amazon S3.
		|
	*/
	AssetUrl: env.StringOr("ASSET_URL", "http://asset.localhost"),

	/*
		|--------------------------------------------------------------------------
		| Line separator
		|--------------------------------------------------------------------------
		|
		| Determine what the line separator should be for your application
		|
	*/
	LineSeparator: env.StringOr("LINE_SEPARATOR", "\n"),

	/*
		|--------------------------------------------------------------------------
		| Encryption Key
		|--------------------------------------------------------------------------
		|
		| This key is used by the Confetti encrypter service and should be set
		| to a random, 32 character string, otherwise these encrypted strings
		| will not be safe. Please do this before deploying an application!
		|
	*/
	Key: env.String("APP_KEY"),

	/*
		|--------------------------------------------------------------------------
		| Cipher type
		|--------------------------------------------------------------------------
		|
		| Source encrypted values are encrypted using OpenSSL and this
		| cipher. Furthermore, all encrypted values are signed with a message
		| authentication code (MAC) to detect any modifications to the encrypted
		| string.
		|
	*/
	Cipher: "AES-256-CBC",

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
	Env: env.StringOr("APP_ENV", "production"),

	/*
	   |--------------------------------------------------------------------------
	   | Application Port
	   |--------------------------------------------------------------------------
	   |
	   | Determines which port the application should listen to. For a development
	   | environment this will be 80 and for an environment with a certificate
	   | this will be 443. If you have a proxy server (such as Docker, NGINX or a
	   | load balancer), you can forward the proxy server to a port number
	   | specified here.
	   |
	*/
	Port: env.IntOr("APP_PORT", 8080),

	/*
	   |--------------------------------------------------------------------------
	   | Application Host
	   |--------------------------------------------------------------------------
	   |
	   | This specifies the TCP address for the server to listen on. The service
	   | names are defined in RFC 6335 and assigned by IANA. See net.Dial for
	   | details of the address format (without :port). To match routers with a
	   | domain name, see the routing documentation.
	   |
	*/
	Host: env.StringOr("APP_HOST", ""),

	/*
	   |--------------------------------------------------------------------------
	   | Application Command-line Arguments
	   |--------------------------------------------------------------------------
	   |
	   | Args hold the command-line arguments, starting with the program name.
	   |
	*/
	OsArgs: os.Args,

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
	Debug: env.BoolOr("APP_DEBUG", false),

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
	Locale: env.LangOr("LOCALE", language.AmericanEnglish),

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
	FallbackLocale: env.LangOr("FALLBACK_LOCALE", language.AmericanEnglish),

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
	FakerLocale: env.LangOr("FAKE_LOCALE", language.AmericanEnglish),
}
