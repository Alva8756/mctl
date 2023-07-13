## mctl get firmware

Get all firmware components on a server

```
mctl get firmware {-s | --server-id} <server uuid> [flags]
```

### Options

```
  -h, --help               help for firmware
  -s, --server-id string   the server id to look up
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.mctl.yml)
  -o, --output string   {json|text} (default "json")
      --reauth          re-authenticate with oauth services
```

### SEE ALSO

* [mctl get](mctl_get.md)	 - Get resource

###### Auto generated by spf13/cobra on 13-Jul-2023