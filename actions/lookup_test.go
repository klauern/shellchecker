package actions_test

func (as *ActionSuite) Test_LookupCode_GET() {
	var tt = []struct {
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
	for _, tc := range tt {
		res := as.HTML("/code/" + tc.scCode).Get()
		as.Equal(tc.code, res.Code)
	}
}
