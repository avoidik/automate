// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/policy.proto

package v2

import (
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
	policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"
)

func init() {
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/CreatePolicy", "iam:policies", "iam:policies:create", "POST", "/iam/v2/policies", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreatePolicyReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/GetPolicy", "iam:policies:{id}", "iam:policies:get", "GET", "/iam/v2/policies/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetPolicyReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/ListPolicies", "iam:policies", "iam:policies:list", "GET", "/iam/v2/policies", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/DeletePolicy", "iam:policies:{id}", "iam:policies:delete", "DELETE", "/iam/v2/policies/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeletePolicyReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/UpdatePolicy", "iam:policies:{id}", "iam:policies:update", "PUT", "/iam/v2/policies/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdatePolicyReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/GetPolicyVersion", "iam:policyVersion", "iam:policies:get", "GET", "/iam/v2/policy_version", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/ListPolicyMembers", "iam:policies:{id}:members", "iam:policyMembers:get", "GET", "/iam/v2/policies/{id}/members", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.ListPolicyMembersReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/ReplacePolicyMembers", "iam:policies:{id}:members", "iam:policyMembers:update", "PUT", "/iam/v2/policies/{id}/members", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.ReplacePolicyMembersReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/RemovePolicyMembers", "iam:policies:{id}:members", "iam:policyMembers:delete", "POST", "/iam/v2/policies/{id}/members:remove", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.RemovePolicyMembersReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/AddPolicyMembers", "iam:policies:{id}:members", "iam:policyMembers:create", "POST", "/iam/v2/policies/{id}/members:add", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.AddPolicyMembersReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/CreateRole", "iam:roles", "iam:roles:create", "POST", "/iam/v2/roles", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateRoleReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/ListRoles", "iam:roles", "iam:roles:list", "GET", "/iam/v2/roles", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/GetRole", "iam:roles:{id}", "iam:roles:get", "GET", "/iam/v2/roles/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetRoleReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/DeleteRole", "iam:roles:{id}", "iam:roles:delete", "DELETE", "/iam/v2/roles/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteRoleReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/UpdateRole", "iam:roles:{id}", "iam:roles:update", "PUT", "/iam/v2/roles/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateRoleReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/CreateProject", "iam:projects", "iam:projects:create", "POST", "/iam/v2/projects", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateProjectReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/UpdateProject", "iam:projects:{id}", "iam:projects:update", "PUT", "/iam/v2/projects/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateProjectReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/GetProject", "iam:projects:{id}", "iam:projects:get", "GET", "/iam/v2/projects/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetProjectReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/ListProjects", "iam:projects", "iam:projects:list", "GET", "/iam/v2/projects", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/DeleteProject", "iam:projects:{id}", "iam:projects:delete", "DELETE", "/iam/v2/projects/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteProjectReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/UpgradeToV2", "system:iam:upgradeToV2", "system:iam:upgrade", "", "", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/ResetToV1", "system:iam:resetToV1", "system:iam:reset", "", "", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Policies/IntrospectAllProjects", "iam:introspect", "iam:introspect:getAllProjects", "GET", "/iam/v2/introspect_projects", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
