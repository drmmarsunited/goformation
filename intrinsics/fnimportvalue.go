package intrinsics

// FnImportValue is not implemented, and always returns "dummyvalue.
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-importvalue.html
func FnImportValue(name string, input interface{}, template interface{}) interface{} {

	// { "Fn::ImportValue" : sharedValueToImport }
	return "dummyvalue"
}
