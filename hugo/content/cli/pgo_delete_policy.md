---
title: "pgo_delete_policy"
---
## pgo delete policy

Delete a SQL policy

### Synopsis

Delete a policy. For example:

    pgo delete policy mypolicy

```
pgo delete policy [flags]
```

### Options

```
  -h, --help        help for policy
  -n, --no-prompt   No command line confirmation.
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

* [pgo delete](/cli/pgo_delete/)	 - Delete a backup, cluster, pgbouncer, pgpool, label, policy, upgrade, or user

###### Auto generated by spf13/cobra on 8-Mar-2019