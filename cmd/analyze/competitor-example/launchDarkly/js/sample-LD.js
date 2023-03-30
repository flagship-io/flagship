FeatureFlag.start("your_env_id", "your_api_key");

const visitor = FeatureFlag.newVisitor({
    visitorId: "your_visitor_id",
    context: { key: "value" },
});

visitor.on("ready", (error) => {
    if (error) {
        return;
    }

    const btnColorFlag = client.variation("LD-string-flag-js", 'red');
    const backgroundSize = client.variation("LD-number-flag-js", 16);
    const showBackground = client.variation("LD-bool-flag-js", true);

});

visitor.fetchFlags();


