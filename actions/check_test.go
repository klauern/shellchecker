package actions_test

func (as *ActionSuite) Test_CheckShellScripts() {

	var checkScriptTests = []struct {
		inFile           string
		expectStatusCode int
		expectErrorCount int
	}{
		{
			inFile: `#!/bin/sh
			`,
			expectStatusCode: 200,
			expectErrorCount: 3,
		},
	}
	for _, tt := range checkScriptTests {
		resp := as.JSON("/check").Post([]byte(tt.inFile))
		as.Equal(tt.expectStatusCode, resp.Code)
	}

}
