+++
title = "IAM v2 API Reference"
description = "IAM v2 API Reference"
draft = false
bref = ""
toc = true
[menu]
  [menu.docs]
    parent = "authorization"
    weight = 50
+++

This API reference details Chef Automate IAM v2 features from the command line.
If you are not already on IAM v2, [upgrade to IAM v2]({{< relref "iam-v2-guide.md#upgrade-to-iam-v2" >}}) first to ensure expected outcomes when using `curl` commands.

## Getting Started

Keep in mind the following context when exploring this API reference:

1. If you do not have one already, generate an admin token for IAM v2.
   To do so, run this command, filling in a token name of your choice:

   ```bash
   export TOKEN=`chef-automate iam token create <your-token-name-here> --admin`
   ```

2. URIs are relative to `https://<your-domain-here>/apis/iam/v2beta`, unless otherwise noted.

3. Attach the `?pretty` query string to an endpoint to get pretty-printed output.

Putting all of the above information together with the `/policies` endpoint as an example, this command fetches the list of policies with multi-line, formatted output:

```bash
curl -sH "api-token: $TOKEN" \
  https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/policies?pretty
```

For those API methods that take JSON data, typically through the `create` and `update` methods,
you can provide that data to the REST endpoint in a variety of ways.
If the data is a relatively small size payload, you can include it in a `curl` command inline. For example:

```bash
curl -sH "api-token: $TOKEN" -d '<your JSON here>' ...
```

If the payload is larger, it is often convenient to store the data in a file, then pass that file. For example: 

```bash
curl -sH "api-token: $TOKEN" -d @policy.json ...
```

## Projects Property

Teams, tokens, roles, and policies all contain a top-level `projects` property. You will see that in the example JSON for each resource in the following sections.
This `projects` property associates the particular IAM resource with one or more projects.
Users must have permissions for those projects to be able to view or modify those resources.
If the `projects` property is left empty, that resource is categorized as *unassigned*.

To create those permissions for users, there is a separate and distinct `projects` property on each policy statement.
This `projects` property may **not** be empty. You must either specify specific projects, or use a wildcard (`*`) to indicate all projects.

Type               | May be empty ?    | Allows "*" ?         | May specify (unassigned) ?
-------------------|-------------------|----------------------|---------------------------
top-level projects | empty allowed     | wildcard not allowed | (unassigned) not allowed
statement projects | empty not allowed | wildcard allowed     | (unassigned) allowed

## Policies

HTTP request              | Description
--------------------------|------------
GET /policies             | list both *Chef-managed* and *Custom* policies
GET /policies/{**id**}    | get the specified policy
POST /policies            | create a *Custom* policy
PUT /policies/{**id**}    | update a *Custom* policy
DELETE /policies/{**id**} | delete a *Custom* policy

### Example Policy

This policy has two statements, each enforcing a different set of permissions.
The first policy statement grants "list" and "get" access to view users.
The second statement grants access to all the actions comprising the `editor` role.
Note that you may specify inline **actions** or use a **role** reference.

```json
{
  "name": "Alpha Beta Policy",
  "id": "alpha-beta-policy-1",
  "members": ["team:local:alpha","team:ldap:beta"],
  "statements": [
    {
      "effect": "ALLOW",
      "actions": [
        "iam:users:list",
        "iam:users:get"
      ],
      "projects": [
        "*"
      ]
    },
    {
      "effect": "ALLOW",
      "role": "editor",
      "projects": [
        "*"
      ]
    }
  ],
  "projects": [
    "admin-group-1",
    "it-support",
  ]
}
```

### Listing Policies

The output for this command includes both *Chef-managed* and *Custom* policies.

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/policies?pretty
```

### Getting a Policy

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/policies/{policy-id}?pretty
```

### Creating a Policy

Create a policy by composing JSON with all the necessary policy properties (see [Example Policy]({{< relref "iam-v2-api-reference.md#example-policy" >}})).

Assuming you store the JSON content in the file "policy.json", pass the file to `curl` to create the policy:

```bash
curl -sSH "api-token: $TOKEN" -d @policy.json -X POST \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/policies?pretty
```

### Updating a Policy

When updating a policy, supply all of a policy's properties, not just the ones you wish to update. Properties that you do not include are reset to empty values.
You can modify the *membership* of both Chef-managed and custom policies.
You can modify the *definition* of custom policies, but not Chef-managed policies.
The policy ID is immutable and it can only be set at creation time.

```bash
curl -sSH "api-token: $TOKEN" -d @policy.json -X PUT \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/policies/{policy-id}?pretty
```

### Deleting a Policy

Deleting a policy is permanent and cannot be undone.

```bash
curl -sSH "api-token: $TOKEN" -X DELETE \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/policies/{policy-id}?pretty
```

## Policy Membership

Chef Automate supports two methods for updating policy membership:

1. Using `PUT` lets you set membership for the entire policy. For more information, see [Updating All Members On a Policy]({{< relref "iam-v2-api-reference/#updating-all-members-on-a-policy" >}}).

2. Using `POST` lets you add members to, or remove members from, an existing membership list.
For more information, see [Adding Members to a Policy]({{< relref "iam-v2-api-reference/#adding-members-to-a-policy" >}}) and [Removing Members from a Policy]({{< relref "iam-v2-api-reference/#removing-members-from-a-policy" >}}).

HTTP request                           | Description
---------------------------------------|------------
GET /policies/{**id**}/members         | list the membership
PUT /policies/{**id**}/members         | replace the membership
POST /policies/{**id**}/members:add    | add to the membership
POST /policies/{**id**}/members:remove | remove from the membership

### Example Policy Membership Payload

```json
{
  "members": [
    "team:local:admins",
    "token:admin-token"
  ]
}
```

### Listing Policy Members

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/policies/{policy-id}/members?pretty
```

### Updating All Members on a Policy

Note the `member-id` must be in the correct [Member Expression]({{< relref "iam-v2-guide.md#member-expressions" >}}) format.

Use this HTTP request to replace the policy's membership:

```bash
curl -sSH "api-token: $TOKEN" -X PUT \
-d '{"members":["{member-id}","{member-id}","{member-id}"]}' \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/policies/{policy-id}/members?pretty
```

### Adding Members to a Policy

Specify new members to add to a policy's membership via this HTTP request.
Ensure each `member-id` is in the correct [Member Expression]({{< relref "iam-v2-guide.md#member-expressions" >}}) format.

```bash
curl -sSH "api-token: $TOKEN" -X POST \
-d '{"members":["{member-id}","{member-id}"]}' \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/policies/{policy-id}/members:add?pretty
```

### Removing Members from a Policy

Specify members to remove from a policy's membership via this HTTP request. Ensure each `member-id` is in the correct [Member Expression]({{< relref "iam-v2-guide.md#member-expressions" >}}) format.

The removed members will still exist within Chef Automate, but are no longer associated with this policy.

```bash
curl -sSH "api-token: $TOKEN" -X POST \
-d '{"members":["{member-id}","{member-id}","{member-id}","{member-id}"]}' \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/policies/{policy-id}/members:remove?pretty
```

## Roles

A role encapsulates a list of actions.
IAM v2 provides several default *Chef-managed* roles, which are essential to the operation of Chef Automate and cannot be altered.
Roles that you create are *Custom* roles.

HTTP request              | Description
--------------------------|------------
GET /roles                | list both *Chef-managed* and *Custom* roles
GET /roles/{**id**}       | get the specified role
POST /roles               | create a *Custom* role
PUT /roles/{**id**}       | update a *Custom* role
DELETE /roles/{**id**}    | delete a *Custom* role

### Default Chef-managed Roles

Name | ID| Actions
-----------------------|-----|--------
Owner              | owner         | \*
Viewer             | viewer        | secrets:\*:get, secrets:\*:list, infra:\*:get, infra:\*:list, compliance:\*:get, compliance:\*:list, system:\*:get, system:\*:list, event:\*:get, event:\*:list, ingest:\*:get, ingest:\*:list, iam:projects:list, iam:projects:get, applications:\*:list, applications:\*:get
Editor             | editor        | infra:\*, compliance:\*, system:\*, event:\*, ingest:\*, secrets:\*, telemetry:\*, iam:projects:list, iam:projects:get, iam:projects:assign, applications:\*
Project Owner      | project-owner | infra:\*, compliance:\*, system:\*, event:\*, ingest:\*, secrets:\*, telemetry:\*, iam:projects:list, iam:projects:get, iam:projects:assign, iam:policies:list, iam:policies:get, iam:policyMembers:\*, iam:teams:list, iam:teams:get, iam:teamUsers:\*, iam:users:get, iam:users:list
Ingest             | ingest        | infra:ingest:\*, compliance:profiles:get, compliance:profiles:list

### Example Role

```json
{
  "name": "Advocate",
  "id": "advocate-role",
  "actions": [
    "infra:*",
    "compliance:*",
    "teams:*",
    "users:*"
  ],
  "projects": [
    "east-region",
    "west-region"
  ]
}
```

### Listing Roles

The output from this `curl` command includes both *Chef-managed* and *Custom* roles.

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/roles?pretty
```

### Getting a Role

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/roles/{role-id}?pretty
```

### Creating a Role

Create a role by composing JSON with all the necessary role properties (see [Example Role]({{< relref "iam-v2-api-reference.md#example-role" >}})).

Assuming you store the JSON content in the file "role.json", pass the file to `curl` to create the role:

```bash
curl -sSH "api-token: $TOKEN" -d @role.json -X POST \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/roles?pretty
```

### Updating a Role

When updating a role, supply all of a role's properties, not just the ones you wish to update. Properties that you do not include are reset to empty values.

The role ID is immutable and it can only be set at creation time.

```bash
curl -sSH "api-token: $TOKEN" -d @role.json -X PUT \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/roles/{role-id}?pretty
```

### Deleting a Role

Deleting a role is permanent and cannot be undone.

```bash
curl -sSH "api-token: $TOKEN" -X DELETE \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/roles/{role-id}?pretty
```

## Projects

HTTP request              | Description
--------------------------|------------
GET /projects             | list all projects
GET /projects/{**id**}    | get the specified role
POST /projects            | create a project
PUT /projects/{**id**}    | update a project
DELETE /projects/{**id**} | delete a project

### Example Project

```json
{
  "name": "eastern region",
  "id": "eastern-region",
  "status": "NO_RULES",
  "type": "CUSTOM"
}
```

### Listing Projects

```bash
curl -sH "api-token: $TOKEN" -X GET \
  https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/projects?pretty
```

### Getting a Project

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/projects/{project-id}?pretty
```

### Creating a Project

Create a project by composing JSON with all the necessary project properties (see [Example Project]({{< relref "iam-v2-api-reference.md#example-project" >}})).

```bash
curl -sSH "api-token: $TOKEN" -X POST \
-d '{"name": "Test Project", "id": "test-project-1" }' \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/projects?pretty
```

### Updating a Project

Remember when updating a project, the project name is the only property of a project you can modify. The ID is immutable and it can only be set at creation time.

```bash
curl -sSH "api-token: $TOKEN" -X PUT \
-d '{"name": "Test Role", "actions": ["compliance:*", "teams:*", "users:*"] }' \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/projects/{[project]-id}?pretty
```

### Deleting a Project

Deleting a project is permanent and cannot be undone.

```bash
curl -sSH "api-token: $TOKEN" -X DELETE \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/projects/{project-id}?pretty
```

## Project Rules

HTTP request                                          | Description
------------------------------------------------------|------------
GET /projects/***project_id***/rules                  | list rules for a project
GET /projects/***project_id***/rules/***rule_id***    | get a project rule
POST /projects/***project_id***/rules                 | create a project rule
PUT /projects/***project_id***/rules/***rule_id***    | update a project rule
DELETE /projects/***project_id***/rules/***rule_id*** | delete a project rule

A project rule specifies a non-empty array of conditions.
Each project rule **condition** specifies values to match for a particular **attribute**.
The choice of attributes depends on the rule **type**.

Rule Type  | Available Attributes
-----------|---------------------
EVENT      | CHEF_ORGANIZATION, CHEF_SERVER
NODE       | CHEF_ORGANIZATION, CHEF_SERVER, ENVIRONMENT, CHEF_ROLE, CHEF_TAG, CHEF_POLICY_NAME, CHEF_POLICY_GROUP

Each rule condition also specifies an **operator**, either `MEMBER_OF` or `EQUALS`.
To match a single value, use `EQUALS` and include just a single value in the **values** array.
To match any one of a set of values, use `MEMBER_OF` and include the choices in the **values** array.

All rule conditions must be true for a rule to apply.
If, for example, you want to match when `chef_organization` is either `project1` or `project2`, this will **not** work:

```json
{
  "operator": "EQUALS", "attribute": "CHEF_SERVER", "values": [ "project1" ],
  "operator": "EQUALS", "attribute": "CHEF_SERVER", "values": [ "project2" ]
},
```

The conjunction of those conditions will always be false because both can never be true at the same time.
Rather, write the rule with a single condition like this:

```json
{
  "operator": "MEMBER_OF",
  "attribute": "CHEF_SERVER",
  "values": [ "project1", "project2" ]
},
```

### Example Project Rule

```json
{
  "id": "org1-filter",
  "name": "server and org filter",
  "type": "NODE",
  "project_id": "eastern-region",
  "conditions": [
    {
      "operator": "MEMBER_OF",
      "attribute": "CHEF_SERVER",
      "values": [
        "prod",
        "staging"
      ]
    },
    {
      "operator": "EQUALS",
      "attribute": "CHEF_ORGANIZATION",
      "values": [
        "org1"
      ]
    }
  ]
}
```

### Listing Rules for a Project

The output lists all rules for a specified project.

```bash
curl -sH "api-token: $TOKEN" -X GET \
  https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/projects/{project-id}/rules?pretty
```

### Getting a Project Rule

This shows the details of a single project, selected by its ID.

```bash
curl -sH "api-token: $TOKEN" -X GET \
  https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/projects/{project-id}/rules/{rule-id}?pretty
```

### Creating a Project Rule

Create a rule by composing JSON with all the necessary rule properties (see [Example Project Rule]({{< relref "iam-v2-api-reference.md#example-project-rule" >}})).

Assuming you store the JSON content in the file "project-rule.json", pass the file to `curl` to create the project rule:

```bash
curl -sSH "api-token: $TOKEN" -d @project-rule.json -X POST \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/projects/{project-id}/rules?pretty
```

### Updating a Project Rule

When updating a Project Rule, supply all of a rule's properties, not just the ones you wish to update. If you do not, properties that you do not include are reset to empty values. Know that the `Rule ID` and `Type` are immutable and both properties can only be set at creation time. Note that if you need to change the rule's `Type`, then you will need to create a new rule instead.

```bash
curl -sSH "api-token: $TOKEN" -d @project-rule.json -X PUT \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/projects/{project-id}/rules/{rule-id}?pretty
```

### Deleting a Project Rule

Deleting a project rule is permanent and cannot be undone.

```bash
curl -sSH "api-token: $TOKEN" -X DELETE \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/projects/{project-id}/rules/{rule-id}?pretty
```

## Applying Rules

HTTP request                                          | Description
------------------------------------------------------|------------
GET    /apply-rules                                   | get status of a rule application
POST   /apply-rules                                   | start applying pending rule or project edits
DELETE /apply-rules                                   | stop applying pending rule or project edits

**NOTE**: With a large amount of historical compliance data, rule application
can take a considerable amount of time. If making a large number of rule
changes, it's best to batch them up and apply them all at once.

## Users

HTTP request              | Description
--------------------------|------------
GET /users                | list all users
GET /users/{**id**}       | get the specified user
POST /users               | create a user
PUT /users/{**id**}       | update a user
DELETE /users/{**id**}    | delete a user
GET /users/{**id**}/teams | list teams for a user
PUT /self/{**id**}        | update your own user

### Example User

```json
{
  "name": "Douglas Adams",
  "id": "doug42",
  "password": "secret_pwd"
}
```

### Listing Users

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/users?pretty
```

### Getting a User

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/users/{user-id}?pretty
```

### Creating a User

Create a user by composing JSON with all the necessary user properties (see [Example User]({{< relref "iam-v2-api-reference.md#example-user" >}})).

Assuming you store the JSON content in the file "user.json", pass the file to `curl` to create the user:

```bash
curl -sSH "api-token: $TOKEN" -d @user.json -X POST \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/users?pretty
```

### Updating a User

When updating a user, supply all of a user's properties, not just the ones you wish to update. 
Properties that you do not include are reset to empty values.
The user ID is immutable and it can only be set at creation time.

```bash
curl -sSH "api-token: $TOKEN" -d @user.json -X PUT \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/users/{user-id}?pretty
```

### Deleting a User

Deleting a user is permanent and cannot be undone.

```bash
curl -sSH "api-token: $TOKEN" -X DELETE \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/users/{user-id}?pretty
```

### List Teams for a User

The output for the command below lists all teams that contain the specified user.
Note, the output uses the `membership_id` property of the user instead of the user's `id`.

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/users/{membership-id}/teams?pretty
```

### Updating Your Own User

When updating your own user, supply all of your user properties, not just the ones you wish to update.
Properties that you do not include are reset to empty values.
Your user ID is immutable and it can only be set at creation time.

```bash
curl -sSH "api-token: $TOKEN" -d @user.json -X PUT \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/self/{user-id}?pretty
```

## Teams

HTTP request              | Description
--------------------------|------------
GET /teams               | list all teams
GET /teams/{**id**}      | get the specified team
POST /teams              | create a team
PUT /teams/{**id**}      | update a team
DELETE /teams/{**id**}   | delete a team

### Example Team

```json
{
  "name": "team 1",
  "id": "team-1",
  "projects": [
    "east-region",
    "west-region"
  ]
}
```

### Listing Teams

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/teams?pretty
```

### Getting a Team

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/teams/{team-id}?pretty
```

### Creating a Team

Create a team by composing JSON with all the necessary team properties (see [Example Team]({{< relref "iam-v2-api-reference.md#example-team" >}})).

Assuming you store the JSON content in the file "team.json", pass the file to `curl` to create the team:

```bash
curl -sSH "api-token: $TOKEN" -d @team.json -X POST \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/teams?pretty
```

### Updating a Team

When updating a team, supply all of a team's properties, not just the ones you wish to update.
Properties that you do not include are reset to empty values.
The team ID is immutable and it can only be set at creation time.

```bash
curl -sSH "api-token: $TOKEN" -d @team.json -X PUT \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/teams/{team-id}?pretty
```

### Deleting a Team

Deleting a team is permanent and cannot be undone.

```bash
curl -sSH "api-token: $TOKEN" -X DELETE \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/teams/{team-id}?pretty
```

## Team Membership

A local team consists exclusively of local users.
Note that the `user_ids` of the local users used here are GUIDs, generated when you create users.
A user's GUID is reported in the `membership_id` property when you fetch a user.

HTTP request                           | Description
---------------------------------------|------------
GET /teams/{**id**}/users              | list the membership
POST /teams/{**id**}/users:add         | add to the membership
POST /teams/{**id**}/users:remove      | remove from the membership

### Example Team Membership Payload

```json
{
  "id": "team-1",
  "user_ids": [
    "1a5a63d7-0b14-465a-aec1-5eea08a72343",
    "6c4e3e8a-fade-4a9c-ac9f-99b86677530b"
  ]
}
```

### Listing Team Users

The output lists all local users on the specified team.

```bash
curl -sSH "api-token: $TOKEN"  -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/teams/{team-id}/users?pretty
```

### Adding Users to a Team

```bash
curl -sSH "api-token: $TOKEN" -X POST \
-d '{"user_ids":["{membership_id}","{membership_id}"]}' \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/teams/{team-id}/users:add?pretty
```

### Removing Users from a Team

The removed users still exists within Chef Automate, but are no longer associated with this team.

```bash
curl -sSH "api-token: $TOKEN" -X POST \
-d '{"user_ids":["{membership_id}","{membership_id}"]}' \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/teams/{team-id}/users:remove?pretty
```

## Tokens

HTTP request              | Description
--------------------------|------------
GET /tokens               | list all tokens
GET /tokens/{**id**}      | get the specified token
POST /tokens              | create a token
PUT /tokens/{**id**}      | update a token
DELETE /tokens/{**id**}   | delete a token

### Example Token

```json
{
  "name": "token 1",
  "id": "token-1",
  "active": true,
  "projects": [
    "east-region",
    "west-region"
  ]
}
```

### Listing Tokens

This command lists admin and non-admin tokens.

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/tokens?pretty
```

### Getting a Token

```bash
curl -sSH "api-token: $TOKEN" -X GET \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/tokens/{token-id}?pretty
```

### Creating a Token

Create a token by composing JSON with all the necessary token properties (see [Tokens]({{< relref "iam-v2-api-reference.md#tokens" >}})).

Assuming you store the JSON content in the file "token.json", pass the file to `curl` to create the token:

```bash
curl -sSH "api-token: $TOKEN" -d @token.json -X POST \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/tokens?pretty
```

Note that this creates **non**-admin tokens that may then be assigned permissions via policies just like users or teams (unless you have already created policies that encompass all tokens using `tokens:*`).

You cannot create admin tokens via the REST API.
Admin tokens can only be created by specifying the `--admin` flag to this `chef-automate` sub-command:

```bash
  chef-automate iam token create <your-token-name> --admin
```

### Updating a Token

When updating a token, supply all of a token's properties, not just the ones you wish to update.
Properties that you do not include are reset to empty values.
The token ID is immutable and it can only be set at creation time.

```bash
curl -sSH "api-token: $TOKEN" -d @token.json -X PUT \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/tokens/{token-id}?pretty
```

### Deleting a Token

Deleting a token is permanent and cannot be undone.

```bash
curl -sSH "api-token: $TOKEN" -X DELETE \
https://{{< example_fqdn "automate" >}}/apis/iam/v2beta/tokens/{token-id}?pretty
```
