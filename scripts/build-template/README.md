# build-template
- This is a build template tool for AWS CloudFormation templates. The tool was built to put all resources together in the same template.
- Using the tool we **avoid** to use the [Outputs](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/outputs-section-structure.html) of each resource and the [nested stacks](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-nested-stacks.html) of AWS CloudFormation.
- Folder structure sample:
  ```
  root folder
    |_resources
        |_api.yaml
        |_cognito-user-pool-client.yaml
        |_cognito.user-pool.yaml
    |_template.yaml
  ```
- After run the tool the folder structure should be in this state:
  ```
  root folder
    |_resources
        |_api.yaml
        |_cognito-user-pool-client.yaml
        |_cognito.user-pool.yaml
    |_template.yaml
    |_final_template.yaml
  ```
- Sample of `template.yaml` file:
  ```yaml
  Transform: AWS::Serverless-2016-10-31
  Globals:
    Function:
      Runtime: go.x
      Timeout: 3
      MemorySize: 128

  Parameters:
    Stage:
      Type: String
      Default: dev
      AllowedValues:
        - dev
        - hml
        - prod
      Description: "Enter the stage name: dev, hml, or prod. Default is dev."

  Resources:

  ```
- All content of the resources will be add at the end of the `template.yaml` file after the `Resources` section.
- The resources should be written like this:
  ```yaml
  Resources:
    CognitoUserPoolClient:
      Type: AWS::Cognito::UserPoolClient
      Properties:
        UserPoolId: !Ref CognitoUserPool
        ClientName: !Join [ "", [ !Ref Stage, _tutor_user_pool_client ] ]
        GenerateSecret: false
        ExplicitAuthFlows:
          - ALLOW_USER_PASSWORD_AUTH

  ```
- The `final_template.yaml` file:
  ```yaml
  Transform: AWS::Serverless-2016-10-31
  Globals:
    Function:
      Runtime: go.x
      Timeout: 3
      MemorySize: 128

  Parameters:
    Stage:
      Type: String
      Default: dev
      AllowedValues:
        - dev
        - hml
        - prod
      Description: "Enter the stage name: dev, hml, or prod. Default is dev."

  Resources:
    CognitoUserPoolClient:
      Type: AWS::Cognito::UserPoolClient
      Properties:
        UserPoolId: !Ref CognitoUserPool
        ClientName: !Join [ "", [ !Ref Stage, _tutor_user_pool_client ] ]
        GenerateSecret: false
        ExplicitAuthFlows:
          - ALLOW_USER_PASSWORD_AUTH
  ```
