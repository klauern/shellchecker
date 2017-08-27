package actions

import (
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

// ErrorCodeRegexp represents the error code from ShellCheck.  The format is generally
// SC1111, or something similar.
const ErrorCodeRegexp = `(?m)^[Ss][Cc][0-9]{4}$`

// ShellCheckLoc represents the submodule that we load all of the various SCXXXX codes from.
const ShellCheckLoc = "assets/shellcheck.wiki"

// LookupShellCheckError will search for the particular error code returned
// by shellcheck and give the Wiki page from the shellcheck.wiki site.
func LookupShellCheckError(code string) (string, error) {
	log := buffalo.NewLogger("DEBUG")
	log.Debug("lookup shellcheck code for "+code)
	normCode := normalizeCode(code)
	shellCheckCodeFile := path.Join(ShellCheckLoc, normCode+".md")
	_, err := os.Stat(shellCheckCodeFile)
	if os.IsNotExist(err) {
		log.Debugf("%v does not exist", shellCheckCodeFile)
		return "", errors.WithMessage(err, "file "+shellCheckCodeFile+" does not exist")
	}
	return shellCheckCodeFile, nil
}

func normalizeCode(code string) string {
	log := buffalo.NewLogger("DEBUG")
	log.Debug("parsing regex code for "+code)
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
	file, err := LookupShellCheckError(c.Param("code"))
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Render(200, r.HTML(file))
}
