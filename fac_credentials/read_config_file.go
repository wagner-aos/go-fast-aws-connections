package awscredentials

import (
	"bufio"

	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/kataras/golog"
)

var config *AWSConfigFile

//AWSConfigFile - temp
type AWSConfigFile struct {
	Profiles map[string]AWSProfile
}

//AWSProfile - type
type AWSProfile struct {
	Region    string
	AccountID string
	Role      string
	RoleARN   string
}

//Init - initializes config maps
func init() {
	config = &AWSConfigFile{}
	config.ReadConfig(sharedConfigFilename())
}

// ReadConfig reads info from config file
func (c *AWSConfigFile) ReadConfig(configfile string) {

	//New profiles map
	c.Profiles = make(map[string]AWSProfile)

	if _, err := os.Stat(configfile); err == nil {

		//Opening file
		file, err := os.Open(configfile)
		if err != nil {
			golog.Error(err)
		}
		defer file.Close()

		//Reading aws config file in order to extract values
		//for profile and role_arn
		scanner := bufio.NewScanner(file)
		keyMapStr := ""
		roleArnStr := ""
		regionStr := ""
		accountIDStr := ""
		roleStr := ""

		for scanner.Scan() {
			lineRead := scanner.Text()

			if strings.HasPrefix(lineRead, "[") {
				keyMapStr = strings.Replace(lineRead, "[", "", 1)
				keyMapStr = strings.Replace(keyMapStr, "profile", "", 1)
				keyMapStr = strings.Replace(keyMapStr, "]", "", 1)
				keyMapStr = strings.Trim(keyMapStr, " ")

			}
			if strings.Contains(lineRead, "role_arn") {
				accountIDStr, roleStr = splitRoleARN(lineRead)
				roleArnStr = strings.Replace(lineRead, "role_arn", "", 1)
				roleArnStr = strings.Replace(roleArnStr, "=", "", 1)
				roleArnStr = strings.Trim(roleArnStr, " ")
			}
			if strings.Contains(lineRead, "region") {
				regionStr = strings.Replace(lineRead, "region", "", 1)
				regionStr = strings.Replace(regionStr, "=", "", 1)
				regionStr = strings.Trim(regionStr, " ")
			}

			if len(keyMapStr) > 0 && len(roleArnStr) > 0 {

				c.Profiles[keyMapStr] = AWSProfile{
					Region:    regionStr,
					AccountID: accountIDStr,
					Role:      roleStr,
					RoleARN:   roleArnStr,
				}
				keyMapStr = ""
				roleArnStr = ""
				regionStr = ""
				accountIDStr = ""
				roleStr = ""
			}
		}
	} else {
		golog.Warnf("AWS Config: %s not found!!!", configfile)
	}

}

func splitRoleARN(lineRead string) (string, string) {

	accountID := ""
	role := ""

	resultRoleARNStr := strings.Split(lineRead, ":")

	for _, s := range resultRoleARNStr {
		re := regexp.MustCompile("[0-9]+")
		resultStr := re.FindAllString(s, -1)
		if len(resultStr) > 0 {
			accountID = s
		}
		if strings.Contains(s, "/") {
			role = strings.Replace(s, "role/", "", 1)
		}
	}

	return accountID, role

}

//GetProfile - role arn and transform in AWSProfile object
func GetProfile(profile string) AWSProfile {
	return config.Profiles[profile]
}

func sharedConfigFilename() string {
	return filepath.Join(UserHomeDir(), ".aws", "config")
}

// UserHomeDir returns the home directory for the user the process is
// running under.
func UserHomeDir() string {
	if runtime.GOOS == "windows" { // Windows
		return os.Getenv("USERPROFILE")
	}
	// *nix
	return os.Getenv("HOME")
}

//PrintProfiles - prints all profiles obtained from .aws/config file
func (c *AWSConfigFile) PrintProfiles() {
	for key, value := range c.Profiles {
		golog.Infof("PROFILE: %10s \n", key)
		golog.Infof("ACCOUNT_ID: %10s \n", value.AccountID)
		golog.Infof("REGION: %10s \n", value.Region)
		golog.Infof("ROLE: %10s \n", value.Role)
		golog.Infof("ROLE_ARN: %10s \n\n", value.RoleARN)
	}
}
