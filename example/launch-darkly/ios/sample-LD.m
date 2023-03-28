/// Create visitor
FFVisitor * visitor1 = [[[[FeatureFlag sharedInstance] newVisitor:@"visitor_1" instanceType:InstanceSHARED_INSTANCE] withContextWithContext:@{@"age":@18} ] build];

/// Fetch flags
[visitor1 fetchFlagsOnFetchCompleted:^{
  LDFlag * flag = [visitor1 stringVariationForKey:@"LD-string-flag-m" defaultValue:@"red"];
  LDFlag * flag = [visitor1 boolVariationForKey:@"LD-bool-flag-m" defaultValue:TRUE];
  LDFlag * flag = [visitor1 intVariationForKey:@"LD-int-flag-m" defaultValue:13];
  LDFlag * flag = [visitor1 doubleVariationForKey:@"LD-double-flag-m" defaultValue:14.5];
}];
