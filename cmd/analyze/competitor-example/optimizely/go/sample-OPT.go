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

	flagValue, _ := featureFlagVisitor.GetFeatureVariableBoolean(featureKey, "OPT-flag-go", userID, userAttributes)

}
