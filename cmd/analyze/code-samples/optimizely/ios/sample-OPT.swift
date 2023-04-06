// Create the visitor
let visitor1 = FeatureFlag.sharedInstance.newVisitor("visitor_1").build()

// Fetch flags
visitor1.fetchFlags {

    let flag = visitor1.getFeatureVariableString(featureKey:"my_feature_key", variableKey:"OPT-flag-swift", userId:"user_123", attributes:attributes)

}