# Remitly IAM Validator
Excercise done as part of the 2024 Remitly internship recruitment.

## Excercise
Write a method verifying the input JSON data. Input data format is defined as AWS::IAM::Role Policy. Input JSON might be read from a file. 
Method shall return logical false if an input JSON Resource field contains a single asterisk and true in any other case.

## AWS::IAM::Role Policy
Policy has the following JSON format and is further documented at: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-role-policy.html
```json
{
  "PolicyName": "root",
  "PolicyDocument": {
    "Version": "2012-10-17",
    "Statement": [
      {
        "Sid": "IamListAccess",
        "Effect": "Allow",
        "Action": ["iam:ListRoles", "iam:ListUsers"],
        "Resource": "*"
      }
    ]
  }
}
```
A single Policy can have multiple Statements, therefore the program returns an array of bool (true/false) values corresponding to respective Statements.
The program validates the input considering the validity of its fields as per documentation and returns an error if there is a problem with the input JSON.

## Usage

### Run locally
```bash
git clone https://github.com/Spynacz/remitlyiamvalidator
```

To run the included example
```bash
cd remitlyiamvalidator/main
go run .
```

To run tests
```bash
cd remitlyiamvalidator/iamvalid
go test
```

### Use inside Go project
To use this package inside another Go project
```go
import "github.com/spynacz/remitlyiamvalidator/iamvalid"

res, err := iamvalid.IsValid("example.json")
```

```go
func IsValid(file string) (returns []bool, error error)
```
IsValid returns an array of bools which elements correspond to Statement fields in Policy
