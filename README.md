# skills-text-extration-go

Build app
LD_LIBRARY_PATH=${PWD}/libs CGO_CFLAGS=-I/${PWD} CGO_LDFLAGS=-L"${PWD}/libs -lrt -lm" CGO_ENABLED=1 GOOS=linux go build -o application .

Run app
export LD_LIBRARY_PATH=${PWD}/libs
./application


## /predict

POST Body

```json

{
	"query": "what is the show timing for starwars 9",
	"content": "The movies available for you to watch today are star wars 9, fast and furious, ABCD, EFGH, Rush Hour.  The first show for star wars 9 is at 10am and the following shows at 10:30am, 1pm, 2pm , 3pm, 7pm and midnight 12am."
}

```

Response:

```json
{
    "matches": [
        "The first show for star wars 9 is at 10am",
        "10am",
        "at 10am",
        "10am and the following shows at 10:30am, 1pm, 2pm , 3pm, 7pm and midnight 12am.",
        "at 10am and the following shows at 10:30am, 1pm, 2pm , 3pm, 7pm and midnight 12am."
    ]
}
```