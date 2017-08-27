package actions

import (
	"regexp"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// ErrorCodeRegexp represents the error code from ShellCheck.  The format is generally
// SC1111, or something similar.
const ErrorCodeRegexp = `(?m)^[Ss][Cc][0-9]{4}$`

// ShellCheckLoc represents the submodule that we load all of the various SCXXXX codes from.
const ShellCheckLoc = "shellcheck.wiki"

func normalizeCode(code string) string {
	log := buffalo.NewLogger("DEBUG")
	log.Debug("parsing regex code for " + code)
	var re = regexp.MustCompile(ErrorCodeRegexp)
	if len(re.FindString(code)) > 0 {
		log.Debug("found shellcheck code in regex")
		norm := strings.ToUpper(code)
		norm = strings.Trim(norm, " ")
		return norm
	}
	return code
}

// LookupShellCheckErrorHandler is a handler for /code/{code} lookups.
func LookupShellCheckErrorHandler(c buffalo.Context) error {
	code := normalizeCode(c.Param("code"))
	debugLog.Debugf("lookup code for " + code)
	p := ShellCheckLoc + "/" + code + ".md"
	if r.TemplatesBox.Has(p) {
		debugLog.Debugf("TemplatesBox has code %v", p)
		return c.Render(200, r.HTML(p))
	} else {
		debugLog.Debugf("TemplatesBox can't find %v", p)
	}
	return c.Error(404, errors.Errorf("could not find %s", c.Param("code")))

}
