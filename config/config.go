package config

import (
	"strings"

	"github.com/fitraditya/cloudfs/storage"
	"github.com/fitraditya/cloudfs/storage/dropbox"
	"github.com/fitraditya/cloudfs/storage/filesystem"
	"github.com/fitraditya/cloudfs/storage/gdrive"
	"github.com/fitraditya/cloudfs/storage/s3"
	"github.com/obrel/go-lib/pkg/log"
	"github.com/spf13/viper"
)

// GetStorage
func GetStorage() string {
	return viper.GetString("storage")
}

// GetStorageOption
func GetStorageOption() []storage.Option {
	switch GetStorage() {
	case "filesystem":
		return []storage.Option{
			filesystem.BasePath(viper.GetString("storage_option.base_path")),
		}
	case "s3":
		return []storage.Option{
			s3.Region(viper.GetString("storage_option.region")),
			s3.Bucket(viper.GetString("storage_option.bucket")),
			s3.AccessKey(viper.GetString("storage_option.access_key")),
			s3.AccessSecret(viper.GetString("storage_option.access_secret")),
		}
	case "dropbox":
		return []storage.Option{
			dropbox.Token(viper.GetString("storage_option.token")),
		}
	case "gdrive":
		return []storage.Option{
			gdrive.ClientId(viper.GetString("storage_option.client_id")),
			gdrive.ClientSecret(viper.GetString("storage_option.client_secret")),
			gdrive.Token(viper.GetString("storage_option.token")),
		}
	}

	return nil
}

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//
		} else {
			log.For("config", "init").Fatal(err)
		}
	}
}
