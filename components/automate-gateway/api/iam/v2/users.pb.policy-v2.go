// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/users.proto

package v2

import (
	request "github.com/chef/automate/components/automate-gateway/api/iam/v2/request"
	policyv2 "github.com/chef/automate/components/automate-gateway/authz/policy_v2"
)

func init() {
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Users/CreateUser", "iam:users", "iam:users:create", "POST", "/iam/v2/users", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateUserReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "password":
					return m.Password
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Users/ListUsers", "iam:users", "iam:users:list", "GET", "/iam/v2/users", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Users/GetUser", "iam:users:{id}", "iam:users:get", "GET", "/iam/v2/users/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.GetUserReq); ok {
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
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Users/DeleteUser", "iam:users:{id}", "iam:users:delete", "DELETE", "/iam/v2/users/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.DeleteUserReq); ok {
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
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Users/UpdateUser", "iam:users:{id}", "iam:users:update", "PUT", "/iam/v2/users/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateUserReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "password":
					return m.Password
				default:
					return ""
				}
			})
		}
		return ""
	})
	policyv2.MapMethodTo("/chef.automate.api.iam.v2.Users/UpdateSelf", "iam:usersSelf:{id}", "iam:usersSelf:update", "PUT", "/iam/v2/self/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateSelfReq); ok {
			return policyv2.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "name":
					return m.Name
				case "password":
					return m.Password
				case "previous_password":
					return m.PreviousPassword
				default:
					return ""
				}
			})
		}
		return ""
	})
}
