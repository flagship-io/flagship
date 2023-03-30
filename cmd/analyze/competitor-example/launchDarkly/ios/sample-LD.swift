// Create the visitor
let visitor1 = FeatureFlag.sharedInstance.newVisitor("visitor_1").build()

// Fetch flags
visitor1.fetchFlags {

    let flag = visitor1.stringVariation(forKey: "LD-string-flag-swift", defaultValue: "red")
    let flag = visitor1.boolVariation(forKey: "LD-bool-flag-swift", defaultValue: false)
    let flag = visitor1.intVariation(forKey: "LD-int-flag-swift", defaultValue: 13)
    let flag = visitor1.doubleVariation(forKey: "LD-double-flag-swift", defaultValue: 13.5)

    let flag = visitor1.stringVariationDetail(forKey: "LD-string-flag-swift-1", defaultValue: "red")
    let flag = visitor1.boolVariationDetail(forKey: "LD-bool-flag-swift-1", defaultValue: false)
    let flag = visitor1.intVariationDetail(forKey: "LD-int-flag-swift-1", defaultValue: 13)
    let flag = visitor1.doubleVariationDetail(forKey: "LD-double-flag-swift-1", defaultValue: 13.5)

}