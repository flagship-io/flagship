val visitor1 = FeatureFlag.newVisitor("visitor_1")
            .context(hashMapOf("age" to "32", "isVIP" to true))
            .hasConsented(true)
            .isAuthenticated(true)
            .build()

val visitor1 = FeatureFlag.newVisitor("visitor_1").build()
visitor1.updateContext("isVip", true)
visitor1.fetchFlags().invokeOnCompletion {

    val btnColorFlag = visitor1.stringVariation("LD-string-flag-kotlin", "#000000")
    val btnSizeFlag = visitor1.intVariation("LD-number-flag-kotlin", 12)
    val showBackgroundFlag = visitor1.boolVariation("LD-bool-flag-kotlin", false)

    val btnColorFlag = visitor1.stringVariationDetail("LD-string-flag-kotlin-1", "#000000")
    val btnSizeFlag = visitor1.intVariationDetail("LD-number-flag-kotlin-1", 12)
    val showBackgroundFlag = visitor1.boolVariationDetail("LD-bool-flag-kotlin-1", false)
}
