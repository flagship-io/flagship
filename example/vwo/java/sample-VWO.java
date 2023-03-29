Visitor visitor = featureFlag.newVisitor("visitor_unique_id").build();
visitor.updateContext("isVip", true);
visitor.fetchFlags().whenComplete((instance, error) -> {
    
    Flag<String> btnColorFlag = visitor.getFeatureVariableValue(campaignKey, "VWO-flag-java", "user_123", options);

});

