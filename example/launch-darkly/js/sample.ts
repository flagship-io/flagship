import { Flagship } from "@flagship.io/js-sdk";

Flagship.start("your_env_id", "your_api_key");

const visitor = Flagship.newVisitor({
    visitorId: "your_visitor_id",
    context: { isVip: true },
});

visitor.on("ready",  (error) => {
    if (error) {
        return;
    }

    const btnColorFlag: string = client.variation("LD-analyze-btnColor", 'red');
    const backgroundSize: number = client.variation("LD-analyze-backgroundSizeFloat", 16.3);
    const showBackground: boolean = client.variation("LD-analyze-showBackground", true);

    console.log('btnColorFlag : ', btnColorFlag);
    console.log('backgroundColorFlag : ', backgroundColorFlag);
});

visitor.fetchFlags();