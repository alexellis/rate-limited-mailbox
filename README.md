# rate-limited-mailbox


## Deploy your mailbox

```
docker service rm mailbox
docker service create --network=func_functions --name mailbox --env gateway_url=http://gateway:8080  alexellis2/mailbox:latest
docker service logs -f mailbox
```

Now build/deploy your functions

```
sh ./test.sh
```

Finally invoke the `feeder` function. It uses an `api` function which is artifcally rate-limited to 5 requests.

> You can reset the rate-limit value by posting "reset" (no, new lines) to the api function. `echo -n .... | faas-cli invoke`

```
sh ./test.sh

successfully tagged api:latest
Image: api built.
[0] < Builder done.
Deploying: feeder.
Removing old function.
Deployed.
URL: http://localhost:8080/function/feeder

200 OK
Deploying: api.
Removing old function.
Deployed.
URL: http://localhost:8080/function/api

200 OK
bash-3.2$



bash-3.2$ echo test | faas-cli invoke feeder
None None
It's OK
bash-3.2$ echo test | faas-cli invoke feeder
None None
It's OK
bash-3.2$ echo test | faas-cli invoke feeder
None None
It's OK
bash-3.2$ echo test | faas-cli invoke feeder
None None
It's OK
bash-3.2$ echo test | faas-cli invoke feeder
None None
It's OK
bash-3.2$ echo test | faas-cli invoke feeder
None None
Not OK: "FAIL
"
Posting to Mailbox:  202
bash-3.2$
```

Now monitor the logs of the mailbox:

```
mailbox.1.mm87d5orcsxq@moby    | Try again with: feeder
mailbox.1.mm87d5orcsxq@moby    | 2017/11/09 17:13:07 Posting to http://gateway:8080/async-function/feeder/, status: 202 Accepted
mailbox.1.mm87d5orcsxq@moby    | 2017/11/09 17:13:07 header: map[Content-Length:[5] Accept-Encoding:[gzip, deflate] X-Delay-Duration:[5] X-Max-Retries:[5] Connection:[keep-alive] Accept
:[*/*] User-Agent:[python-requests/2.18.4] X-Retries:[3]]
mailbox.1.mm87d5orcsxq@moby    | 2017/11/09 17:13:07 Accepted work for feeder, retries: 3 max: 5, delay: 5.000000 secs
mailbox.1.mm87d5orcsxq@moby    | 2017-11-09 17:13:08.065040522 +0000 UTC m=+533.004187166 items= 1
mailbox.1.mm87d5orcsxq@moby    | 2017-11-09 17:13:09.063006622 +0000 UTC m=+534.002153251 items= 1
mailbox.1.mm87d5orcsxq@moby    | 2017-11-09 17:13:10.063095657 +0000 UTC m=+535.002242277 items= 1
mailbox.1.mm87d5orcsxq@moby    | 2017-11-09 17:13:11.062920488 +0000 UTC m=+536.002067113 items= 1
mailbox.1.mm87d5orcsxq@moby    | 2017-11-09 17:13:12.064587657 +0000 UTC m=+537.003734300 items= 1
mailbox.1.mm87d5orcsxq@moby    | 2017-11-09 17:13:13.063389758 +0000 UTC m=+538.002536404 items= 1
mailbox.1.mm87d5orcsxq@moby    | Try again with: feeder
mailbox.1.mm87d5orcsxq@moby    | 2017/11/09 17:13:13 Posting to http://gateway:8080/async-function/feeder/, status: 202 Accepted
mailbox.1.mm87d5orcsxq@moby    | 2017/11/09 17:13:13 header: map[Connection:[keep-alive] Accept-Encoding:[gzip, deflate] Accept:[*/*] User-Agent:[python-requests/2.18.4] X-Delay-Duratio
n:[5] X-Retries:[4] Content-Length:[5] X-Max-Retries:[5]]
mailbox.1.mm87d5orcsxq@moby    | 2017/11/09 17:13:13 Accepted work for feeder, retries: 4 max: 5, delay: 5.000000 secs
```


Before the function gets to its maximum retry limit, you can reset the API:

```
bash-3.2$ echo -n reset | faas-cli invoke api
OK
```

