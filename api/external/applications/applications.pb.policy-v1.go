// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: api/external/applications/applications.proto

package applications

import policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"

func init() {
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetServiceGroups", "service_groups", "list", "GET", "/applications/service-groups", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetServiceGroupsHealthCounts", "service_groups", "list", "GET", "/applications/service_groups_health_counts", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetServices", "service_groups", "list", "GET", "/applications/services", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetServicesDistinctValues", "service_groups", "list", "GET", "/applications/services-distinct-values", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*ServicesDistinctValuesReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "field_name":
					return m.FieldName
				case "query_fragment":
					return m.QueryFragment
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetServicesBySG", "service_groups", "list", "GET", "/applications/service-groups/{service_group_id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*ServicesBySGReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "service_group_id":
					return m.ServiceGroupId
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetServicesStats", "service_groups", "list", "GET", "/applications/stats", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetDisconnectedServices", "service_groups", "list", "GET", "/applications/disconnected_services", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/DeleteDisconnectedServices", "service_groups", "delete", "POST", "/applications/delete_disconnected_services", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetVersion", "service_info:version", "read", "GET", "/applications/version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetDisconnectedServicesConfig", "service_groups:scheduler:disconnected_services", "read", "GET", "/retention/service_groups/disconnected_services/config", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/UpdateDisconnectedServicesConfig", "service_groups:scheduler:disconnected_services", "configure", "POST", "/retention/service_groups/disconnected_services/config", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*PeriodicMandatoryJobConfig); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "threshold":
					return m.Threshold
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/GetDeleteDisconnectedServicesConfig", "service_groups:scheduler:delete_disconnected_services", "read", "GET", "/retention/service_groups/delete_disconnected_services/config", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.applications.ApplicationsService/UpdateDeleteDisconnectedServicesConfig", "service_groups:scheduler:delete_disconnected_services", "configure", "POST", "/retention/service_groups/delete_disconnected_services/config", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*PeriodicJobConfig); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "threshold":
					return m.Threshold
				default:
					return ""
				}
			})
		}
		return ""
	})
}
