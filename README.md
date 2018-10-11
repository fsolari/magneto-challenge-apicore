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

  Total:  239.5674 secs.  
  Slowest:  1.4562 secs.  
  Fastest:  0.1817 secs.  
  Average:  0.2395 secs.  
  
  Requests/sec: 41.7419  
  Total Data Received:  56525 bytes.  
  Response Size per Request:  5 bytes.

**Response time histogram:**

  0.227 [9469]  |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎  
  0.383 [304] |∎  
  0.481 [92]  |  
  0.612 [112] |  
  0.716 [13]  |  
  0.819 [1] |  
  0.904 [2] |  
  0.992 [3] |    
  1.082 [2] |  
  1.441 [2] |

**Latency distribution:**
  
  10% in 0.2266 secs.  
  25% in 0.2266 secs.  
  50% in 0.2266 secs.  
  75% in 0.2266 secs.  
  90% in 0.2266 secs.  
  95% in 0.3832 secs.  
  99% in 0.6122 secs.

