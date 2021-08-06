package vmware

import "github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
	return "VMware"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
	return []string{
		"VMware (AVS)",
	}
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azurerm_vmware_private_cloud": dataSourceVmwarePrivateCloud(),
	}
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*pluginsdk.Resource {
	return map[string]*pluginsdk.Resource{
		"azurerm_vmware_private_cloud":               resourceVmwarePrivateCloud(),
		"azurerm_vmware_cluster":                     resourceVmwareCluster(),
		"azurerm_vmware_express_route_authorization": resourceVmwareExpressRouteAuthorization(),
	}
}
