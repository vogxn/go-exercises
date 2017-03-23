package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

/**
{
  "ResponseMetadata": {
    "HTTPHeaders": {
      "content-type": "text/xml",
      "content-length": "1075",
      "date": "Mon, 10 Oct 2016 07:53:49 GMT",
      "x-amzn-requestid": "ae0173ef-8ebe-11e6-a70a-a16d9da03b47"
    },
    "RequestId": "ae0173ef-8ebe-11e6-a70a-a16d9da03b47",
    "HTTPStatusCode": 200,
    "RetryAttempts": 0
  },
  "Credentials": {
    "AccessKeyId": "ASIAJLETHLL5HNJYJLZQ",
    "Expiration": "2016-10-10T08:53:49Z",
    "SessionToken": "FQoDYXdzEEEaDBM98NmTKmLuCjS4rSLdAbR7+TzNHDXgQfsdXdPWjIgQNW5yEvUgC8eR0j0mwxg9lFKy3Y/w/3RtwX2hDPOpMo7NPVsAFUcmWm1T3FYwjxHlSyZ1bauymt7QAqV6tCiAGVsl+aV3Q2Ga1bj3Yf42ONpNCKorDq9TyGBzyzH4RXpr79tHqQBrIwjvmorZY3kUYhuiCxngzSSdNaRg+ylNrwNGDspA3LA4ZY7n1ZIUqeRBk
VcprDqGB4ZVFU3h/L3kh8s7bFtQI7Debf7EQ4HMWR/UA0vOmyNJzCnxspPIUq6WZt9JWjC4U1PsF4RfKI2S7b8F",
    "SecretAccessKey": "iIPQ/1ZpiGSNP2l8LcVVDJdbJdjAGGNUscEC1pFg"
  },
  "AssumedRoleUser": {
    "Arn": "arn:aws:sts::147460200367:assumed-role/tsp-role-devexternal/AWS-CLI-session-1476086028",
    "AssumedRoleId": "AROAIUVMKI67DJMLIGXAE:AWS-CLI-session-1476086028"
  }
}
*/

type Role struct {
	AccessKey    string
	SecretKey    string
	SessionToken string
	Expiration   time.Time
}

func main() {
	t, _ := time.Parse(time.RFC3339, "2016-10-10T08:53:49Z")
	r := Role{"ASIAJLETHLL5HNJYJLZQ", "iIPQ/1ZpiGSNP2l8LcVVDJdbJdjAGGNUscEC1pFg", "FQoDYXdzEEEaDBM98NmTKmLuCjS4rSLd", t}
	j := json.NewEncoder(os.Stdout)
	fmt.Println("writing role")
	j.SetIndent("", "  ")
	j.Encode(r)
}
