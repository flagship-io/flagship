import { Flagship } from "@flagship.io/js-sdk";

Flagship.start("your_env_id", "your_api_key");

const visitor = Flagship.newVisitor({
    visitorId: "your_visitor_id",
    context: { key: "value" },
});

visitor.on("ready", (error) => {
    if (error) {
        return;
    }

    const btnColorFlag = client.variation("LD-analyze-btnColor", 'red');
    const backgroundSize = client.variation("LD-analyze-backgroundSize", 16);
    const backgroundSize = client.variation("LD-analyze-backgroundSizeFloat", 16.3);
    const showBackground = client.variation("LD-analyze-showBackground", true);

    console.log('btnColorFlag : ', btnColorFlag)
    console.log('backgroundColor : ', backgroundColorFlag)
});

visitor.fetchFlags();


