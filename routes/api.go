package routes

import (
	"github.com/lanvard/contract/inter"
	. "github.com/lanvard/routing"
	"lanvard/src/controllers"
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
var Api = []inter.RouteCollection{
	Get("/users", controllers.User.Index),
	Get("/users/{user_ids?}", controllers.User.Index),
	Get("/user/{user_id}", controllers.User.Show),
	Post("/users", controllers.User.Store),
	Delete("/users/{user_ids}", controllers.User.Destroy),
	// Match([]string{"GET", "POST"}, "/", controllers.User.Store),
}
