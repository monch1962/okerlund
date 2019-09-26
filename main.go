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

// IsKubelessEnv checks whether code is running in a Kubeless function environment
func IsKubelessEnv() bool {
	if os.Getenv("REQ_MB_LIMIT") != "" { //Is there a better approach than this...?
		return true
	}
	return false
}

// IsGoogleCloudRunEnv checks whether code is running in a Google Cloud Run container environment
func IsGoogleCloudRunEnv() bool {
	if os.Getenv("K_SERVICE") != "" {
		return true
	}
	return false
}

// IsServerlessEnv checks whether code is running in a serverless environment (any platform)
func IsServerlessEnv() bool {
	if IsLambdaEnv() || IsAzureFunctionEnv() || IsGcpFunctionEnv() || IsKubelessEnv() || IsGoogleCloudRunEnv() {
		return true
	}
	return false
}
