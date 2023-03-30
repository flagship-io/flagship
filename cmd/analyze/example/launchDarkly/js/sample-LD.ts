FeatureFlag.start("your_env_id", "your_api_key");

const visitor = FeatureFlag.newVisitor({
    visitorId: "your_visitor_id",
    context: { isVip: true },
});

visitor.on("ready",  (error) => {
    if (error) {
        return;
    }

    const btnColorFlag: string = client.variation("LD-string-flag-ts", 'red');
    const backgroundSize: number = client.variation("LD-number-flag-ts", 16.3);
    const showBackground: boolean = client.variation("LD-bool-flag-ts", true);

    const btnColorFlag: string = client.variationDetail("LD-string-flag-ts-1", 'green');
    const backgroundSize: number = client.variationDetail("LD-number-flag-ts-1", 16);
    const showBackground: boolean = client.variationDetail("LD-bool-flag-ts-1", false);

});

visitor.fetchFlags();