package generates

import (
	"math/rand"
	"time"
)


func init()  {
	rand.Seed(time.Now().UnixNano())
}
