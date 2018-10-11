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

Benchmarking was done using [pla](http://github.com/mercadolibre/pla) 

**Summary:**

  Total:	43.9388 secs.  
  Slowest:	1.2116 secs.  
  Fastest:	0.2885 secs.  
  Average:	0.4311 secs.
  
  Requests/sec:	227.5895  
  Total Data Received:	624657 bytes.  
  Response Size per Request:	62 bytes.
 
**Status code distribution:**

   [200]	4708 responses  
   [403]	26 responses  
   [400]	5266 responses  

**Response time histogram:**

  0.310 [12]	|  
  0.369 [116]	|  
  0.425 [9546]	|∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎  
  0.538 [155]	|  
  0.591 [45]	|  
  0.674 [37]	|  
  0.743 [42]	|  
  0.813 [20]	|  
  0.875 [26]	|  
  1.212 [1]	|

**Latency distribution:**
  
  10% in 0.4251 secs.  
  25% in 0.4251 secs.  
  50% in 0.4251 secs.  
  75% in 0.4251 secs.  
  90% in 0.4251 secs.  
  95% in 0.4251 secs.  
  99% in 0.6744 secs.

