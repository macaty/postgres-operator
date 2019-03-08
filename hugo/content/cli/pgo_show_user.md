---
title: "pgo_show_user"
---
## pgo show user

Show user information

### Synopsis

Show users on a cluster. For example:

	pgo show user mycluster

```
pgo show user [flags]
```

### Options

```
      --expired string    Shows passwords that will expire in X days.
  -h, --help              help for user
  -s, --selector string   The selector to use for cluster filtering.
```

### Options inherited from parent commands

```
      --apiserver-url string     The URL for the PostgreSQL Operator apiserver.
      --debug                    Enable debugging when true.
      --namespace string         The namespace to use for pgo requests.
      --pgo-ca-cert string       The CA Certificate file path for authenticating to the PostgreSQL Operator apiserver.
      --pgo-client-cert string   The Client Certificate file path for authenticating to the PostgreSQL Operator apiserver.
      --pgo-client-key string    The Client Key file path for authenticating to the PostgreSQL Operator apiserver.
```

### SEE ALSO

* [pgo show](/cli/pgo_show/)	 - Show the description of a cluster

###### Auto generated by spf13/cobra on 8-Mar-2019