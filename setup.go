package eureka

import (
	"os"
	"github.com/sirupsen/logrus"
)



func init(){
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.WarnLevel)
}
