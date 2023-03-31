
FeatureFlag::start("your_env_id", "your_api_key");
$visitor = FeatureFlag::newVisitor("your_visitor_id")->build();
$visitor->updateContext("isVip", true);
$visitor->fetchFlags();

$flag = $visitor->getFeatureVariableDouble('my_feature_key', 'OPT-flag-php', 'user_123', $attributes);

