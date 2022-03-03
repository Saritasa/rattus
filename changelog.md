# Changelog

## 0.3.1

* Correct usage of `aws-region`
* Update `aws-sdk-go` to `1.43.10`

## 0.3.0

### Add support for session tokens to aws provider

Added support for temporary keys for AWS secrets provider(`AWS_SESSION_TOKEN`). To use needed aws profile set `AWS_PROFILE`.

Now rattus can be used locally like this:

First auth into aws and get temp keys

```bash
saml2aws login --idp-account={idp_account}
```

Then export them as environment variables

```bash
eval $(saml2aws script --idp-account={idp_account})
```

And just call rattus

```bash
AWS_PROFILE=saml rattus -aws-secret-name={secret_name} -template={template} > {to}
```

In `0.2` version is not possible because we only send `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY` and request to get secrets fails with 400 error.

### 0.2.0

Update read.me

### 0.1.0

Initial release
