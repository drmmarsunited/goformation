package intrinsics_test

import (
	"encoding/base64"
	"encoding/json"
	. "github.com/drmmarsunited/goformation/v7/intrinsics"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AWS CloudFormation intrinsic function processing", func() {

	Context("with a template that contains invalid JSON", func() {
		input := `{`
		processed, err := ProcessJSON([]byte(input), nil)
		It("should fail to process the template", func() {
			Expect(processed).To(BeNil())
			Expect(err).ToNot(BeNil())
		})
	})

	Context("with a template that contains a 'Ref' intrinsic function", func() {

		input := `{"Parameters":{"FunctionTimeout":{"Type":"Number","Default":120}},"Resources":{"MyServerlessFunction":{"Type":"AWS::Serverless::Function","Properties":{"Runtime":"nodejs6.10","Timeout":{"Ref":"FunctionTimeout"}}}}}`
		processed, err := ProcessJSON([]byte(input), nil)
		It("should successfully process the template", func() {
			Expect(processed).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})

		var result interface{}
		err = json.Unmarshal(processed, &result)
		It("should be valid JSON, and marshal to a Go type", func() {
			Expect(processed).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		template := result.(map[string]interface{})
		resources := template["Resources"].(map[string]interface{})
		resource := resources["MyServerlessFunction"].(map[string]interface{})
		properties := resource["Properties"].(map[string]interface{})

		It("should have the correct values", func() {
			Expect(properties["Timeout"]).To(Equal(float64(120)))
			Expect(properties["Runtime"]).To(Equal("nodejs6.10"))
		})

	})

	Context("with a template that contains a 'Ref' intrinsic function with MaxLength param string", func() {

		input := `{"Parameters":{"Name":{"Type":"String", "Default":"test", "MaxLength": "10"}, "FunctionTimeout":{"Type":"Number","Default":120}},"Resources":{"MyServerlessFunction":{"Type":"AWS::Serverless::Function","Properties":{"Runtime":"nodejs6.10","Timeout":{"Ref":"FunctionTimeout"}}}}}`
		processed, err := ProcessJSON([]byte(input), nil)
		It("should successfully process the template", func() {
			Expect(err).Should(BeNil())
			Expect(processed).ShouldNot(BeNil())
		})

		var result interface{}
		err = json.Unmarshal(processed, &result)
		It("should be valid JSON, and marshal to a Go type", func() {
			Expect(processed).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		template := result.(map[string]interface{})
		parameters := template["Parameters"].(map[string]interface{})
		pName := parameters["Name"].(map[string]interface{})
		resources := template["Resources"].(map[string]interface{})
		resource := resources["MyServerlessFunction"].(map[string]interface{})
		properties := resource["Properties"].(map[string]interface{})

		It("should have the correct values", func() {
			Expect(pName["MaxLength"]).To(Equal(strconv.Itoa(10)))
			Expect(properties["Timeout"]).To(Equal(float64(120)))
			Expect(properties["Runtime"]).To(Equal("nodejs6.10"))
		})

	})

	Context("with a template that contains a 'Ref' intrinsic function of a pseudo parameter", func() {

		input := `{"Parameters": {"FunctionTimeout": {"Type": "Number", "Default": 120}}, "Resources": {"MyServerlessFunction": {"Type": "AWS::Serverless::Function", "Properties": {"Runtime": "nodejs6.10", "Timeout": {"Ref": "FunctionTimeout"}}}}, "Outputs" : {"Partition" : {"Description" : "Value of the AWS partition", "Value" : {"Ref": "AWS::Partition"}}}}`
		processed, err := ProcessJSON([]byte(input), nil)
		It("should successfully process the template", func() {
			Expect(processed).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})

		var result interface{}
		err = json.Unmarshal(processed, &result)
		It("should be valid JSON, and marshal to a Go type", func() {
			Expect(processed).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		template := result.(map[string]interface{})
		resources := template["Resources"].(map[string]interface{})
		resource := resources["MyServerlessFunction"].(map[string]interface{})
		properties := resource["Properties"].(map[string]interface{})
		outputs := template["Outputs"].(map[string]interface{})
		partitionOutput := outputs["Partition"].(map[string]interface{})

		It("should have the correct values", func() {
			Expect(properties["Timeout"]).To(Equal(float64(120)))
			Expect(properties["Runtime"]).To(Equal("nodejs6.10"))
			Expect(partitionOutput["Value"]).To(Equal("aws"))
		})

	})

	Context("with a template that contains a 'Ref' intrinsic function and parameter overrides", func() {

		input := `{"Parameters":{"FunctionTimeout":{"Type":"Number","Default":120}},"Resources":{"MyServerlessFunction":{"Type":"AWS::Serverless::Function","Properties":{"Runtime":"nodejs6.10","Timeout":{"Ref":"FunctionTimeout"}}}}}`
		opts := ProcessorOptions{
			ParameterOverrides: map[string]interface{}{"FunctionTimeout": float64(42)},
		}
		processed, err := ProcessJSON([]byte(input), &opts)
		It("should successfully process the template", func() {
			Expect(processed).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})

		var result interface{}
		err = json.Unmarshal(processed, &result)
		It("should be valid JSON, and marshal to a Go type", func() {
			Expect(processed).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		template := result.(map[string]interface{})
		resources := template["Resources"].(map[string]interface{})
		resource := resources["MyServerlessFunction"].(map[string]interface{})
		properties := resource["Properties"].(map[string]interface{})

		It("should have the correct values", func() {
			Expect(properties["Timeout"]).To(Equal(float64(42)))
			Expect(properties["Runtime"]).To(Equal("nodejs6.10"))
		})

	})

	Context("with a template that contains a 'Fn::GetAZs' intrinsic function", func() {

		const input = `{
			"Resources": {
				"ExampleResource": {
					"Type": "AWS::Example::Resource",
					"Properties": {
						"AZEmpty": { "Fn::GetAZs": "" },
						"AZDefault": { "Fn::GetAZs": { "Ref": "AWS::Region" } },
						"AZParam": { "Fn::GetAZs": "eu-west-1" },
						"FirstAZ": {
							"Fn::Select" : [ "0", { "Fn::GetAZs" : "eu-central-1" } ]
						}
					}
				}
			}
		}`
		processed, err := ProcessJSON([]byte(input), nil)
		It("should successfully process the template", func() {
			Expect(processed).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})

		var result interface{}
		err = json.Unmarshal(processed, &result)
		It("should be valid JSON, and marshal to a Go type", func() {
			Expect(processed).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		template := result.(map[string]interface{})
		resources := template["Resources"].(map[string]interface{})
		resource := resources["ExampleResource"].(map[string]interface{})
		properties := resource["Properties"].(map[string]interface{})

		It("should have the correct values", func() {
			Expect(properties["AZEmpty"]).To(Equal([]interface{}{"us-east-1a", "us-east-1b", "us-east-1c", "us-east-1d"}))
			Expect(properties["AZDefault"]).To(Equal([]interface{}{"us-east-1a", "us-east-1b", "us-east-1c", "us-east-1d"}))
			Expect(properties["AZParam"]).To(Equal([]interface{}{"eu-west-1a", "eu-west-1b", "eu-west-1c"}))
			Expect(properties["FirstAZ"]).To(Equal("eu-central-1a"))
		})

	})

	Context("with a template that contains a 'Ref' intrinsic function with a 'Fn::Join' inside it", func() {

		input := `{"Parameters":{"FunctionTimeout":{"Type":"Number","Default":120}},"Resources":{"MyServerlessFunction":{"Type":"AWS::Serverless::Function","Properties":{"Runtime":"nodejs6.10","Timeout":{"Ref":{"Fn::Join":["Function","Timeout"]}}}}}}`
		processed, err := ProcessJSON([]byte(input), nil)
		It("should successfully process the template", func() {
			Expect(processed).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})

		var result interface{}
		err = json.Unmarshal(processed, &result)
		It("should be valid JSON, and marshal to a Go type", func() {
			Expect(processed).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		template := result.(map[string]interface{})
		resources := template["Resources"].(map[string]interface{})
		resource := resources["MyServerlessFunction"].(map[string]interface{})
		properties := resource["Properties"].(map[string]interface{})

		It("should have the correct values", func() {
			Expect(properties["Timeout"]).To(Equal(float64(120)))
			Expect(properties["Runtime"]).To(Equal("nodejs6.10"))
		})

	})

	Context("with a YAML template that contains intrinsic functions in tag form", func() {

		t := "AWSTemplateFormatVersion: '2010-09-09'\n"
		t += "Transform: AWS::Serverless-2016-10-31\n"
		t += "Description: SAM template for testing intrinsic functions with YAML tags\n"
		t += "Parameters:\n"
		t += "  TestParameter:\n"
		t += "    Type: String\n"
		t += "    Default: test-parameter-value\n"
		t += "Resources:\n"
		t += "  CodeUriWithS3LocationSpecifiedAsString:\n"
		t += "    Type: AWS::Serverless::Function\n"
		t += "    Properties:\n"
		t += "      Runtime: !Sub test-${ThisWontResolve}\n"
		t += "      Timeout: !Ref ThisWontResolve\n"
		t += "      FunctionName: !Ref TestParameter\n"
		t += "      Handler: !Sub method.${TestParameter}.${ThisWontResolve}\n"
		t += "      Role: !Sub ${ThisWontResolve.Arn}\n"
		t += "      CodeUri:\n"
		t += "        Fn::Sub:\n"
		t += "          - s3://${Bucket}/${Key}\n"
		t += "          - { \"Bucket\": \"test-bucket\", \"Key\": \"test-key\" }\n"

		processed, err := ProcessYAML([]byte(t), nil)
		It("should successfully process the template", func() {
			Expect(processed).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})

		var result interface{}
		err = json.Unmarshal(processed, &result)
		It("should be valid JSON, and marshal to a Go type", func() {
			Expect(processed).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		template := result.(map[string]interface{})
		resources := template["Resources"].(map[string]interface{})
		resource := resources["CodeUriWithS3LocationSpecifiedAsString"].(map[string]interface{})
		properties := resource["Properties"].(map[string]interface{})

		It("should convert an unresolvable !Ref to nil", func() {
			Expect(properties["Timeout"]).To(BeNil())
		})

		It("should handle unresolvable references in !Sub", func() {
			Expect(properties["Runtime"]).To(Equal("test-"))
		})

		It("should handle Fn::Sub with an array of replacements", func() {
			Expect(properties["CodeUri"]).To(Equal("s3://test-bucket/test-key"))
		})

		It("should resolved a !Ref that points to a valid parameter", func() {
			Expect(properties["FunctionName"]).To(Equal("test-parameter-value"))
		})

		It("should resolve a reference within a !Sub", func() {
			Expect(properties["Handler"]).To(Equal("method.test-parameter-value."))
		})

		// GetAtt isn't implemented today
		It("should resolve a resource attribute within a !Sub", func() {
			Expect(properties["Role"]).To(Equal("dummyvalue"))
		})

	})

	Context("with a template that contains primitives, intrinsics, and nested intrinsics", func() {

		const template = `{
			"Mappings" : {
				"RegionMap" : {
					"us-east-1" : { "32" : "ami-6411e20d", "64" : "ami-7a11e213" },
					"us-west-1" : { "32" : "ami-c9c7978c", "64" : "ami-cfc7978a" },
					"eu-west-1" : { "32" : "ami-37c2f643", "64" : "ami-31c2f645" },
					"ap-southeast-1" : { "32" : "ami-66f28c34", "64" : "ami-60f28c32" },
					"ap-northeast-1" : { "32" : "ami-9c03a89d", "64" : "ami-a003a8a1" }
				}
			},
			"Resources": {
				"ExampleResource": {
					"Type": "AWS::Example::Resource",
					"Properties": {
						"StringProperty": "Simple string example",
						"BooleanProperty": true,
						"NumberProperty": 123.45,
						"JoinIntrinsicPropertyString": { "Fn::Join": [ "some", "name" ] },
						"JoinIntrinsicPropertyArray": { "Fn::Join": [ "-", [ "some", "hyphenated", "name" ] ] },
						"JoinNestedIntrinsicProperty": { "Fn::Join": [ "some", { "Fn::Join": [ "joined", "value" ] } ] },
						"SubIntrinsicProperty": { "Fn::Sub": [ "some ${replaced}", { "replaced": "value" } ] },
						"SplitIntrinsicProperty": { "Fn::Split" : [ ",", "some,string,to,be,split" ] },
						"SelectIntrinsicProperty": { "Fn::Select" : [ 1, [ 0, 1, 2, 3 ] ] },
						"FindInMapIntrinsicProperty": { "Fn::FindInMap" : [ "RegionMap", "eu-west-1", "64"] },
						"Base64IntrinsicProperty": { "Fn::Base64" : "some-string-to-base64" },
						"RefAWSAccountId": { "Ref": "AWS::AccountId" },
						"RefAWSNotificationARNs": { "Ref": "AWS::NotificationARNs" },
						"RefNoValue": { "Ref": "AWS::NoValue" },
						"RefAWSRegion": { "Ref": "AWS::Region" },
						"RefAWSStackId": { "Ref": "AWS::StackId" },
						"RefAWSStackName": { "Ref": "AWS::StackName" }
					}
				}
			}
		}`

		Context("with no processor options", func() {

			processed, err := ProcessJSON([]byte(template), nil)
			It("should successfully process the template", func() {
				Expect(processed).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
			})

			var result interface{}
			err = json.Unmarshal(processed, &result)
			It("should be valid JSON, and marshal to a Go type", func() {
				Expect(processed).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			template := result.(map[string]interface{})
			resources := template["Resources"].(map[string]interface{})
			resource := resources["ExampleResource"].(map[string]interface{})
			properties := resource["Properties"].(map[string]interface{})

			It("should have the correct value for a primitive string property", func() {
				Expect(properties["StringProperty"]).To(Equal("Simple string example"))
			})

			It("should have the correct value for a primitive boolean property", func() {
				Expect(properties["BooleanProperty"]).To(Equal(true))
			})

			It("should have the correct value for a primitive number property", func() {
				Expect(properties["NumberProperty"]).To(Equal(123.45))
			})

			It("should have the correct value for a Fn::Join intrinsic property with an array of strings", func() {
				Expect(properties["JoinIntrinsicPropertyString"]).To(Equal("somename"))
			})

			It("should have the correct value for a Fn::Join intrinsic property with a delimiter and an array of strings", func() {
				Expect(properties["JoinIntrinsicPropertyArray"]).To(Equal("some-hyphenated-name"))
			})

			It("should have the correct value for a nested Fn::Join intrinsic property", func() {
				Expect(properties["JoinNestedIntrinsicProperty"]).To(Equal("somejoinedvalue"))
			})

			It("should have the correct value for a Fn::Sub intrinsic property", func() {
				Expect(properties["SubIntrinsicProperty"]).To(Equal("some value"))
			})

			It("should have the correct value for a Fn::Split intrinsic property", func() {
				Expect(properties["SplitIntrinsicProperty"]).To(HaveLen(5))
				results := []string{}
				for _, element := range properties["SplitIntrinsicProperty"].([]interface{}) {
					if str, ok := element.(string); ok {
						results = append(results, str)
					}
				}
				Expect(results).To(Equal([]string{"some", "string", "to", "be", "split"}))
			})

			It("should have the correct value for a Fn::Select intrinsic property", func() {
				Expect(properties["SelectIntrinsicProperty"]).To(Equal(float64(1)))
			})

			It("should have the correct value for a Fn::FindInMap intrinsic property", func() {
				Expect(properties["FindInMapIntrinsicProperty"]).To(Equal("ami-31c2f645"))
			})

			It("should have the correct value for a Fn::Base64 intrinsic property", func() {
				encoded, ok := properties["Base64IntrinsicProperty"].(string)
				Expect(ok).ToNot(BeNil())
				decoded, err := base64.StdEncoding.DecodeString(encoded)
				Expect(string(decoded)).To(Equal("some-string-to-base64"))
				Expect(err).To(BeNil())
			})

			It("should have the correct value for a AWS::AccountId pseudo parameter intrinsic property", func() {
				Expect(properties["RefAWSAccountId"]).To(Equal("123456789012"))
			})

			It("should have the correct value for a AWS::NotificationARNs pseudo parameter intrinsic property", func() {
				Expect(properties["RefAWSNotificationARNs"]).To(ContainElement("arn:aws:sns:us-east-1:123456789012:MyTopic"))
			})

			It("should have the correct value for a AWS::NoValue pseudo parameter intrinsic property", func() {
				Expect(properties["RefAWSNoValue"]).To(BeNil())
			})

			It("should have the correct value for a AWS::Region pseudo parameter intrinsic property", func() {
				Expect(properties["RefAWSRegion"]).To(Equal("us-east-1"))
			})

			It("should have the correct value for a AWS::StackId pseudo parameter intrinsic property", func() {
				Expect(properties["RefAWSStackId"]).To(Equal("arn:aws:cloudformation:us-east-1:123456789012:stack/MyStack/1c2fa620-982a-11e3-aff7-50e2416294e0"))
			})

			It("should have the correct value for a AWS::StackName pseudo parameter intrinsic property", func() {
				Expect(properties["RefAWSStackName"]).To(Equal("goformation-stack"))
			})

		})

		Context("with a processor options override for the Fn::Join function", func() {

			opts := &ProcessorOptions{
				IntrinsicHandlerOverrides: map[string]IntrinsicHandler{
					"Fn::Join": func(name string, input interface{}, template interface{}) interface{} {
						return "overridden"
					},
				},
			}

			processed, err := ProcessJSON([]byte(template), opts)
			It("should successfully process the template", func() {
				Expect(processed).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
			})

			result := map[string]interface{}{}
			err = json.Unmarshal(processed, &result)

			It("should be valid JSON, and marshal to a Go type", func() {
				Expect(processed).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			resources := result["Resources"].(map[string]interface{})
			resource := resources["ExampleResource"].(map[string]interface{})
			properties := resource["Properties"].(map[string]interface{})

			It("should have the correct value for a primitive string property", func() {
				Expect(properties["StringProperty"]).To(Equal("Simple string example"))
			})

			It("should have the correct value for a primitive boolean property", func() {
				Expect(properties["BooleanProperty"]).To(Equal(true))
			})

			It("should have the correct value for a primitive number property", func() {
				Expect(properties["NumberProperty"]).To(Equal(123.45))
			})

			It("should have the correct value for an intrinsic property", func() {
				Expect(properties["JoinIntrinsicPropertyString"]).To(Equal("overridden"))
			})

			It("should have the correct value for a nested intrinsic property", func() {
				Expect(properties["JoinNestedIntrinsicProperty"]).To(Equal("overridden"))
			})

			It("should have the correct value for an intrinsic property that's not supposed to be overridden", func() {
				Expect(properties["SubIntrinsicProperty"]).To(Equal("some value"))
			})

		})

	})

	Context("with a template that contains primitives, intrinsics, and nested intrinsics 2", func() {

		const template = `{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "The CloudFormation template for AWS resources required by AWS RedShift.",
    "Metadata": {
        "AWS::CloudFormation::Interface": {
            "ParameterGroups": [
                {
                    "Label": {
                        "default": "Tag Configuration"
                    },
                    "Parameters": [
                        "pSystemOwner",
                        "pApplicationName",
                        "pApplicationNameLC",
                        "pSubDivision",
                        "pDataSteward",
                        "pDataSensitivity",
                        "pResourceOwner",
                        "pResourceExposure",
                        "pResourceSensitivity",
                        "pPurpose",
                        "pProject",
                        "pSupportGroup",
                        "pAppGroup"
                    ]
                },
                {
                    "Label": {
                        "default": "Redshift Configuration"
                    },
                    "Parameters": [
                        "pPermissionBoundaryArn",
                        "pAutomatedSnapshotRetentionPeriod",
                        "pClusterIdentifier",
                        "pClusterSubnetGroupDescription",
                        "pClusterSubnetId1",
                        "pClusterSubnetId2",
                        "pClusterType",
                        "pDBName",
                        "pEncrypted",
                        "pEnhancedVpcRouting",
                        "pMasterUsername",
                        "pMasterUserPassword",
                        "pNodeType",
                        "pNumberOfNodes",
                        "pPort",
                        "pBucketName",
                        "pS3KeyPrefix"
                    ]
                }
            ],
            "ParameterLabels": {
                "pSystemOwner": {
                    "default": "System Owner"
                },
                "pApplicationName": {
                    "default": "Application Name"
                },
                "pApplicationNameLC": {
                    "default": "Application Name Lowercase"
                },
                "pSubDivision": {
                    "default": "Sub Division"
                },
                "pDataSteward": {
                    "default": "Data Steward"
                },
                "pDataSensitivity": {
                    "default": "Data Sensitivity"
                },
                "pResourceOwner": {
                    "default": "Resource Owner"
                },
                "pResourceExposure": {
                    "default": "Resource Exposure"
                },
                "pResourceSensitivity": {
                    "default": "Resource Sensitivity"
                },
                "pPurpose": {
                    "default": "Purpose of the resource"
                },
                "pProject": {
                    "Default": "Project"
                },
                "pSupportGroup": {
                    "Default": "Support Group"
                },
                "pAppGroup": {
                    "Default": "Application Group"
                },
                "pPermissionBoundaryArn": {
                    "default": "Permission Boundary Arn"
                },
                "pAutomatedSnapshotRetentionPeriod": {
                    "default": "Automated Snapshot Retention Period"
                },
                "pClusterIdentifier": {
                    "default": "Cluster Identifier"
                },
                "pClusterSubnetGroupDescription": {
                    "default": "Cluster Subnet Group Description"
                },
                "pClusterSubnetId1": {
                    "default": "Cluster Subnet Id 1"
                },
                "pClusterSubnetId2": {
                    "default": "Cluster Subnet Id 2"
                },
                "pClusterType": {
                    "default": "Cluster Type"
                },
                "pDBName": {
                    "default": "DB Name"
                },
                "pEncrypted": {
                    "default": "Encrypted"
                },
                "pEnhancedVpcRouting": {
                    "default": "Enhanced Vpc Routing"
                },
                "pMasterUsername": {
                    "default": "Master Username"
                },
                "pMasterUserPassword": {
                    "default": "Master User Password"
                },
                "pNodeType": {
                    "default": "Node Type"
                },
                "pNumberOfNodes": {
                    "default": "Number Of Nodes"
                },
                "pPort": {
                    "default": "Port"
                },
                "pBucketName": {
                    "default": "Bucket Name"
                },
                "pS3KeyPrefix": {
                    "default": "S3 Key Prefix"
                }
            }
        }
    },
    "Parameters": {
        "pSystemOwner": {
            "Type": "String",
            "Description": "Enter the ID of the system owner If you are not sure then use the requestor from the account request spreadsheet.\nIt cannot be more than 20 char, and cannot include [ ] : ; | = + * ? < > / \\ , @ or non printable characters\n",
            "AllowedPattern": "(^[^\\x00-\\x21\\x7f-\\xff\\[\\]:;|=+*?<>/\\\\,@]{1,20}$)",
            "Default": "CC"
        },
        "pApplicationName": {
            "Type": "String",
            "Description": "Enter the application name. If you are not sure then use the value from the project name in GitLab.",
            "ConstraintDescription": "The application name must be between 1 and 8 characters in length.",
            "AllowedPattern": "(^[a-zA-Z0-9][a-zA-Z0-9._()-]{0,7}$)",
            "MinLength": "1",
            "MaxLength": "8",
            "Default": "TEST"
        },
        "pApplicationNameLC": {
            "Type": "String",
            "Description": "Enter the application name in Lowercase. If you are not sure then use the value from the project name in GitLab.",
            "ConstraintDescription": "The application name must be between 1 and 8 characters in length.",
            "AllowedPattern": "(^[a-z0-9][a-z0-9._()-]{0,7}$)",
            "MinLength": "1",
            "MaxLength": "8",
            "Default": "test"
        },
        "pSubDivision": {
            "Description": "Define SubDivision which the resources should be tagged with.",
            "Type": "String",
            "MinLength": 1,
            "Default": "-"
        },
        "pDataSteward": {
            "Type": "String",
            "Description": "Enter the Active Directory group or user that is the data steward of this solution. If you are not sure then use the requestor from the account request spreadsheet.",
            "AllowedPattern": "(^[^\\x00-\\x21\\x7f-\\xff\\[\\]:;|=+*?<>/\\\\,@]{1,20}$)",
            "Default": "CC"
        },
        "pDataSensitivity": {
            "Type": "String",
            "Description": "Select the data sensitivity of data hosted by this solution",
            "ConstraintDescription": "The data sensitivity must be an allowed value",
            "AllowedValues": [
                "Red",
                "Amber",
                "Green"
            ],
            "Default": "Amber"
        },
        "pResourceOwner": {
            "Type": "String",
            "Description": "Enter the Active Directory group or user that is the resource owner of this solution. If you are not sure then use the requestor from the account request spreadsheet.",
            "AllowedPattern": "(^[^\\x00-\\x21\\x7f-\\xff\\[\\]:;|=+*?<>/\\\\,@]{1,20}$)",
            "Default": "CC"
        },
        "pResourceExposure": {
            "Type": "String",
            "Description": "Select the most exposure that this solution is exposed to",
            "ConstraintDescription": "The resource exposure must be an allowed value",
            "AllowedValues": [
                "Public",
                "Partner",
                "Saas",
                "OtherSecCloud",
                "AwsGovCloud",
                "SecOnPremise"
            ],
            "Default": "SecOnPremise"
        },
        "pResourceSensitivity": {
            "Type": "String",
            "Description": "Select the sensitivity of resources in this solution",
            "ConstraintDescription": "The resource sensitivity must be an allowed value",
            "AllowedValues": [
                "Red",
                "Amber",
                "Green"
            ],
            "Default": "Amber"
        },
        "pPurpose": {
            "Type": "String",
            "Description": "Select the general purpose of the resource",
            "ConstraintDescription": "The resource purpose must be an allowed value",
            "AllowedValues": [
                "wwws",
                "appl",
                "daba",
                "devo",
                "netw",
                "secr",
                "comp",
                "dago"
            ],
            "Default": "appl"
        },
        "pProject": {
            "Type": "String",
            "Description": "Enter the project name which is used for naming and tagging project-applicationname-purpose",
            "Default": "cc"
        },
        "pSupportGroup": {
            "Type": "String",
            "Description": "Enter in the contact information for the group supporting the resources",
            "Default": "TEST.CC.Foundation@test.com"
        },
        "pAppGroup": {
            "Type": "String",
            "Description": "Enter in the acronym from https://confluence.test.com/display/TEST/TEST+AWS+Organization+Account+Details",
            "Default": "test"
        },
        "pPermissionBoundaryArn": {
            "Description": "This is permission boundary arn for IAM Role",
            "Type": "String",
            "Default": "arn:aws:iam::123456789012:policy/cc-ScopePermissions-test-Dev-Developer"
        },
        "pAutomatedSnapshotRetentionPeriod": {
            "Description": "The number of days that automated snapshots are retained. If the value is 0, automated snapshots are disabled. Even if automated snapshots are disabled, you can still create manual snapshots when you want",
            "Type": "Number",
            "Default": 7,
            "MinValue": 0,
            "MaxValue": 35
        },
        "pClusterIdentifier": {
            "Description": "A unique identifier for the cluster. You use this identifier to refer to the cluster for any subsequent cluster operations such as deleting or modifying. The identifier also appears in the Amazon Redshift console.",
            "Type": "String",
            "Default": ""
        },
        "pClusterSubnetGroupDescription": {
            "Description": "The Description of a cluster subnet group to be associated with this cluster.",
            "Type": "String",
            "Default": "Redshift Cluster Subnet Group Description"
        },
        "pClusterSubnetId1": {
            "Description": "The Cluster Subnet ID 1.",
            "Type": "String",
            "Default": "subnet-0c3636c28010a6107"
        },
        "pClusterSubnetId2": {
            "Description": "The Cluster Subnet ID 2.",
            "Type": "String",
            "Default": "subnet-06af5b42cc8fc9aa4"
        },
        "pClusterType": {
            "Description": "The type of cluster",
            "Type": "String",
            "Default": "single-node",
            "AllowedValues": [
                "single-node",
                "multi-node"
            ]
        },
        "pDBName": {
            "Description": "The name of the first database to be created when the cluster is created.",
            "Type": "String",
            "Default": "dev",
            "AllowedPattern": "([a-z]|[0-9])+"
        },
        "pEncrypted": {
            "Description": "If true, the data in the cluster is encrypted at rest.",
            "Type": "String",
            "Default": true,
            "AllowedValues": [
                true,
                false
            ]
        },
        "pEnhancedVpcRouting": {
            "Description": "An option that specifies whether to create the cluster with enhanced VPC routing enabled.",
            "Type": "String",
            "Default": true,
            "AllowedValues": [
                true,
                false
            ]
        },
        "pMasterUsername": {
            "Description": "The user name that is associated with the master user account for the cluster that is being created.",
            "Type": "String",
            "Default": "redshiftadmin",
            "AllowedPattern": "([a-z])([a-z]|[0-9])*"
        },
        "pMasterUserPassword": {
            "Description": "The password that is associated with the master user account for the cluster that is being created.",
            "Type": "String",
            "Default": "Redshiftadmin123!",
            "NoEcho": true
        },
        "pNodeType": {
            "Description": "The type of node to be provisioned",
            "Type": "String",
            "Default": "ra3.xlplus",
            "AllowedValues": [
                "ds2.xlarge",
                "ds2.8xlarge",
                "dc1.large",
                "dc1.8xlarge",
                "dc2.large",
                "dc2.8xlarge",
                "ra3.xlplus",
                "ra3.4xlarge",
                "ra3.16xlarge"
            ]
        },
        "pNumberOfNodes": {
            "Description": "The number of compute nodes in the cluster. For multi-node clusters, the NumberOfNodes parameter must be greater than 1",
            "Type": "Number",
            "Default": 1
        },
        "pPort": {
            "Description": "The port number on which the cluster accepts incoming connections.",
            "Type": "Number",
            "Default": 5439
        },
        "pBucketName": {
            "Description": "The name of an existing S3 bucket where the log files are to be stored.",
            "Type": "String",
            "Default": "test-redshift-audit-logging"
        },
        "pS3KeyPrefix": {
            "Description": "The prefix applied to the log file names.",
            "Type": "String",
            "Default": ""
        }
    },
    "Conditions": {
        "IsNullOrEmptyPermissionBoundaryArn": {
            "Fn::Or": [
                {
                    "Fn::Equals": [
                        {
                            "Ref": "pPermissionBoundaryArn"
                        },
                        ""
                    ]
                },
                {
                    "Fn::Equals": [
                        {
                            "Ref": "pPermissionBoundaryArn"
                        },
                        "AWS::NoValue"
                    ]
                }
            ]
        }
    },
    "Mappings": {
        "NameTagPartitionToPartitionAbbreviation": {
            "aws-us-gov": {
                "PartitionAbbr": "G"
            },
            "aws": {
                "PartitionAbbr": "C"
            }
        },
        "NameTagRegionToRegionAbbreviation": {
            "us-east-1": {
                "RegionAbbr": "USE1"
            },
            "us-east-2": {
                "RegionAbbr": "USE2"
            },
            "us-west-1": {
                "RegionAbbr": "USW1"
            },
            "us-west-2": {
                "RegionAbbr": "USW2"
            }
        }
    },
    "Resources": {
        "rRedshiftExecutionRole": {
            "Type": "AWS::IAM::Role",
            "Properties": {
                "RoleName": {
                    "Fn::Join": [
                        "-",
                        [
                            "{{resolve:ssm:/account/config/BusinessLowercase}}",
                            {
                                "Ref": "pApplicationNameLC"
                            },
                            "RedshiftExecutionRole"
                        ]
                    ]
                },
                "AssumeRolePolicyDocument": {
                    "Version": "2012-10-17",
                    "Statement": {
                        "Effect": "Allow",
                        "Principal": {
                            "Service": [
                                "redshift.amazonaws.com",
                                "resources.cloudformation.amazonaws.com",
                                "s3.amazonaws.com"
                            ]
                        },
                        "Action": "sts:AssumeRole"
                    }
                },
                "Path": "/",
                "PermissionsBoundary": {
                    "Fn::If": [
                        "IsNullOrEmptyPermissionBoundaryArn - !Ref \"AWS::NoValue\" - !Ref pPermissionBoundaryArn"
                    ]
                },
                "Policies": [
                    {
                        "PolicyName": "RedShiftLogsAndS3Policy",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Effect": "Allow",
                                    "Action": [
                                        "cloudwatch:PutMetricData",
                                        "logs:DescribeLogGroups",
                                        "logs:CreateLogGroup",
                                        "logs:DescribeLogStream*",
                                        "logs:CreateLogStream",
                                        "logs:PutLogEvents",
                                        "s3:ListBucket",
                                        "s3:ListBucketVersions",
                                        "s3:GetBucketAcl",
                                        "s3:GetObject",
                                        "s3:GetObjectAcl",
                                        "s3:GetObjectVersion",
                                        "s3:GetObjectVersionAcl"
                                    ],
                                    "Resource": "*"
                                }
                            ]
                        }
                    },
                    {
                        "PolicyName": "RedshiftAuditLogging",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": {
                                "Effect": "Allow",
                                "Action": [
                                    "s3:PutObject",
                                    "s3:PutObjectAcl"
                                ],
                                "Resource": [
                                    "arn:aws:s3:::${pBucketName}",
                                    "arn:aws:s3:::${pBucketName}/*",
                                    "arn:aws:s3:::${pBucketName}/AWSLogs",
                                    "arn:aws:s3:::${pBucketName}/AWSLogs/*"
                                ]
                            }
                        }
                    },
                    {
                        "PolicyName": "KMSAccess",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Effect": "Allow",
                                    "Action": [
                                        "kms:GetPublicKey",
                                        "kms:ListKeyPolicies",
                                        "kms:ListRetirableGrants",
                                        "kms:GetKeyPolicy",
                                        "kms:ListResourceTags",
                                        "kms:ListGrants",
                                        "kms:GetParametersForImport",
                                        "kms:DescribeCustomKeyStores",
                                        "kms:ListKeys",
                                        "kms:Encrypt",
                                        "kms:Decrypt",
                                        "kms:GetKeyRotationStatus",
                                        "kms:ListAliases",
                                        "kms:DescribeKey",
                                        "kms:GenerateDataKey"
                                    ],
                                    "Resource": [
                                        {
                                            "Fn::ImportValue": "S3GeneralKeyArn"
                                        }
                                    ]
                                }
                            ]
                        }
                    },
                    {
                        "PolicyName": "ResourceTypePolicy",
                        "PolicyDocument": {
                            "Version": "2012-10-17",
                            "Statement": [
                                {
                                    "Effect": "Allow",
                                    "Action": [
                                        "iam:PassRole",
                                        "redshift:Describe*",
                                        "redshift:List*",
                                        "redshift:View*",
                                        "redshift:Accept*",
                                        "redshift:Cancel*",
                                        "redshift:Create*",
                                        "redshift:*Tags",
                                        "redshift:ModifyCluster*",
                                        "redshift:Modify*",
                                        "redshift:Delete*",
                                        "redshift:Reboot*"
                                    ],
                                    "Resource": "*"
                                }
                            ]
                        }
                    }
                ],
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Fn::Sub": [
                                "A-${Partition}${Region}-${LifeCycle}-${DivisionName}-${pApplicationName}-${Purpose}-${ServiceType}-${RoleName}",
                                {
                                    "Partition": {
                                        "Fn::FindInMap": [
                                            "NameTagPartitionToPartitionAbbreviation",
                                            {
                                                "Ref": "AWS::Partition"
                                            },
                                            "PartitionAbbr"
                                        ]
                                    },
                                    "Region": {
                                        "Fn::FindInMap": [
                                            "NameTagRegionToRegionAbbreviation",
                                            {
                                                "Ref": "AWS::Region"
                                            },
                                            "RegionAbbr"
                                        ]
                                    },
                                    "LifeCycle": "{{resolve:ssm:/account/config/LifeCycleLowercase}}",
                                    "DivisionName": "{{resolve:ssm:/account/config/BusinessLowercase}}",
                                    "Purpose": {
                                        "Ref": "pPurpose"
                                    },
                                    "ServiceType": "IAM-Role",
                                    "RoleName": {
                                        "Fn::Sub": "${pApplicationNameLC}-RedshiftExecutionRole"
                                    }
                                }
                            ]
                        }
                    },
                    {
                        "Key": "LifeCycle",
                        "Value": "{{resolve:ssm:/account/config/LifeCycle}}"
                    },
                    {
                        "Key": "SystemOwner",
                        "Value": {
                            "Ref": "pSystemOwner"
                        }
                    },
                    {
                        "Key": "ApplicationName",
                        "Value": {
                            "Ref": "pApplicationName"
                        }
                    },
                    {
                        "Key": "Business",
                        "Value": "{{resolve:ssm:/account/config/Business}}"
                    },
                    {
                        "Key": "SupportBusiness",
                        "Value": "{{resolve:ssm:/account/config/SupportBusiness}}"
                    },
                    {
                        "Key": "SupportCCGroup",
                        "Value": "{{resolve:ssm:/account/config/SupportCCGroup}}"
                    },
                    {
                        "Key": "DataSteward",
                        "Value": {
                            "Ref": "pDataSteward"
                        }
                    },
                    {
                        "Key": "DataSensitivity",
                        "Value": {
                            "Ref": "pDataSensitivity"
                        }
                    },
                    {
                        "Key": "ResourceOwner",
                        "Value": {
                            "Ref": "pResourceOwner"
                        }
                    },
                    {
                        "Key": "ResourceSensitivity",
                        "Value": {
                            "Ref": "pResourceSensitivity"
                        }
                    },
                    {
                        "Key": "ResourceExposure",
                        "Value": {
                            "Ref": "pResourceExposure"
                        }
                    },
                    {
                        "Key": "SubDivision",
                        "Value": {
                            "Ref": "pSubDivision"
                        }
                    },
                    {
                        "Key": "Project",
                        "Value": {
                            "Ref": "pProject"
                        }
                    },
                    {
                        "Key": "SupportGroup",
                        "Value": {
                            "Ref": "pSupportGroup"
                        }
                    },
                    {
                        "Key": "AppGroup",
                        "Value": {
                            "Ref": "pAppGroup"
                        }
                    }
                ]
            }
        },
        "rRedshiftClusterSubnetGroup": {
            "Type": "AWS::Redshift::ClusterSubnetGroup",
            "Properties": {
                "Description": {
                    "Ref": "pClusterSubnetGroupDescription"
                },
                "SubnetIds": [
                    {
                        "Ref": "pClusterSubnetId1"
                    },
                    {
                        "Ref": "pClusterSubnetId2"
                    }
                ],
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Ref": "pClusterSubnetGroupDescription"
                        }
                    },
                    {
                        "Key": "LifeCycle",
                        "Value": "{{resolve:ssm:/account/config/LifeCycle}}"
                    },
                    {
                        "Key": "SystemOwner",
                        "Value": {
                            "Ref": "pSystemOwner"
                        }
                    },
                    {
                        "Key": "ApplicationName",
                        "Value": {
                            "Ref": "pApplicationName"
                        }
                    },
                    {
                        "Key": "Business",
                        "Value": "{{resolve:ssm:/account/config/Business}}"
                    },
                    {
                        "Key": "SupportBusiness",
                        "Value": "{{resolve:ssm:/account/config/SupportBusiness}}"
                    },
                    {
                        "Key": "SupportCCGroup",
                        "Value": "{{resolve:ssm:/account/config/SupportCCGroup}}"
                    },
                    {
                        "Key": "DataSteward",
                        "Value": {
                            "Ref": "pDataSteward"
                        }
                    },
                    {
                        "Key": "DataSensitivity",
                        "Value": {
                            "Ref": "pDataSensitivity"
                        }
                    },
                    {
                        "Key": "ResourceOwner",
                        "Value": {
                            "Ref": "pResourceOwner"
                        }
                    },
                    {
                        "Key": "ResourceSensitivity",
                        "Value": {
                            "Ref": "pResourceSensitivity"
                        }
                    },
                    {
                        "Key": "ResourceExposure",
                        "Value": {
                            "Ref": "pResourceExposure"
                        }
                    },
                    {
                        "Key": "SubDivision",
                        "Value": {
                            "Ref": "pSubDivision"
                        }
                    },
                    {
                        "Key": "Project",
                        "Value": {
                            "Ref": "pProject"
                        }
                    },
                    {
                        "Key": "SupportGroup",
                        "Value": {
                            "Ref": "pSupportGroup"
                        }
                    },
                    {
                        "Key": "AppGroup",
                        "Value": {
                            "Ref": "pAppGroup"
                        }
                    }
                ]
            }
        },
        "rRedshiftClusterParameterGroup": {
            "Type": "AWS::Redshift::ClusterParameterGroup",
            "Properties": {
                "Description": "Cluster parameter group",
                "ParameterGroupFamily": "redshift-1.0",
                "Parameters": [
                    {
                        "ParameterName": "enable_user_activity_logging",
                        "ParameterValue": "true"
                    },
                    {
                        "ParameterName": "require_ssl",
                        "ParameterValue": "true"
                    }
                ]
            }
        },
        "rRedshiftCluster": {
            "Type": "AWS::Redshift::Cluster",
            "Properties": {
                "AutomatedSnapshotRetentionPeriod": {
                    "Ref": "pAutomatedSnapshotRetentionPeriod"
                },
                "ClusterIdentifier": {
                    "Ref": "pClusterIdentifier"
                },
                "ClusterParameterGroupName": {
                    "Ref": "rRedshiftClusterParameterGroup"
                },
                "ClusterSubnetGroupName": {
                    "Fn::GetAtt": [
                        "rRedshiftClusterSubnetGroup",
                        "ClusterSubnetGroupName"
                    ]
                },
                "ClusterType": {
                    "Ref": "pClusterType"
                },
                "DBName": {
                    "Ref": "pDBName"
                },
                "Encrypted": {
                    "Ref": "pEncrypted"
                },
                "EnhancedVpcRouting": {
                    "Ref": "pEnhancedVpcRouting"
                },
                "IamRoles": [
                    {
                        "Fn::GetAtt": [
                            "rRedshiftExecutionRole",
                            "Arn"
                        ]
                    }
                ],
                "KmsKeyId": {
                    "Fn::ImportValue": "RedshiftKeyId"
                },
                "MasterUsername": {
                    "Ref": "pMasterUsername"
                },
                "MasterUserPassword": {
                    "Ref": "pMasterUserPassword"
                },
                "NodeType": {
                    "Ref": "pNodeType"
                },
                "NumberOfNodes": {
                    "Ref": "pNumberOfNodes"
                },
                "Port": {
                    "Ref": "pPort"
                },
                "PubliclyAccessible": false,
                "LoggingProperties": {
                    "BucketName": {
                        "Ref": "pBucketName"
                    },
                    "S3KeyPrefix": {
                        "Ref": "pS3KeyPrefix"
                    }
                },
                "Tags": [
                    {
                        "Key": "Name",
                        "Value": {
                            "Fn::Sub": [
                                "A-${Partition}-${Region}-${LifeCycle}-${DivisionName}-${pApplicationName}-${Purpose}-${ServiceType}-001-_NONAME",
                                {
                                    "Partition": {
                                        "Fn::FindInMap": [
                                            "NameTagPartitionToPartitionAbbreviation",
                                            {
                                                "Ref": "AWS::Partition"
                                            },
                                            "PartitionAbbr"
                                        ]
                                    },
                                    "Region": {
                                        "Fn::FindInMap": [
                                            "NameTagRegionToRegionAbbreviation",
                                            {
                                                "Ref": "AWS::Region"
                                            },
                                            "RegionAbbr"
                                        ]
                                    },
                                    "LifeCycle": "{{resolve:ssm:/account/config/LifeCycleLowercase}}",
                                    "DivisionName": "{{resolve:ssm:/account/config/BusinessLowercase}}",
                                    "pApplicationName": {
                                        "Ref": "pApplicationName"
                                    },
                                    "Purpose": "Redshift Cluster",
                                    "ServiceType": "Redshift Cluster"
                                }
                            ]
                        }
                    },
                    {
                        "Key": "LifeCycle",
                        "Value": "{{resolve:ssm:/account/config/LifeCycle}}"
                    },
                    {
                        "Key": "SystemOwner",
                        "Value": {
                            "Ref": "pSystemOwner"
                        }
                    },
                    {
                        "Key": "ApplicationName",
                        "Value": {
                            "Ref": "pApplicationName"
                        }
                    },
                    {
                        "Key": "Business",
                        "Value": "{{resolve:ssm:/account/config/Business}}"
                    },
                    {
                        "Key": "SupportBusiness",
                        "Value": "{{resolve:ssm:/account/config/SupportBusiness}}"
                    },
                    {
                        "Key": "SupportCCGroup",
                        "Value": "{{resolve:ssm:/account/config/SupportCCGroup}}"
                    },
                    {
                        "Key": "DataSteward",
                        "Value": {
                            "Ref": "pDataSteward"
                        }
                    },
                    {
                        "Key": "DataSensitivity",
                        "Value": {
                            "Ref": "pDataSensitivity"
                        }
                    },
                    {
                        "Key": "ResourceOwner",
                        "Value": {
                            "Ref": "pResourceOwner"
                        }
                    },
                    {
                        "Key": "ResourceSensitivity",
                        "Value": {
                            "Ref": "pResourceSensitivity"
                        }
                    },
                    {
                        "Key": "ResourceExposure",
                        "Value": {
                            "Ref": "pResourceExposure"
                        }
                    },
                    {
                        "Key": "SubDivision",
                        "Value": {
                            "Ref": "pSubDivision"
                        }
                    },
                    {
                        "Key": "Project",
                        "Value": {
                            "Ref": "pProject"
                        }
                    },
                    {
                        "Key": "SupportGroup",
                        "Value": {
                            "Ref": "pSupportGroup"
                        }
                    },
                    {
                        "Key": "AppGroup",
                        "Value": {
                            "Ref": "pAppGroup"
                        }
                    }
                ]
            }
        }
    },
    "Outputs": {
        "oRedshiftExecutionRoleArn": {
            "Description": "IAM Role for Executing Redshift",
            "Value": {
                "Fn::GetAtt": [
                    "rRedshiftExecutionRole",
                    "Arn"
                ]
            }
        },
        "oRedshiftClusterSubnetGroupName": {
            "Description": "Redshift Cluster Subnet Group Name",
            "Value": {
                "Fn::GetAtt": [
                    "rRedshiftClusterSubnetGroup",
                    "ClusterSubnetGroupName"
                ]
            }
        },
        "oRedshiftClusterName": {
            "Description": "The Name of the Redshift Cluster",
            "Value": {
                "Ref": "rRedshiftCluster"
            }
        }
    }
}`

		Context("with no processor options", func() {

			processed, err := ProcessJSON([]byte(template), nil)
			It("should successfully process the template", func() {
				Expect(processed).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
			})

			var result interface{}
			err = json.Unmarshal(processed, &result)
			It("should be valid JSON, and marshal to a Go type", func() {
				Expect(processed).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			template := result.(map[string]interface{})
			resources := template["Resources"].(map[string]interface{})
			resource := resources["rRedshiftCluster"].(map[string]interface{})
			properties := resource["Properties"].(map[string]interface{})
			tags := properties["Tags"].([]interface{})

			It("should have the correct value for a primitive integer property", func() {
				Expect(properties["Port"]).To(Equal(float64(5439)))
			})

			It("should have the correct length for the Tags property", func() {
				Expect(properties["Tags"]).To(HaveLen(16))
			})

			It("should have the correct value for the Tags[0] --> Key", func() {
				Expect(tags[0].(map[string]interface{})["Key"]).To(Equal("Name"))
			})

			It("should have the correct value for the Tags --> Name value", func() {
				Expect(tags[0].(map[string]interface{})["Value"]).To(Equal("A-C-USE1-{{resolve:ssm:/account/config/LifeCycleLowercase}}-{{resolve:ssm:/account/config/BusinessLowercase}}-TEST-Redshift Cluster-Redshift Cluster-001-_NONAME"))
			})

		})

	})

	Context("with a processor options that has NoProcess set", func() {

		input := `{"Resources":{"MyBucket":{"Type":"AWS::S3::Bucket","Properties":{"BucketName":{"Ref":"BucketNameParameter"}}}}}`

		opts := &ProcessorOptions{
			NoProcess: true,
		}
		processed, err := ProcessJSON([]byte(input), opts)
		It("should successfully process the template", func() {
			Expect(processed).ShouldNot(BeNil())
			Expect(err).Should(BeNil())
		})

		var result interface{}
		err = json.Unmarshal(processed, &result)
		It("should be valid JSON, and marshal to a Go type", func() {
			Expect(processed).ToNot(BeNil())
			Expect(err).To(BeNil())
		})

		template := result.(map[string]interface{})
		resources := template["Resources"].(map[string]interface{})
		resource := resources["MyBucket"].(map[string]interface{})
		properties := resource["Properties"].(map[string]interface{})
		bucketName := properties["BucketName"].(map[string]interface{})

		It("should have an unprocessed Ref", func() {
			Expect(bucketName["Ref"]).To(Equal("BucketNameParameter"))
		})

	})

	Context("with a template that contains intrinsics and conditions", func() {

		const template = `{
					"Parameters" : {
						"UseBucket":{"Type":"String","Default":"false"}
					},
					"Conditions" : {
						"UseBucketCondition" : { "Fn::Equals": [{"Ref" : "UseBucket"}, "true"]},
						"WithIntrinsic" :      { "Fn::Equals": [{"Ref" : "UseBucket"}, { "Fn::Join": [ "fal", "se" ] }]},
						"NonExistant" : { "Fn::Not": [{"Condition" : "NotAvailable"}]},
						"Eq" : { "Fn::Equals": [{"Ref" : "AWS::AccountId"}, "123456789012"]},
						"EqF" : { "Fn::Equals": [{"Ref" : "AWS::AccountId"}, "Foo"]},
						"NotInline" : { "Fn::Not" : [{ "Fn::Equals": [{"Ref" : "AWS::AccountId"}, "123456789012"]}]},
						"And" : { "Fn::And" : [{ "Condition": "Eq"}, { "Condition": "EqF"}]},
						"AndT" : { "Fn::And" : [{ "Condition": "Eq"}, { "Condition": "WithIntrinsic"}]},
						"Or" : { "Fn::Or" : [{ "Condition": "Eq"}, { "Condition": "EqF"}]},
						"OrF" : { "Fn::Or" : [{ "Condition": "EqF"}, { "Condition": "UseBucketCondition"}]},
						"NotRef" : { "Fn::Not" : [{ "Condition": "Eq"}]}
					},
					"Resources": {
						"ExampleResource": {
							"Type": "AWS::Example::Resource",
							"Properties": {
								"EqProp": { "Fn::If": [ "Eq", "OK", "false" ] },
								"Bucket": { "Fn::If": [ "UseBucketCondition", "Bucket", "NotBucket" ] },
								"NotProp": { "Fn::If": [ "NotInline", "false", "OK" ] },
								"NonExistant": { "Fn::If": [ "NonExistant", "false", "OK" ] }
							}
						}
					}
				}`

		Context("with evaluate conditions processor option", func() {

			processed, err := ProcessJSON([]byte(template), &ProcessorOptions{EvaluateConditions: true})
			It("should successfully process the template", func() {
				Expect(processed).ShouldNot(BeNil())
				Expect(err).Should(BeNil())
			})

			var result interface{}
			err = json.Unmarshal(processed, &result)
			It("should be valid JSON, and marshal to a Go type", func() {
				Expect(processed).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			template := result.(map[string]interface{})
			resources := template["Resources"].(map[string]interface{})
			conditions := template["Conditions"].(map[string]interface{})
			resource := resources["ExampleResource"].(map[string]interface{})
			properties := resource["Properties"].(map[string]interface{})

			It("should have the correct value for a equals condition", func() {
				Expect(conditions["Eq"]).To(Equal(true))
				Expect(conditions["EqF"]).To(Equal(false))
			})

			It("should have the correct value for a not condition", func() {
				Expect(conditions["NotInline"]).To(Equal(false))
				Expect(conditions["NotRef"]).To(Equal(false))
			})

			It("should have the correct value for a and condition that references another condition", func() {
				Expect(conditions["And"]).To(Equal(false))
				Expect(conditions["AndT"]).To(Equal(true))
			})

			It("should have the correct value for a or condition that references another condition", func() {
				Expect(conditions["Or"]).To(Equal(true))
				Expect(conditions["OrF"]).To(Equal(false))
			})

			It("should have the correct value for a condition that uses intrinsics", func() {
				Expect(conditions["WithIntrinsic"]).To(Equal(true))
			})

			It("should have a nil value for a non-existant condition", func() {
				Expect(conditions["NonExistant"]).To(BeNil())
			})

			It("should have the correct value for a if property", func() {
				Expect(properties["EqProp"]).To(Equal("OK"))
				Expect(properties["NotProp"]).To(Equal("OK"))
				Expect(properties["Bucket"]).To(Equal("NotBucket"))
				Expect(properties["NonExistant"]).To(BeNil())
			})
		})
	})
})
