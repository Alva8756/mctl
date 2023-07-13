## mctl completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	mctl completion fish | source

To load completions for every new session, execute once:

	mctl completion fish > ~/.config/fish/completions/mctl.fish

You will need to start a new shell for this setup to take effect.


```
mctl completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.mctl.yml)
      --reauth          re-authenticate with oauth services
```

### SEE ALSO

* [mctl completion](mctl_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 13-Jul-2023