package actionsgithubcom

const (
	LabelKeyRunnerTemplateHash = "runner-template-hash"
	LabelKeyPodTemplateHash    = "pod-template-hash"
)

const (
	EnvVarRunnerJITConfig      = "ACTIONS_RUNNER_INPUT_JITCONFIG"
	EnvVarRunnerExtraUserAgent = "GITHUB_ACTIONS_RUNNER_EXTRA_USER_AGENT"
)

// Environment variable names used to set proxy variables for containers
const (
	EnvVarHTTPProxy  = "http_proxy"
	EnvVarHTTPSProxy = "https_proxy"
	EnvVarNoProxy    = "no_proxy"
)
