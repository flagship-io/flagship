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

vip_feature_enabled = visitor.get_treatment(key, 'SPLIT-flag-py')
