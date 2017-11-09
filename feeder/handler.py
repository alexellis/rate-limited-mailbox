import requests
import os


def requeue(r, st):
    max_retries = int(os.getenv("Http_X_Max_Retries", "5"))
    retries = int(os.getenv("Http_X_Retries", "0"))
    delay_duration = int(os.getenv("Http_X_Delay_Duration", "5"))

    retries = retries + 1

    headers = {
        "X-Retries": str(retries),
        "X-Max-Retries": str(max_retries),
        "X-Delay-Duration": str(delay_duration)
    }

    r = requests.post("http://mailbox:8080/deadletter/feeder", data=st, json=False, headers= headers)

    print "Posting to Mailbox: ", r.status_code
    if r.status_code!= 202:
        print "Mailbox says: ", r.text

def handle(st):
    print os.getenv("Http_X_Max_Retries"), os.getenv("Http_X_Retries")

    r = requests.post("http://api:8080/",  st)
    if str(r.text).strip() == "OK":
        print "It's OK"
    else:
        print "Not OK: \""+ r.text + "\""
        requeue(r, st)
