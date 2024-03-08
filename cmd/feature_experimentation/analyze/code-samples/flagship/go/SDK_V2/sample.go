package SDK_V2

func main() {
	// Using the Decision API (default)
	fsClient, _ := flagship.Start("environmentID", "apiKey")

	// Using the Bucketing mode
	fsClient, _ = flagship.Start("environmentID", "apiKey", client.WithBucketing())

	// Create visitor context
	context := map[string]interface{}{
		"isVip": true,
		"age":   30,
		"name":  "visitor",
	}
	// Create a visitor
	fsVisitor, _ := fsClient.NewVisitor("visitor_id", context)

	// Update a single key
	fsVisitor.UpdateContextKey("vipUser", true)
	fsVisitor.UpdateContextKey("age", 30)

	// Update the whole context
	newContext := map[string]interface{}{
		"isVip": true,
		"age":   30,
		"name":  "visitor",
	}
	fsVisitor.UpdateContext(newContext)

	discountName, err := fsVisitor.GetModificationString("btnColor", "VString", true)
	discountName, err := fsVisitor.GetModificationNumber("btnSize", 13, true)
	discountName, err := fsVisitor.GetModificationBool("showBtn", false, true)

	// these flags will be analyzed using the custom-regex file example-regex.json
	// try the command: flagship analyze flag list --custom-regex-json ./cmd/analyze/flag/example-regex.json --directory ./cmd/analyze/code-samples/flagship/go/
	flagValue, _ := example.BoolVariation("my-boolean-flag", false)
	flagValue1, _ := example.StringVariation("my-string-flag", "defaltVal")
	flagValue2, _ := example.Float64Variation("my-numbers-flag", 13.6)
}
