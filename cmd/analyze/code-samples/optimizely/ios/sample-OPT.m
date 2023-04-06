/// Create visitor
FFVisitor * visitor1 = [[[[FeatureFlag sharedInstance] newVisitor:@"visitor_1" instanceType:InstanceSHARED_INSTANCE] withContextWithContext:@{@"age":@18} ] build];

/// Fetch flags
[visitor1 fetchFlagsOnFetchCompleted:^{
    flag = [visitor1 getFeatureVariableString:@"featureKey", variableKey:@"OPT-flag-m", userId:@"user_123", attributes:attributes];
}];
