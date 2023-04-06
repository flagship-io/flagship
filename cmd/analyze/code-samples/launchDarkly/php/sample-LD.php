
FeatureFlag::start("your_env_id", "your_api_key");
$visitor = FeatureFlag::newVisitor("your_visitor_id")->build();
$visitor->updateContext("isVip", true);
$visitor->fetchFlags();

$flag = $visitor->variation("LD-bool-flag-php", $context, false);
$flag = $visitor->variation("LD-number-flag-php", $context, 13);
$flag = $visitor->variation("LD-string-flag-php", $context, "red");

$flag = $visitor->variationDetail("LD-bool-flag-php", $context, false);
$flag = $visitor->variationDetail("LD-number-flag-php", $context, 13);
$flag = $visitor->variationDetail("LD-string-flag-php", $context, "red");
