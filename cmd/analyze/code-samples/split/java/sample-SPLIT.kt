val visitor1 = FeatureFlag.newVisitor("visitor_1")
            .context(hashMapOf("age" to "32", "isVIP" to true))
            .hasConsented(true)
            .isAuthenticated(true)
            .build()

val visitor1 = FeatureFlag.newVisitor("visitor_1").build()
visitor1.updateContext("isVip", true)
visitor1.fetchFlags().invokeOnCompletion {

    val btnColorFlag = visitor1.getTreatment(Key, "SPLIT-flag-kt");
}
