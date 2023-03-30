Visitor visitor = featureFlag.newVisitor("visitor_unique_id").build();
visitor.updateContext("isVip", true);
visitor.fetchFlags().whenComplete((instance, error) -> {
    
    Flag<String> btnColorFlag = visitor.getTreatment(key, "SPLIT-flag-java");

});

