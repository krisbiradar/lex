package internal

import "os"

type Setup struct {
}

func (s *Setup) LoadModel(config *Config) bool {
	if config.ModelPath == "" {
		GetLogger(config).logError("No model path specified")
		return false
	}
	_, err := os.Stat(config.ModelPath)
	if err == nil {
		// start loading and spinning things up
		return true
	} else if os.IsNotExist(err) {
		GetLogger(config).logError("model path does not exist \n load Failed")
		return false
	}

	GetLogger(config).logErrorException(err, "something went wrong")
	return false

}
