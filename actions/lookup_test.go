package actions_test

func (as *ActionSuite) Test_LookupCode_GET() {
	as.HTML("/code/SC1024").Get()
}

