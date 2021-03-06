// Code generated by protoc-gen-policy. DO NOT EDIT.
// source: components/automate-gateway/api/auth/tokens/tokens.proto

package tokens

import (
	request "github.com/chef/automate/components/automate-gateway/api/auth/tokens/request"
	policy "github.com/chef/automate/components/automate-gateway/api/authz/policy"
)

func init() {
	policy.MapMethodTo("/chef.automate.api.tokens.TokensMgmt/GetTokens", "auth:api_tokens", "read", "GET", "/auth/tokens", func(unexpandedResource string, input interface{}) string {
		return unexpandedResource
	})
	policy.MapMethodTo("/chef.automate.api.tokens.TokensMgmt/CreateToken", "auth:api_tokens", "create", "POST", "/auth/tokens", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.CreateToken); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "description":
					return m.Description
				case "value":
					return m.Value
				case "id":
					return m.Id
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.tokens.TokensMgmt/UpdateToken", "auth:api_tokens:{id}", "update", "PUT", "/auth/tokens/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.UpdateToken); ok {
			return policy.ExpandParameterizedResource(unexpandedResource, func(want string) string {
				switch want {
				case "id":
					return m.Id
				case "description":
					return m.Description
				default:
					return ""
				}
			})
		}
		return ""
	})
	policy.MapMethodTo("/chef.automate.api.tokens.TokensMgmt/GetToken", "auth:api_tokens:{id}", "read", "GET", "/auth/tokens/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Uuid); ok {
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
	policy.MapMethodTo("/chef.automate.api.tokens.TokensMgmt/DeleteToken", "auth:api_tokens:{id}", "delete", "DELETE", "/auth/tokens/{id}", func(unexpandedResource string, input interface{}) string {
		if m, ok := input.(*request.Uuid); ok {
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
}
