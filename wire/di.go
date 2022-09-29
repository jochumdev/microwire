package wire

// DiStage1ConfigStore is a marker that the config has been loaded from compiled in opts
type DiStage1ConfigStore struct{}

// DiStage2ConfigStore is a marker that the config has been loaded from different sources (yaml,json,toml,name it here) on top of DiStage1Config
type DiStage2ConfigStore struct{}

// DiStage3ConfigStore is a marker that the config has been loaded from cli on top of DiStage2Config
type DiStage3ConfigStore struct{}
