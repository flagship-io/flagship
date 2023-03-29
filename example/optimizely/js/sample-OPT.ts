FeatureFlag.start("your_env_id", "your_api_key");

const visitor = FeatureFlag.newVisitor({
    visitorId: "your_visitor_id",
    context: { isVip: true },
});

visitor.on("ready",  (error) => {
    if (error) {
        return;
    }

    const btnColorFlag: string = client.getFeatureVariable('my_feature_key', 'OPT-flag-ts', 'user_123', attributes);

});

visitor.fetchFlags();