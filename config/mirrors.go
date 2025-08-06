package config

const (
	GitHubRawBase      = "https://raw.githubusercontent.com/notmeower77463955/catman-files/refs/heads/main/packages"
	GitHubPackageList  = "https://raw.githubusercontent.com/notmeower77463955/catman-files/refs/heads/main/packages/package_list"
	VersionServer	   = "https://raw.githubusercontent.com/notmeower77463955/catman-files/refs/heads/main/VERSION"
	MetadataURLTemplate = GitHubRawBase + "/%s/%s.cat"
	BuildScriptURLTemplate = GitHubRawBase + "/%s/%s.sh"
)

// actually make it be useful eta 2079