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

var flag = visitor.GetTreatment(key, "SPLIT-flag-cs");
