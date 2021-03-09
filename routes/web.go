package routes

import (
	. "github.com/confetti-framework/foundation/http/routing"
	"src/app/http/controllers"
	"src/app/http/middleware"
)

// Web contains all route configurations for all your web pages. Here is where
// you can register Web routes for your website. By default these routes are
// loaded within a group which is assigned the "Web" middleware. Enjoy building
// your website!
var Web = Group(
	Get("/", controllers.Homepage),
).Middleware(middleware.Web...)
