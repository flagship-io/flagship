package http_request

import (
	"github.com/flagship-io/flagship/utils/http_request/common"
	"github.com/flagship-io/flagship/utils/http_request/feature_experimentation"
	"github.com/flagship-io/flagship/utils/http_request/web_experimentation"
)

type HTTPResource interface {
	Init(*common.RequestConfig)
}

var ResourceRequester common.ResourceRequest

var HTTPResources = []HTTPResource{&ResourceRequester}

// feature experimentation
var CampaignRequester feature_experimentation.CampaignRequester = feature_experimentation.CampaignRequester{ResourceRequest: &ResourceRequester}
var AccountEnvironmentFERequester feature_experimentation.AccountEnvironmentFERequester = feature_experimentation.AccountEnvironmentFERequester{ResourceRequest: &ResourceRequester}
var FlagRequester feature_experimentation.FlagRequester = feature_experimentation.FlagRequester{ResourceRequest: &ResourceRequester}
var GoalRequester feature_experimentation.GoalRequester = feature_experimentation.GoalRequester{ResourceRequest: &ResourceRequester}
var ProjectRequester feature_experimentation.ProjectRequester = feature_experimentation.ProjectRequester{ResourceRequest: &ResourceRequester}
var UserRequester feature_experimentation.UserRequester = feature_experimentation.UserRequester{ResourceRequest: &ResourceRequester}
var TargetingKeyRequester feature_experimentation.TargetingKeyRequester = feature_experimentation.TargetingKeyRequester{ResourceRequest: &ResourceRequester}
var VariationGroupRequester feature_experimentation.VariationGroupRequester = feature_experimentation.VariationGroupRequester{ResourceRequest: &ResourceRequester}
var VariationRequester feature_experimentation.VariationRequester = feature_experimentation.VariationRequester{ResourceRequest: &ResourceRequester}
var PanicRequester feature_experimentation.PanicRequester = feature_experimentation.PanicRequester{ResourceRequest: &ResourceRequester}

// web experimentation

var TestRequester web_experimentation.TestRequester = web_experimentation.TestRequester{ResourceRequest: &ResourceRequester}
var GlobalCodeRequester web_experimentation.GlobalRequester = web_experimentation.GlobalRequester{ResourceRequest: &ResourceRequester}
