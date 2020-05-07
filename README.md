# theAnalyst
Financial news web page with sentiment analysis. Website deployed at: https://fanjason.github.io/the-analyst/

## Run Locally:
* Clone the repository
* Get API keys from https://monkeylearn.com/ and https://newsapi.org/
* Create a .env file in the root directory and paste in:
```sh
NEWS: REPLACE_WITH_NEWSAPI_KEY
SENTIMENT: REPLACE_WITH_MONKEYLEARN_APIKEY
```
* To start the server:
```sh
cd server
go run main.go
```
This will run the server on localhost:8080

* To start the client:
```sh
Open a new terminal in the root directory of the project
cd client
npm start
```
Project is ready to go on localhost:3000!
