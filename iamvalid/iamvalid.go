package iamvalid

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
)

type IAMRolePolicy struct {
	PolicyName     string
	PolicyDocument PolicyDocument
}

type PolicyDocument struct {
	Version   string
	Statement []Statement
}

type Statement struct {
	Sid      string
	Effect   string
	Action   []string
	Resource string
}

func IsValid(file string) (returns []bool, error error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %w", err)
	}

	var policy IAMRolePolicy
	err = json.Unmarshal(data, &policy)
	if err != nil {
		return nil, fmt.Errorf("ERROR: %w", err)
	}

	var policyNamePattern = regexp.MustCompile(`(?m)^[\w+=,.@-]+$`)
	if !(policyNamePattern.MatchString(policy.PolicyName) && (len(policy.PolicyName) >= 1 && len(policy.PolicyName) <= 128)) {
		errMsg := "ERROR: JSON does not contain a valid PolicyName"
		return nil, errors.New(errMsg)
	}

	if !(policy.PolicyDocument.Version == "2012-10-17" || policy.PolicyDocument.Version == "2008-10-17") {
		errMsg := fmt.Sprintf("ERROR: Policy %s does not contain a valid Version", policy.PolicyName)
		return nil, errors.New(errMsg)
	}

	if len(policy.PolicyDocument.Statement) == 0 {
		errMsg := fmt.Sprintf("ERROR: Policy %s does not contain a valid Statement", policy.PolicyName)
		return nil, errors.New(errMsg)
	}

	for index, stat := range policy.PolicyDocument.Statement {
		if !(stat.Effect == "Allow" || stat.Effect == "Deny") {
			errMsg := fmt.Sprintf("ERROR: Invalid value for field Effect of Statement %d, policy %s", index, policy.PolicyName)
			return nil, errors.New(errMsg)
		}

		if len(stat.Action) == 0 {
			errMsg := fmt.Sprintf("ERROR: Statement %d of policy %s does not contain any Actions", index, policy.PolicyName)
			return nil, errors.New(errMsg)
		}

		if stat.Resource == "*" {
			returns = append(returns, false)
		} else {
			returns = append(returns, true)
		}
	}

	return returns, nil
}
