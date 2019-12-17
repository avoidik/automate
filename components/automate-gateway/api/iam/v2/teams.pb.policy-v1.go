// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/teams.proto

package v2

import (
	policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/ListTeams", "auth:teams", "read", "GET", "/iam/v2/teams", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/GetTeam", "auth:teams:{id}", "read", "GET", "/iam/v2/teams/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetTeamReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/CreateTeam", "auth:teams", "create", "POST", "/iam/v2/teams", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateTeamReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/UpdateTeam", "auth:teams:{id}", "update", "PUT", "/iam/v2/teams/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateTeamReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/DeleteTeam", "auth:teams:{id}", "delete", "DELETE", "/iam/v2/teams/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteTeamReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/GetTeamMembership", "auth:teams:{id}:users", "read", "GET", "/iam/v2/teams/{id}/users", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetTeamMembershipReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/AddTeamMembers", "auth:teams:{id}", "create", "POST", "/iam/v2/teams/{id}/users:add", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.AddTeamMembersReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/RemoveTeamMembers", "auth:teams:{id}", "delete", "POST", "/iam/v2/teams/{id}/users:remove", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.RemoveTeamMembersReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/GetTeamsForMember", "auth:users:{id}:teams", "read", "POST", "/iam/v2/users/{id}/teams", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetTeamsForMemberReq); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
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
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/ApplyV2DataMigrations", "auth:teams", "update", "", "", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.iam.v2.Teams/ResetAllTeamProjects", "auth:teams", "update", "", "", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
}
