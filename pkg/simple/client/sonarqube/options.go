package sonarqube

import (
	"github.com/spf13/pflag"
)

type SonarQubeOptions struct {
	Host  string `json:",omitempty" yaml:",omitempty" description:"SonarQube service host address"`
	Token string `json:",omitempty" yaml:",omitempty" description:"SonarQube service token"`
}

func NewSonarQubeOptions() *SonarQubeOptions {
	return &SonarQubeOptions{
		Host:  "",
		Token: "",
	}
}

func NewDefaultSonarQubeOptions() *SonarQubeOptions {
	return NewSonarQubeOptions()
}

func (s *SonarQubeOptions) Validate() []error {
	errors := []error{}

	return errors
}

func (s *SonarQubeOptions) ApplyTo(options *SonarQubeOptions) {
	if options == nil {
		return
	}

	if s.Host != "" {
		options.Host = s.Host
	}

	if s.Token != "" {
		options.Token = s.Token
	}
}

func (s *SonarQubeOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.Host, "sonarqube-host", s.Host, ""+
		"Sonarqube service address if enabled.")

	fs.StringVar(&s.Token, "sonarqube-token", s.Token, ""+
		"Sonarqube service access token.")
}
