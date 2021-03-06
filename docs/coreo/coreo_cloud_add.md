## coreo cloud add

Add a cloud account

### Synopsis


Add a cloud account. The result would be shown as the following if successful.
         -----------------------------  -----------------------  -----------------------------
               Cloud Account ID           Cloud Account Name               Team ID
         -----------------------------  -----------------------  -----------------------------
             *********************           ************           ************************
         -----------------------------  -----------------------  -----------------------------


```
coreo cloud add [flags]
```

### Examples

```
  coreo cloud add --name YOUR_NEW_ACCOUNT_NAME --role NAME_FOR_NEW_ROLE
  coreo cloud add --name YOUR_NEW_ACCOUNT_NAME --arn YOUR_ROLE_ARN --external-id EXTERNAL_ID_OF_YOUR_ROLE
```

### Options

```
      --arn string                The arn of the role to connect
      --aws-profile string        Aws shared credential file. If empty default provider chain will be used to look for credentials with the following order.
                                    1. Environment variables.
                                    2. Shared credentials file.
                                    3. If your application is running on an Amazon EC2 instance, IAM role for Amazon EC2.
      --aws-profile-path string   The file path of aws profile. If empty will look for AWS_SHARED_CREDENTIALS_FILE env variable. If the env value is empty will default to current user's home directory.
                                    Linux/OSX: "$HOME/.aws/credentials"
                                    Windows:   "%USERPROFILE%\.aws\credentials"
      --external-id string        The external-id used to assume the provided role
  -h, --help                      help for add
  -n, --name string               Name flag
      --policy-arn string         The arn of the policy you'd like to attach for role creation, SecurityAudit policy arn by default (default "arn:aws:iam::aws:policy/SecurityAudit")
      --role string               The name of the role you want to create
```

### Options inherited from parent commands

```
      --api-key string      Coreo API Key (default "None")
      --api-secret string   Coreo API Secret (default "None")
      --endpoint string     Coreo API endpoint. Overrides $CC_API_ENDPOINT. (default "https://app.cloudcoreo.com/api")
      --home string         Location of your Coreo config. Overrides $COREO_HOME. (default "/Users/Jiangz/.cloudcoreo")
      --json                Output in json format
      --profile string      Coreo profile to use. Overrides $COREO_PROFILE. (default "default")
      --team-id string      Coreo team id (default "None")
      --verbose             Enable verbose output
```

### SEE ALSO

* [coreo cloud](coreo_cloud.md)	 - Manage Cloud Accounts

###### Auto generated by spf13/cobra on 15-Nov-2018
