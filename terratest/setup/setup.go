package setup

import "github.com/gruntwork-io/terratest/modules/terraform"

// Helper function to abstract away Terratest options setup.
func SetupTerraformOptions(tfDir string, vars map[string]interface{}, varFiles []string, awsProfile string, awsRegion string, teamName string) (terraformOptions *terraform.Options) {
	tags := map[string]string{"Team": teamName}
	vars["tags"] = tags

	terraformOptions = &terraform.Options{
		// Optionally use a different binary.
		//TerraformBinary: "/usr/local/bin/terraform12.3",

		// The path to where our Terraform code is located
		TerraformDir: tfDir,

		// Variables to pass to our Terraform code using -var options
		Vars: vars,

		// Variables to pass to our Terraform code using -var-file options.
		VarFiles: varFiles,

		// Environment variables to set when running Terraform
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
			"AWS_PROFILE":        awsProfile,
		},

		NoColor: true,
	}
	return
}
