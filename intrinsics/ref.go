package intrinsics

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/drmmarsunited/goformation/v7/cloudformation/utils"
	"strings"
)

var pph = newPseudoParamHelper()

// Ref resolves the 'Ref' AWS CloudFormation intrinsic function.
// Currently, this only resolves against CloudFormation Parameter default values
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html
func Ref(name string, input interface{}, template interface{}) interface{} {
	// Dang son, this has got more nest than a bald eagle
	// Check the input is a string
	if name, ok := input.(string); ok {

		switch name {

		case "AWS::AccountId":
			pph.parseAwsAccountId()
			return pph.accountId
		case "AWS::NotificationARNs":
			// This data is not knowable outside the context of the AWS CloudFormation service, using a dummy value
			return []string{"arn:aws:sns:us-east-1:123456789012:MyTopic"}
		case "AWS::NoValue":
			return nil
		case "AWS::Partition":
			pph.parseAwsPartition()
			return pph.partition
		case "AWS::Region":
			pph.parseAwsRegion()
			return pph.region
		case "AWS::StackId":
			// This data is not knowable outside the context of the AWS CloudFormation service, using a dummy value
			return "arn:aws:cloudformation:us-east-1:123456789012:stack/MyStack/1c2fa620-982a-11e3-aff7-50e2416294e0"
		case "AWS::StackName":
			// This data is not knowable outside the context of the AWS CloudFormation service, using a dummy value
			return "goformation-stack"

		default:

			// This isn't a pseudo 'Ref' parameter, so we need to look inside the CloudFormation template
			// to see if we can resolve the reference. This implementation just looks at the Parameters section
			// to see if there is a parameter matching the name, and if so, return the default value.

			// Check the template is a map
			if template, ok := template.(map[string]interface{}); ok {
				// Check there is a parameters section
				if uparameters, ok := template["Parameters"]; ok {
					// Check the parameters section is a map
					if parameters, ok := uparameters.(map[string]interface{}); ok {
						// Check there is a parameter with the same name as the Ref
						if uparameter, ok := parameters[name]; ok {
							// Check the parameter is a map
							if parameter, ok := uparameter.(map[string]interface{}); ok {
								// Check the parameter has a default
								if def, ok := parameter["Default"]; ok {
									if strings.Contains(parameter["Type"].(string), "CommaDelimitedList") {
										return strings.Split(def.(string), ",")
									}

									if strings.Contains(parameter["Type"].(string), "AWS::SSM::Parameter::Value") {
										// Set up AWS helper
										awsHelper, err := utils.NewAwsHelper()
										if err != nil {
											fmt.Printf("Could not setup AWS Helper: %s\n", err.Error())
											return "dummyvalue"
										}

										// Set up SSM client
										svc := ssm.NewFromConfig(awsHelper.Cfg)

										// Get the SSM parameter value
										pn := def.(string)
										res, err := svc.GetParameter(context.TODO(), &ssm.GetParameterInput{Name: &pn})
										if err != nil {
											fmt.Printf("Could not get SSM parameter %s: %s\n", pn, err.Error())
											return "dummyvalue"
										}

										return *res.Parameter.Value
									}

									return def
								}
							}
						}
					}
				}
			}
		}

	}

	return nil

}
