Visitor visitor = featureFlag.newVisitor("visitor_unique_id").build();
visitor.updateContext("isVip", true);
visitor.fetchFlags().whenComplete((instance, error) -> {
    
    Flag<String> btnColorFlag = visitor.getFeatureVariableString("my_feature_key", "OPT-flag-java", "user_123", attributes);

});

