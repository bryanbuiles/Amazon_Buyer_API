# Amazon Buyer API

Amazon Api buyer is an API where you can display the purchases of an user, see all their products and get some some recommendations for future purchases. The API takes the information from the amazon API in a determined day and saved it in the database.

The api is still under development and the database is not populate with movies yet.

## Functionalities of this api:

- Save the information of buyers, products and transactions of a given day of amazon.
- It shows all the buyers stored in the database who have made purchases at amazon.
- Shows all products and some purchase recommendations for a given buyer.
- Shows all buyers that have the same ip.

## Table of Content

- [Environment](#environment-and-requirements)
- [Run api locally](#Run-api-locally)
- [Endpoints](#Endpoints)
- [Future improvements](#Future-improvements)
- [Bugs](#bugs)
- [Authors](#authors)
- [License](#license)

## Environment and requirements

This web-application was interpreted/tested on Ubuntu 20.04 LTS using go (version 1.15.6)

### General Requirements

- go chi/chi: https://github.com/go-chi/chi
- Dgraph
- Goland
- Vue.js

## Run api locally

- Clone this repository: `git clone "https://github.com/bryanbuiles/Amazon_Buyer_API.git"`
- Access to cmd folder: `cd Amazon-Buyer-API/cmd`
- Update dependecies
  ```
  ~/Amazon-Buyer-API$ go mod tidy
  ```
- Run Dgraph database: https://dgraph.io/docs/get-started/
- Run the api:
  ```
  ~/Amazon-Buyer-API$ go run main.go server.go
  ```
- Request the endpoints with curl or postmant:

## Endpoints

### Save all data by day:

- GET /load - Save data of the current day
- GET /load?date=YYYY-MM-dd - Save data for a given day

  Example:

  ```
  curl -X GET http://localhost:3000/load?date=2015-10-15
  ```

  Output:

  ```
  {"message": "Data saved successfully"}
  ```

### Buyer information:

- GET /buyer - Display all buyers
- GET /buyer/{userID} - Get a buyer by id also display. It also shows users with the same ip and some product recommendation

  Example:

  ```
  curl -X GET http://localhost:3000/buyer/a063c6e2'
  ```

  Output:

  ```
  {
    "BuyerInfo": [
        {
            "id": "a063c6e2",
            "name": "Lenz",
            "purchases": [
                {
                    "transactionID": "0000561f0a47",
                    "ip": "129.165.209.171",
                    "products": [
                        {
                            "id": "555c06e1",
                            "name": "Vegetable barley soups",
                            "price": 4605
                        },
                        {
                            "id": "e287ef01",
                            "name": "Progresso Vegetable Classics Hearty Tomato Soup",
                            "price": 9582
                        }
                    ]
                },
                {
                    "transactionID": "0000561f00e9",
                    "ip": "154.189.124.137",
                    "products": [
                        {
                            "id": "bae389d5",
                            "name": "Campbell's well yes soup chickpea and red pepper",
                            "price": 4202
                        }
                    ]
                }
            ]
        }
    ],
    "sameIP": [
        {
            "ip": "129.165.209.171",
            "id": "e2475dfe",
            "name": "Pavior"
        },
        {
            "ip": "129.165.209.171",
            "id": "6596e02e",
            "name": "Whitelaw"
        },
        {
            "ip": "129.165.209.171",
            "id": "c60786ac",
            "name": "Neom"
        }
    ],
    "recomendations": [
        {
            "id": "467880b7",
            "name": "Jasmine rice",
            "price": 2713
        },
        {
            "id": "443067fa",
            "name": "Spinach & cheddar organic tortilla",
            "price": 3075
        }
    ]
  }
  ```

## Future improvements

- Add Frontend - doing
- Docker

## Bugs

No known bugs at this time.

## Authors

- Brayam Builes - [Github](https://github.com/bryanbuiles) / [Twitter](https://twitter.com/bryan_builes) / [Linkedin](https://www.linkedin.com/in/brayam-steven-builes-echavarria/)

## License

Apache-2.0 License.
