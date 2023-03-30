
FeatureFlag::start("your_env_id", "your_api_key");
$visitor = FeatureFlag::newVisitor("your_visitor_id")->build();
$visitor->updateContext("isVip", true);
$visitor->fetchFlags();

$flag = $visitor->getTreatment($key, 'SPLIT-flag-php');

