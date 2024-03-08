Visitor visitor = featureFlag.newVisitor("visitor_unique_id").build();
visitor.updateContext("isVip", true);
visitor.fetchFlags().whenComplete((instance, error) -> {
    
    Flag<String> btnColorFlag = visitor.stringVariation("LD-string-flag-java", context, "red");
    Flag<Number> backgroundColorFlag = visitor.intVariation("LD-number-flag-java", context, 13);
    Flag<Boolean> backgroundColorFlag = visitor.boolVariation("LD-bool-flag-java", context, true);

    Flag<String> btnColorFlag = visitor.stringVariationDetail("LD-string-flag-java-1", context, "red");
    Flag<Number> backgroundColorFlag = visitor.intVariationDetail("LD-number-flag-java-1", context, 13);
    Flag<Boolean> backgroundColorFlag = visitor.boolVariationDetail("LD-bool-flag-java-1", context, true);

});

