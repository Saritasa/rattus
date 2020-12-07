# Development

### Example of Visual Code configuration for testing and debugging
```
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "envFile": "${workspaceFolder}/.env",
            "args": [
                "-template=.env.template",
                "-aws-secret-name=aws-secret"
            ]
        },
        {
            "name": "Launch shared creds",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "AWS_REGION":"us-east-1",
                "AWS_PROFILE": "saml"
            },
            "args": [
                "-template=.env.template",
                "-aws-secret-name=aws-secret"
            ]
        },
    ]
}
```

### Deployment

* Add a new tag with version and add it to `changelog.md` and describe what changed
* Then on Github open `Actions` tab pick `Release`, then click `Run workflow` and 
  fill inputs and press `Run workflow`

`GitHub` action file is at `.github/workflows/release.yml`


### Known development issues
* Release file compression works only with `upx 3.94`(so for now will use `ubuntu-18.04` instead of `ubuntu-20.04`(it has `upx 3.95`)). Exception is `freebsd-amd64`, because it not supported.

