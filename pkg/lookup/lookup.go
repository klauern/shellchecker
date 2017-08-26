package lookup

import (
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

// ErrorCodeRegexp represents the error code from ShellCheck.  The format is generally
// SC1111, or something similar.
const ErrorCodeRegexp = `(?m)^[Ss][Cc][0-9]{4}$`

// ShellCheckLoc represents the submodule that we load all of the various SCXXXX codes from.
const ShellCheckLoc = "shellcheck.wiki"

// LookupShellCheckError will search for the particular error code returned
// by shellcheck and give the Wiki page from the shellcheck.wiki site.
func LookupShellCheckError(code string) ([]byte, error) {
	normCode := normalizeCode(code)
	shellCheckCodeFile := path.Join(ShellCheckLoc, normCode+".md")
	_, err := os.Stat(shellCheckCodeFile)
	if os.IsNotExist(err) {
		return nil, errors.WithMessage(err, "file "+normCode+".md does not exist")
	}
	return ioutil.ReadFile(shellCheckCodeFile)
}

func normalizeCode(code string) string {
	var re = regexp.MustCompile(ErrorCodeRegexp)
	if len(re.FindString(code)) > 0 {
		norm := strings.ToUpper(code)
		norm = strings.Trim(norm, " ")
		return norm
	}
	return code
}
