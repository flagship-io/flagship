var visitor = Ff.NewVisitor("<VISITOR_ID>")
  .IsAuthenticated(true)
  .HasConsented(true)
  .WithContext(new Dictionary<string, object> {
    ["isVIP"] = true,
    ["country"] = "NL",
    ["loginProvider"] = "Google"
    })
  .Build();

  await visitor.FetchFlags();

var flag = visitors.GetFeatureVariableDouble("my_feature_key", "OPT-flag-cs", "user_123", attributes);
