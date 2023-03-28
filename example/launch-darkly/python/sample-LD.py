class CustomEventHandler(FeatureFlagEventHandler):
    def __init__(self):
        FeatureFlagEventHandler.__init__(self)

    def on_log(self, level, message):
        print("Log >> " + message)
        pass

    def on_exception_raised(self, exception, traceback):
        FeatureFlagEventHandler.on_exception_raised(self, exception, traceback)
        print("Exception >> " + str(exception))
        pass


 FeatureFlag.instance().start("your_env_id", "your_api_key"
        Config(event_handler=CustomEventHandler())

FeatureFlag.instance().start("your_env_id", "your_api_key", Config())

visitor = FeatureFlag.instance().create_visitor("user_#1234", {'isVip':True})
visitor.synchronize_modifications()

vip_feature_enabled = visitor.variation('LD-bool-flag-py', context, False)
vip_feature_color = visitor.variation('LD-string-flag-py', context, "red")
vip_feature_size = visitor.variation('LD-number-flag-py', context, 16)
