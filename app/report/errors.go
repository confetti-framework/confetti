package report

import (
	"github.com/confetti-framework/errors"
	"github.com/confetti-framework/syslog/log_level"
	net "net/http"
)

// UserError Below you will find a list of all user errors. By default these errors
// are not logged by config `config.Errors.NoLogging`. These errors are
// usually errors with http status 499 and lower.
var UserError = errors.WithLevel(errors.WithStatus(errors.New(""), net.StatusBadRequest), log_level.INFO)
var ValidationError = errors.WithStatus(UserError, net.StatusUnprocessableEntity)
var NotFoundError = errors.WithStatus(UserError, net.StatusNotFound)
var PageNotFoundError = errors.Wrap(NotFoundError, "page not found")

// SystemError This list contains errors that indicate that the system is not working
// properly. The message is not displayed to the user on a production environment,
// but will be logged (if MinLevel of the logger allows it).
var SystemError = errors.WithLevel(errors.WithStatus(errors.New(""), net.StatusInternalServerError), log_level.EMERGENCY)
