package config

import (
	"github.com/confetti-framework/framework/inter"
	"github.com/confetti-framework/framework/support/env"
)

var Database = struct {
	Default     string
	Connections map[string]inter.Connection
}{

	/*
	   |--------------------------------------------------------------------------
	   | Default Database Connection Name
	   |--------------------------------------------------------------------------
	   |
	   | Here you may specify which of the database connections below you wish
	   | to use as your default connection for all database work. Of course
	   | you may use many connections at once.
	   |
	*/
	Default: env.StringOr("DB_CONNECTION", "mysql"),

	/*
	   |--------------------------------------------------------------------------
	   | Database Connections
	   |--------------------------------------------------------------------------
	   |
	   | Here are each of the database connections setup for your application.
	   | Of course, examples of configuring each database platform that is
	   | supported by Confetti is shown below to make development simple.
	   |
	*/

	Connections: map[string]inter.Connection{
		/*
			"mysql": &db.MySQL{
				Host:         env.StringOr("DB_HOST", "127.0.0.1"),
				Port:         env.IntOr("DB_PORT", 3306),
				Database:     env.StringOr("DB_DATABASE", "confetti"),
				Username:     env.StringOr("DB_USERNAME", "app"),
				Password:     env.StringOr("DB_PASSWORD", ""),
				QueryTimeout: 10 * time.Second,
			},
		*/

		/*
			"postgresql": &db.PostgreSQL{
				Host:         env.StringOr("DB_HOST", "127.0.0.1"),
				Port:         env.IntOr("DB_PORT", 5432),
				Database:     env.StringOr("DB_DATABASE", "confetti"),
				Username:     env.StringOr("DB_USERNAME", "app"),
				Password:     env.StringOr("DB_PASSWORD", ""),
				QueryTimeout: 10 * time.Second,
			},
		*/
	},
}
