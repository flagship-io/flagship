FeatureFlag.start("your_env_id", "your_api_key");

const visitor = FeatureFlag.newVisitor({
    visitorId: "your_visitor_id",
    context: { isVip: true },
});

visitor.on("ready",  (error) => {
    if (error) {
        return;
    }

    const btnColorFlag: string = client.getTreatment(key, 'SPLIT-flag-ts');

});

visitor.fetchFlags();