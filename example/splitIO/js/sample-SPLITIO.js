FeatureFlag.start("your_env_id", "your_api_key");

const visitor = FeatureFlag.newVisitor({
    visitorId: "your_visitor_id",
    context: { key: "value" },
});

visitor.on("ready", (error) => {
    if (error) {
        return;
    }

    const btnColorFlag = client.getTreatment(key, 'SPLITIO-flag-js');

});

visitor.fetchFlags();


