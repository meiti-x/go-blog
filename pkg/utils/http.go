package utils

// Get config path for local or docker
func GetConfigPath(path string) string {
	if path == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}
