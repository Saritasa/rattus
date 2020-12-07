# Changelog

### 0.3
**Add support for session tokens to aws provider**

Added support for temporary keys for AWS secrets provider(`AWS_SESSION_TOKEN`). To use needed aws profile set `AWS_PROFILE`.

Now rattus can be used locally like this 
This auth into was and get temp keys
```
saml2aws login --idp-account={idp_account}
```
Then export them as environment variables
```
eval $(saml2aws script --idp-account={idp_account})
```
And just call rattus
```
AWS_PROFILE=saml rattus -aws-secret-name={secret_name} -template={template} > {to}
```
In 0.2 version is not possible because we only send `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` and request to get secrets fails with 400 error.

### 0.2
**Update read.me**

### 0.1
**Initial release**
