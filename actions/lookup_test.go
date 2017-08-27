package actions_test

func (as *ActionSuite) Test_LookupCode_GET() {
	var badCodeLookupTests = []struct {
		scCode string
		code   int
	}{
		{"sc1020", 200},
		{"X", 404},
		{"sc1001", 200},
		{"SC", 404},
		{"SC444", 404},
		{"SC2222", 404},
	}
	for _, tt := range badCodeLookupTests {
		res := as.HTML("/code/" + tt.scCode).Get()
		as.Equal(tt.code, res.Code)
	}
}
