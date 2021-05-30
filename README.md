# config-loader

Config-loader allows you to load configuration values from a file, environment variable and flag whilst using [cobra](https://github.com/spf13/cobra)

## Usage

In your `init()` function you will need to call `config.Init(cmd)` and to register any flags used in the config structure.

You can then define your configuration struct like so with support for both `env` and `flag` tags.
```
type exampleConfig struct {
	Env  string `env:"CONFIG_ENV"`
	Flag string `flag:"CONFIG_FLAG"`
}
```

Within your executor for the cobra command you then simply run `config.Load(cmd, cfg)` with `cfg` being a pointer to the configuration struct.

## Example

You can see an example usage of this library in [example/main.go](example/main.go)