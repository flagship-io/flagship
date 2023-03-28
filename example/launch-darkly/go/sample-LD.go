package SDK_V2

func main() {
	// Using the Decision API (default)
	featureFlagClient, _ := featureFlagInstance.Start("environmentID", "apiKey")

	// Create visitor context
	context := map[string]interface{}{
		"isVip": true,
		"age":   30,
		"name":  "visitor",
	}
	// Create a visitor
	featureFlagVisitor, _ := featureFlagClient.NewVisitor("visitor_id", context)

	// Update a single key
	featureFlagVisitor.UpdateContextKey("vipUser", true)
	featureFlagVisitor.UpdateContextKey("age", 30)

	// Update the whole context
	newContext := map[string]interface{}{
		"isVip": true,
		"age":   30,
		"name":  "visitor",
	}

	featureFlagVisitor.UpdateContext(newContext)

	flagValue, _ := featureFlagVisitor.BoolVariation("LD-Boolean-flag-go", context, false)
	flagValue1, _ := featureFlagVisitor.StringVariation("LD-string-flag-go", context, "defaultVal")
	flagValue2, _ := featureFlagVisitor.Float64Variation("LD-Number-flag-go", context, 13.6)

	flagValue, _ := featureFlagVisitor.BoolVariationDetail("LD-Boolean-flag-go-1", context, true)
	flagValue1, _ := featureFlagVisitor.StringVariationDetail("LD-string-flag-go-1", context, "defaultVal1")
	flagValue2, _ := featureFlagVisitor.Float64VariationDetail("LD-Number-flag-go-1", context, 15.6)
}
