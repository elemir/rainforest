package rainforest

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"strings"
)

var rf *RainForest

func init() {
	rf = New()
}

type RainForest struct {
	envKeyReplacer *strings.Replacer
	pflags         map[string](*pflag.Flag)
}

func New() *RainForest {
	rf := new(RainForest)
	rf.pflags = make(map[string](*pflag.Flag))
	rf.envKeyReplacer = strings.NewReplacer("-", "_")

	return rf
}

// BindPFlags binds a full flag set to the configuration, using each flag's long
// name as the config key.
func BindPFlags(flags *pflag.FlagSet) error { return rf.BindPFlags(flags) }
func (rf *RainForest) BindPFlags(flags *pflag.FlagSet) (err error) {
	flags.VisitAll(func(flag *pflag.Flag) {
		if err = rf.BindPFlag(flag.Name, flag); err != nil {
			return
		}
	})
	return nil
}

// BindPFlag binds a specific key to a pflag (as used by cobra).
// Example (where serverCmd is a Cobra instance):
//
//	 serverCmd.Flags().Int("port", 1138, "Port to run Application server on")
//	 RainForest.BindPFlag("port", serverCmd.Flags().Lookup("port"))
//
func BindPFlag(key string, flag *pflag.Flag) error { return rf.BindPFlag(key, flag) }
func (rf *RainForest) BindPFlag(key string, flag *pflag.Flag) error {
	if flag == nil {
		return fmt.Errorf("flag for %q is nil", key)
	}

	rf.pflags[strings.ToLower(key)] = flag
	return nil
}

// Load is used for getting info from environment and push it to flags
func Load() error { return rf.Load() }
func (rf *RainForest) Load() (err error) {
	for name, flag := range rf.pflags {
		name = strings.ToUpper(name)
		if rf.envKeyReplacer != nil {
			name = rf.envKeyReplacer.Replace(name)
		}
		value, ok := os.LookupEnv(name)
		if ok {
			if err = flag.Value.Set(value); err != nil {
				return
			}
			flag.Changed = true
		}
	}
	return nil
}
