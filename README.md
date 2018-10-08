# magneto-challenge-apicore

## API calls

### Examples

POST

```
curl -X POST /mutant -H "Content-Type: application/json" -d '{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}'
```
Response

	HTTP 200 / 400

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
