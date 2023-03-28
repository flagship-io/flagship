// Get the current visitor
  var currentVisitor = FeatureFlag.getCurrentVisitor();

// Fetch flags
    currentVisitor?.fetchFlags().whenComplete(() {
      // Ex: get flag for vip feature
      Flag flag = currentVisitor.boolVariation("LD-bool-flag-dart", false);
      Flag flag = currentVisitor.stringVariation("LD-string-flag-dart", "green");
      Flag flag = currentVisitor.intVariation("LD-int-flag-dart", 16);

      Flag flag = currentVisitor.boolVariationDetail("LD-bool-flag-dart-1", true);
      Flag flag = currentVisitor.stringVariationDetail("LD-string-flag-dart-1", "red");
      Flag flag = currentVisitor.intVariationDetail("LD-int-flag-dart-1", 15);
    });
