package actions_test

func (as *ActionSuite) Test_CheckShellScripts() {

	var checkScriptTests = []struct {
		inFile      string
		expectError bool
		errorCount  int
	}{}

	for _, tt := range checkScriptTests {
		if tt.expectError {
			as.JSON("/check").Post(tt.inFile)
		}
	}

}
