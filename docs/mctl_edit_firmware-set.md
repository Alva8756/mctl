## mctl edit firmware-set

Edit a firmware set

```
mctl edit firmware-set [flags]
```

### Options

```
  -h, --help                           help for firmware-set
      --labels stringToString          Labels to assign to the firmware set - 'vendor=foo,model=bar' (default [])
      --name string                    Update name for the firmware set
      --remove-firmware-uuids string   UUIDs of firmware to be removed from the set
      --uuid string                    UUID of firmware set to be edited
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.mctl.yml)
      --reauth          re-authenticate with oauth services
```

### SEE ALSO

* [mctl edit](mctl_edit.md)	 - Edit resources

###### Auto generated by spf13/cobra on 13-Jul-2023