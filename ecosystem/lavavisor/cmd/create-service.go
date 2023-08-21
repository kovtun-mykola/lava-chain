package lavavisor

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cosmos/cosmos-sdk/client/flags"
	processmanager "github.com/lavanet/lava/ecosystem/lavavisor/pkg/process"
	lvutil "github.com/lavanet/lava/ecosystem/lavavisor/pkg/util"
	"github.com/lavanet/lava/protocol/common"
	"github.com/lavanet/lava/protocol/lavasession"
	"github.com/lavanet/lava/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type ServiceParams struct {
	ServiceType               string
	ServiceConfigFile         string
	FromUser                  string
	KeyringBackend            string
	ChainID                   string
	GeoLocation               uint64
	LogLevel                  string
	Node                      string
	LavavisorServicesDir      string
	LavavisorLogsDir          string
	LavavisorServiceConfigDir string
}

func CreateLavaVisorCreateServiceCobraCommand() *cobra.Command {
	cmdLavavisorCreateService := &cobra.Command{
		Use:   `create-service [service-type: "provider" or "consumer"] [service-config-folder]`,
		Short: "generates service files for each provider/consumer in the config.yml.",
		Long: `The 'create-service' command generates system service files for each provider 
		and consumer specified in the config.yml file. Once these service files are created, 
		the 'lavavisor start' command can utilize them to manage (enable, restart, and check the status of) 
		each service using the 'systemctl' tool. This ensures that each service is properly integrated with 
		the system's service manager, allowing for robust management and monitoring of the LavaVisor services.
		Each service file inside [service-config-folder] must be named exactly the same with corresponding service name
		defined in config.yml`,
		Args: cobra.ExactArgs(2),
		Example: `required flags: --geolocation | --from
			optional flags: --log-level  | --node  | --keyring-backend
			lavavisor create-service ./config --geolocation 1 --from alice --log-level warn
			lavavisor create-service ./config --geolocation 1 --from bob --log-level info`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// 1. read config.yml -> this will tell us what service files this command will create
			dir, _ := cmd.Flags().GetString("directory")
			// Build path to ./lavavisor
			lavavisorPath, err := processmanager.ValidateLavavisorDir(dir)
			if err != nil {
				return err
			}
			// .lavavisor/ main services dir
			lavavisorServicesDir := lavavisorPath + "/services"
			err = os.MkdirAll(lavavisorServicesDir, 0o755)
			if err != nil {
				return utils.LavaFormatError("failed to create services directory", err)
			}

			// .lavavisor/ service logs dir
			lavavisorLogsDir := lavavisorServicesDir + "/logs"
			err = os.MkdirAll(lavavisorLogsDir, 0o755)
			if err != nil {
				return utils.LavaFormatError("failed to create service logs directory", err)
			}

			// .lavavisor/ service config dir
			lavavisorServiceConfigDir := lavavisorServicesDir + "/service_configs"
			err = os.MkdirAll(lavavisorServiceConfigDir, 0o755)
			if err != nil {
				return utils.LavaFormatError("failed to create service logs directory", err)
			}

			// GET SERVICE PARAMS
			serviceType := args[0]
			if serviceType != "provider" && serviceType != "consumer" {
				return utils.LavaFormatError("invalid service type, must be provider or consumer", nil)
			}
			serviceConfigFile := args[1] // the path that contains provider or consumer's configuration yml file
			// from user
			fromUser, _ := cmd.Flags().GetString(flags.FlagFrom)
			// keyring backend
			keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)
			// chainId
			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			// geolocation
			geoLocation, _ := cmd.Flags().GetUint64(lavasession.GeolocationFlag)
			// log-level
			logLevel, _ := cmd.Flags().GetString(flags.FlagLogLevel)
			// log-level
			node, _ := cmd.Flags().GetString(flags.FlagNode)

			serviceParams := &ServiceParams{
				ServiceType:               serviceType,
				ServiceConfigFile:         serviceConfigFile,
				FromUser:                  fromUser,
				KeyringBackend:            keyringBackend,
				ChainID:                   chainID,
				GeoLocation:               geoLocation,
				LogLevel:                  logLevel,
				Node:                      node,
				LavavisorServicesDir:      lavavisorServicesDir,
				LavavisorLogsDir:          lavavisorLogsDir,
				LavavisorServiceConfigDir: lavavisorServiceConfigDir,
			}

			utils.LavaFormatInfo("Creating the service file")

			filename := filepath.Base(serviceConfigFile)
			configName := filename[0 : len(filename)-len(filepath.Ext(filename))]
			configPath := filepath.Dir(serviceConfigFile)

			viper.SetConfigName(configName)
			viper.SetConfigType("yml")
			viper.AddConfigPath(configPath)

			// Read the configuration file
			err = viper.ReadInConfig()
			if err != nil {
				return utils.LavaFormatError("Error reading config file", err)
			}

			serviceFileName, err := CreateServiceFile(serviceParams)
			if err != nil {
				return err
			}
			// Write the name of the service into .lavavisor/config.yml
			return WriteToConfigFile(lavavisorPath, serviceFileName)
		},
	}
	flags.AddTxFlagsToCmd(cmdLavavisorCreateService)
	cmdLavavisorCreateService.MarkFlagRequired(flags.FlagFrom)
	cmdLavavisorCreateService.MarkFlagRequired(flags.FlagChainID)
	cmdLavavisorCreateService.Flags().Uint64(common.GeolocationFlag, 0, "geolocation to run from")
	cmdLavavisorCreateService.MarkFlagRequired(common.GeolocationFlag)
	cmdLavavisorCreateService.Flags().String("directory", os.ExpandEnv("~/"), "Protocol Flags Directory")
	cmdLavavisorCreateService.Flags().String(flags.FlagLogLevel, "debug", "log level")
	return cmdLavavisorCreateService
}

func CreateServiceFile(serviceParams *ServiceParams) (string, error) {
	// working dir:
	out, err := exec.LookPath("lava-protocol")
	if err != nil {
		return "", utils.LavaFormatError("could not detect a linked lava-protocol binary", err)
	}
	workingDir := strings.TrimSpace(filepath.Dir(out) + "/")

	if _, err := os.Stat(serviceParams.ServiceConfigFile); err != nil {
		return "", utils.LavaFormatError("Service config file not found", err)
	}

	configChainID := viper.GetString("endpoints.0.chain-id")
	serviceId := serviceParams.ServiceType + "-" + configChainID
	configPath := serviceParams.LavavisorServiceConfigDir + "/" + filepath.Base(serviceParams.ServiceConfigFile)

	err = lvutil.Copy(serviceParams.ServiceConfigFile, configPath)
	if err != nil {
		return "", utils.LavaFormatError("couldn't copy binary to system path", err)
	}

	content := "[Unit]\n"
	content += "  Description=" + serviceId + " daemon\n"
	content += "  After=network-online.target\n\n"
	content += "[Service]\n"
	content += "  WorkingDirectory=" + workingDir + "\n"
	if serviceParams.ServiceType == "consumer" {
		content += "  ExecStart=" + workingDir + "/lava-protocol rpcconsumer "
	} else if serviceParams.ServiceType == "provider" {
		content += "  ExecStart=" + workingDir + "/lava-protocol rpcprovider "
	}
	content += ".lavavisor/services/service_configs/" + filepath.Base(serviceParams.ServiceConfigFile) + " --from " + serviceParams.FromUser + " --keyring-backend " + serviceParams.KeyringBackend + " --chain-id " + serviceParams.ChainID + " --geolocation " + fmt.Sprint(serviceParams.GeoLocation) + " --log_level " + serviceParams.LogLevel + " --node " + serviceParams.Node + "\n"

	content += "  User=ubuntu\n"
	content += "  Restart=always\n"
	content += "  RestartSec=180\n"
	content += "  LimitNOFILE=infinity\n"
	content += "  LimitNPROC=infinity\n"
	content += "  StandardOutput=append:" + serviceParams.LavavisorLogsDir + "/" + serviceId + ".log\n\n"
	content += "[Install]\n"
	content += "  WantedBy=multi-user.target\n"

	filePath := serviceParams.LavavisorServicesDir + "/" + serviceId + ".service"
	err = os.WriteFile(filePath, []byte(content), os.ModePerm)
	if err != nil {
		return "", utils.LavaFormatError("error writing to service file", err)
	}
	utils.LavaFormatInfo("Service file has been created successfully", utils.Attribute{Key: "Path", Value: filePath})
	// Extract filename from filePath
	filename := filepath.Base(filePath)

	return filename, nil
}

func WriteToConfigFile(lavavisorPath string, serviceFileName string) error {
	configPath := filepath.Join(lavavisorPath, "config.yml")
	file, err := os.OpenFile(configPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return utils.LavaFormatError("error opening config.yml for appending", err)
	}
	defer file.Close()
	// Check if the file is newly created by checking its size
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	if fileInfo.Size() == 0 {
		_, err = file.WriteString("services:\n")
		if err != nil {
			return err
		}
	}
	// Append the new service name
	_, err = file.WriteString("  - " + serviceFileName + "\n")
	if err != nil {
		return err
	}
	return nil
}
