# Purpose:
An example of how to create a performant serverless project written in Golang with AWS SAM.

# Tutorial:
1. Install and configure the AWS SAM CLI with the AWS CLI and have access keys setup: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/manage-sam-cli-versions.html
2. Run `sam sync --stack-name go-serverless-test --watch --region us-east-1` to spawn this stack on a development stack. Test the endpoint against the generated HTTPS url.
3. Delete the stack with `sam delete --stack-name go-severless-test --region us-east-1` to avoid billing.
