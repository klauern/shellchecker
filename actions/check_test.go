package actions_test

func (as *ActionSuite) Test_CheckShellScripts() {

	var tt = []struct {
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
	for _, tc := range tt {
		resp := as.JSON("/check").Post([]byte(tc.inFile))
		as.Equal(tc.expectStatusCode, resp.Code)
	}

}
