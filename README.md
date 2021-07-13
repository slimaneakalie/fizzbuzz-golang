# fizzbuzz-golang
This is a golang microservice for fizzbuzz.

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by "fizz", all multiples of 5 by "buzz", and all multiples of 15 by "fizzbuzz". The output would look like this: 
````
"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,..."
````

This service extends on this, it provides two endpoints:
### 1. `POST /v1/fizzbuzz`
* This endpoint accepts a json query like the following:
```json
{
  "firstInt": 3,
  "secondInt": 5,
  "limit": 100,
  "firstStr": "fizz",
  "secondStr": "buzz"
}
```
and it returns a list of strings with numbers from 1 to limit, where: all multiples of firstInt are replaced by firstStr, all multiples of secondInt are replaced by secondStr, all multiples of firstInt and secondInt are replaced by firstStrSecondStr.
#### Response Example (using the query above)
```json
{
  "fizzbuzzStringList": ["1", "2", "fizz", "4", "buzz", "fizz", "7", "8", "fizz", "buzz", "11", "fizz", "13", "14", "fizzbuzz", "16", "17", "fizz", "19", "buzz", "fizz", "22", "23", "fizz", "buzz", "26", "fizz", "28", "29", "fizzbuzz", "31", "32", "fizz", "34", "buzz", "fizz", "37", "38", "fizz", "buzz", "41", "fizz", "43", "44", "fizzbuzz", "46", "47", "fizz", "49", "buzz", "fizz", "52", "53", "fizz", "buzz", "56", "fizz", "58", "59", "fizzbuzz", "61", "62", "fizz", "64", "buzz", "fizz", "67", "68", "fizz", "buzz", "71", "fizz", "73", "74", "fizzbuzz", "76", "77", "fizz", "79", "buzz", "fizz", "82", "83", "fizz", "buzz", "86", "fizz", "88", "89", "fizzbuzz", "91", "92", "fizz", "94", "buzz", "fizz", "97", "98", "fizz", "buzz"]
}
```

### 2. `GET /v1/fizzbuzz/stats`
This statistics endpoint accepts no parameter and returns the body of the most frequent request (with successful response), and the number of hits.
The body of the service is something like the following:
```json
{
  "mostFrequentQuery": {
    "body": {
      "firstInt": 3,
      "secondInt": 5,
      "limit": 100,
      "firstStr": "fizz",
      "secondStr": "buzz"
    },
    "numberOfHits": 1000
  }
  
}
```

The stats under the hood are relying on prometheus metrics.

## Running
````
make run
````
P.S: You need to have docker installed on your machine

## API documentation
````
make doc
````

## Testing
````
make testing
````