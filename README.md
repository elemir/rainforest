# Overview

RainForest is a simple small library for preload pflag values from environment. It is useful when you use [cobra](https://github.com/spf13/cobra) as cli framework and need to load some arguments from env.
I named it RainForest because cobra lives in rainforests. I agree it isn't very witty

# Why not viper?

[Viper](https://github.com/spf13/viper) is a default solution for such tasks but it has a big problem with design: it uses environment and pflag as independent storages.
For example you have such code from official cobra README:
```go
var author string

func init() {
      rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
      viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
      rootCmd.MarkPersistentFlagRequired("author")
}
```
Content of author variable is agnostic to `AUTHOR` env var. You need to call `viper.Get("author")` for getting right info about author. Secondly it means that required attribute doesn't work correctly with viper:
```
$ AUTHOR="Somebody" program
Error: Required flag(s) "author" have/has not been set
Usage:
  program [flags]

Flags:
      --author string           Author name for copyright attribution
  -h, --help                    help for program

```

# Example

On RainForest code from previous section will look like:

```go
var author string

func init() {
      rootCmd.PersistentFlags().StringVar(&author, "author", "YOUR NAME", "Author name for copyright attribution")
      rainforest.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
      rootCmd.MarkPersistentFlagRequired("author")
      rainforest.Load()
}

```
*Be carefull*: Load() should be run after all plfags preparation will be done

# Contributing
Just create pull request on github

# License
RainForest is released under the LPGLv3 LICENSE. See [COPYING](https://github.com/elemir/rainforest/blob/master/COPYING)
