package okerlund

import (
	"os"
)

//lambdaEnv checks whether code is executing in an AWS Lambda environment
func lambdaEnv() bool {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" { //Is there a better approach than this...?
		return true
	}
	return false
}

func azureFunctionEnv() {}

func gcpFunctionEnv() {}

func kubelessEnv() {}

func serverlessEnv() bool {
	if lambdaEnv() {
		return true
	}
	return false
}
