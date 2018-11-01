package okerlund

import (
	"os"
)

//IsLambdaEnv checks whether code is executing in an AWS Lambda environment
func IsLambdaEnv() bool {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" { //Is there a better approach than this...?
		return true
	}
	return false
}

// IsAzureFunctionEnv checks whether code is running in an Azure Function environment
func IsAzureFunctionEnv() bool {
	if os.Getenv("WEBSITE_INSTANCE_ID") != "" { //Is there a better approach than this...?
		return true
	}
	return false
}

// IsGcpFunctionEnv checks whether code is running in a Google Cloud Platform Function environment
func IsGcpFunctionEnv() bool {
	if os.Getenv("GCP_PROJECT") != "" { //Is there a better approach than this...?
		return true
	}
	return false
}

// IsKubelessEnv checks whether code is running in a Kubless function environment
func IsKubelessEnv() {}

// IsServerlessEnv checks whether code is running in a serverless environment (any platform)
func IsServerlessEnv() bool {
	if IsLambdaEnv() || IsAzureFunctionEnv() || IsGcpFunctionEnv() {
		return true
	}
	return false
}
