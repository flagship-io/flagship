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
var VariationFERequester feature_experimentation.VariationFERequester = feature_experimentation.VariationFERequester{ResourceRequest: &ResourceRequester}
var PanicRequester feature_experimentation.PanicRequester = feature_experimentation.PanicRequester{ResourceRequest: &ResourceRequester}

// web experimentation

var TestRequester web_experimentation.TestRequester = web_experimentation.TestRequester{ResourceRequest: &ResourceRequester}
var GlobalCodeRequester web_experimentation.GlobalCodeRequester = web_experimentation.GlobalCodeRequester{ResourceRequest: &ResourceRequester}
var AccountWERequester web_experimentation.AccountWERequester = web_experimentation.AccountWERequester{ResourceRequest: &ResourceRequester}
var VariationWERequester web_experimentation.VariationWERequester = web_experimentation.VariationWERequester{ResourceRequest: &ResourceRequester}
