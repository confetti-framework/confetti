package report

import (
	"github.com/lanvard/errors"
	"github.com/lanvard/syslog/log_level"
	net "net/http"
)

/**
 * Below you will find a list of all user errors. By default these errors are not logged by config `config.Errors.NoLogging`.
 */
var UserError = errors.New("").Status(net.StatusBadRequest).Level(log_level.INFO)
var PageNotFound = UserError.Wrap("page not found").Status(net.StatusNotFound)

var SystemError = errors.New("").Status(net.StatusInternalServerError)
