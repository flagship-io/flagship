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

var flag = visitor.StringVariation("LD-string-flag-cs", context, 'red');
var flag = visitor.IntVariation("LD-int-flag-cs", context, 13);
var flag = visitor.BoolVariation("LD-bool-flag-cs", context, true);

var flag = visitor.StringVariationDetail("LD-string-flag-cs", context, 'red');
var flag = visitor.IntVariationDetail("LD-int-flag-cs", context, 13);
var flag = visitor.BoolVariationDetail("LD-bool-flag-cs", context, true);