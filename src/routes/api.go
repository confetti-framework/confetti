package routes

import (
	"lanvard/foundation/http/controllers"
	"lanvard/routing"
)

/*
|--------------------------------------------------------------------------
| API MapMethodRoutes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/
var Api = []routing.RouteCollection{
	routing.Get("/users", controllers.User.Index),
	routing.Get("/users/{ids}", controllers.User.Show),
	routing.Post("/users", controllers.User.Store),
	routing.Delete("/users/{ids?}", controllers.User.Destroy),
}
