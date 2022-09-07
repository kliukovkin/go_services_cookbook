# Rate limiter
Что такое рейт лимитер? Почитайте нах в гугле, а сегодня мы реализуем самые популярные алгоритмы rate limiter. 

Для начала, 
- вот наш сервис:
- вот юнит тестик, который будет за секунду n раз дергать наш сервис и проверять его работу
- вот миддлвар, в котором мы будем мутить хуйню-муйню
- теперь нам нужен клиент для редиса и сам редис(зачем и почему)

go get github.com/go-redis/redis/v9


ну и погнали, вот алгоритм такой-то
- токен букет
как работай алгоритма, понимаешь?
почему редис?(если что можно использовать postgres) рейс кондишены? как их обойти по бырому?
- лики букет
-



сказать, что лучше вынести rate-limiter в load-balancer???? - it depends on your application. For example, you can use rate-limiter within your API Gateway together with handwritten caching, you can configure it inside load-balancer(e.g. NGINX) and you can use it in your application itself as a middleware before handler.

Where API gateway will be placed in case of Load Balancer?

## What Are the Major Types of Rate Limiting? - полностью переписать раздел, он украден [отсюда](https://systemsdesign.cloud/SystemDesign/RateLimiter)
There are several major types of rate limiting models that a business can choose between depending on which one offers the best fit for a business based on the nature of the web services that they offer, as we will explore in greater detail below.

- **User-Level Rate Limiting**  
In cases where a system can uniquely identify a user, it can restrict the number of API requests that a user makes in a time period. For example, if the user is only allowed to make two requests per second, the system denies the user’s third request made in the same second. User-level rate limiting ensures fair usage. However, maintaining the usage statistics of each user can create an overhead to the system that if not required for other reasons, could be a drain on resources.
- **Server-Level Rate Limiting**  
Most API-based services are distributed in nature. That means when a user sends a request, it might be serviced by any one of the many servers. In distributed systems, rate limiting can be used for load-sharing among servers. For example, if one server receives a large chunk of requests out of ten servers in a distributed system and others are mostly idle, the system is not fully utilized. There will be a restriction on the number of service requests that a particular server can handle in server-level rate limiting. If a server receives requests that are over this set limit, they are either dropped or routed to another server. Server-level rate limiting ensures the system’s availability and prevents denial of service attacks targeted at a particular server.
 - **Geography-Based Rate Limiting**  
Most API-based services have servers spread across the globe. When a user issues an API request, a server close to the user’s geographic location fulfils it. Organizations implement geography-based rate limiting to restrict the number of service requests from a particular geographic area. This can also be done based on timing. For example, if the number of requests coming from a particular geographic location is small from 1:00 am to 6:00 am, then a web server can have a rate limiting rule for this particular period. If there is an attack on the server during these hours, the number of requests will spike. In the event of a spike, the rate limiting mechanism will then trigger an alert and the organization can quickly respond to such an attack.


## algorithms
### token bucket rewrite
pros:
- allows traffic spikes and burst of traffc.
- memory efficient
- easy to implement
cons:
- race conditions

*To ensure smoother traffic, the refill rate of the token should be different from the rate limit.*

*Say our rate limit is four requests per minute. Instead of replenishing four tokens per minute, we can top up one token per 15 seconds. This prevents the sudden burst of traffic at the reset boundary.* - both paragraphs нагло спизжены и должны быть видоизменены

Where used:  
https://pkg.go.dev/golang.org/x/time/rate
https://docs.aws.amazon.com/AWSEC2/latest/APIReference/throttling.html

### leaky bucket DONE
The main idea behind leaky bucket is it guarantees that speed of proceeding requests will not be greater than the rate limit. You assume that your will proceed no more than let's say 2 requests per second. And If you will have overload, you will store this overload in the bucket.

Посмотреть в репе, по идее этот алгоритм в одном месте реализовывать смысла мало, получается говно. Типа в одном месте очередь, консюмеры и продюсеры
It takes two parameters:
- outflow rate. Speed of proceeding requests(For example we have a rate limit of 1 request per 100 milliseconds).
- bucket(queue) size. When we have second request earlier than 100 milliseconds, it will be places in the queue, if it was empty. If the queue is full, the request will be dropped.

cons:
- slower response time(last in the queue request should wait until all previous requests are processed)
- A burst of traffic fills up the queue with old requests and starve more recent requests from being processed

https://github.com/uber-go/ratelimit  
https://www.nginx.com/blog/rate-limiting-nginx/





### fixed window counter DONE
![fixed window counter](Screenshot%202022-09-06%20at%2018.59.04.png)
This algorithm is very similar to the token bucket algorithm. There are two main parameters:
- rate limit. Speed of proceeding requests.
- window size. The time period in which the rate limit is applied.
There is no refill rate which can make traffic smoother and, as a result the burst of traffic x2 is possible near borders of the time frame. ![](Screenshot%202022-09-06%20at%2019.01.09.png) 

But, unlike the token bucket algorithm, with this approach’s Redis operations are atomic. 
pros:
- memory efficient
- easy to implement
- suitable when it is needed to reduce the number of requests for a particular time period.
cons:
- trafic bursts are not prevented

### sliding window logs

### sliding window counter

## Rate-limiting in distributed systems TODO
https://systemsdesign.cloud/SystemDesign/RateLimiter

https://designgurus.org/path-player?courseid=grokking-the-system-design-interview&unit=grokking-the-system-design-interview_1626970236163_6Unit

https://konghq.com/blog/how-to-design-a-scalable-rate-limiting-algorithm




# Used materials:
https://konghq.com/blog/how-to-design-a-scalable-rate-limiting-algorithm  

https://aaronice.gitbook.io/system-design/system-design-problems/designing-an-api-rate-limiter  

https://betterprogramming.pub/4-rate-limit-algorithms-every-developer-should-know-7472cb482f48  

https://designgurus.org/path-player?courseid=grokking-the-system-design-interview&unit=grokking-the-system-design-interview_1626970236163_6Unit  

https://systemsdesign.cloud/SystemDesign/RateLimiter
