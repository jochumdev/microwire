package di

// DiFlags is a marker that the config has been loaded from compiled in opts
type DiFlags struct{}

// DiConfigData is a marker that the config has been loaded from different sources (yaml,json,toml,name it here) on top of DiStage1Config
type DiConfigData struct{}

type DiConfig string

type DiConfigor struct{}
