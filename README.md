# Ready-to-deploy service

Service to determine if the code can be deployed depending on the weather conditions in Leipzig using a HTTP GET request.

There are two conditions to make the decision:
1. The current temperature is between -10ºC and 25ºC (both included)
2. The current wind speed is lower than 20 km/h

If both conditions are satisfied, you can deploy your code!

An expected response will look like the following one:
```json
{
  "deploy": true, // if the code can be deployed
  "current_temp": 20, // currrent temperature in degrees celsius
  "current_wind": 15, // current wind speed in km/h
  "error": "", // empty string if no errors happened during execution
  "cutie_fox": "http://...jpg" // link to a fox image for the developer enjoyment
}
```

## How to execute the code

To execute the code you will need to be located in the same directory as the `main.go` file and write these two commands:
```bash
go build -v # To compile the code
./main # To start your server in Linux
.\main.exe # To start your server in Windows
```

## Documentation

All the functions are properly documented. You just need to take a look on the files.

However, you can execute the following commands to run a web server to present the documentation as a web page:
```bash
go get golang.org/x/tools/cmd/godoc
godoc -hhtp=:6060
```

The documentation will be available in this [link](http://localhost:6060/pkg/main/server/).

## Tests

Unit tests have been added as well. For this service I have added them for the function that checks the weather condition and returns if the code can be deployed or not.

To execute all the test for all the packages of the project and see the results (PASS or FAIL), write the following command:
```bash
go test ./... -v 
```