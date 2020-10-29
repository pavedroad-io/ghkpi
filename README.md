# ghkpi
Aggregate status for one or more GitHub repositories.
Scope is based on the security of the authenticated user.

You can filter repositories using the -t option and specify a
comma separated list of topics associated with your GitHub
repositories

```bash
ghkpi repo -t one,two,three
```

You can specify a date range with the -r option.  The two options are
**current** or **prior**, for the current or previous month

```bash
ghkpi repo -t one,two,three -r prior
```
