package config

var allowedOrigins = []string{
	"https://barganakukuhraditya.github.io/",
}

// GetAllowedOrigins returns the list of allowed origins
func GetAllowedOrigins() []string {
	return allowedOrigins
}
