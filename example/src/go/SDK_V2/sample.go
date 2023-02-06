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

	btnColor, _ := fsVisitor.GetModificationString("btnColor", "VString", true)
	btnSize, _ := fsVisitor.GetModificationNumber("btnSize", 13, true)
	showBtn, _ := fsVisitor.GetModificationBool("showBtn", false, true)
}
