package database

import "package-config-yaml/config"

func QueryURL() string {
	return config.Config().URL
}
