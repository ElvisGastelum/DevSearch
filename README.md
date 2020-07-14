# DevSearch
A slack bot in go


## Simple use example 

```go
package main

import (
	"github.com/elvisgastelum/devsearchbot"	
)

// Before run this program i set up the enviroment variable "SLACK_ACCESS_TOKEN"
func main() {
	bot := devsearchbot.Bot{}

	bot.Start()
}
```