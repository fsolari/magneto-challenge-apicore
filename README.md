# Magneto API

Magneto wants to recruit as much mutants as possible to fight the X-Men.

This API will help Magneto find out if a human is a mutant, basing on it's ADN sequence.

ADN is a NxN matrix of chars. Available chars are: A,T,C,G
A mutant ADN contains more than one sequence of consecutive chars in horizontal, vertical or diagonal direction.

Written in Golang 1.9.1, deployed on Heroku and provisioned with a JawsDB MySQL.

## Demo

https://magneto-challenge-apicore.herokuapp.com

## API calls

### Examples

POST

```
curl -X POST /mutant -H "Content-Type: application/json" -d '{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}'
```
Response

	HTTP 200 {true} / 403 {false}

GET
```
curl - X GET /stats
```
Response

	HTTP 200
	{
	    "count_mutant_dna":3,
	    "count_human_dna":2,
	    "ratio":0.66
    }

## Benchmark


