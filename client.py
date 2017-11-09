import requests
import json


def postadd():
    data = {"CASE_NAME":"name1",
            "REPORT_PATH":"/111",
            "TYPE":0}
    # headers = {'Content-Type': 'application/x-www-form-urlencoded'}
    headers = {'Content-Type': 'application/json'}

    r = requests.post(
        "http://127.0.0.1:8009/v1/api/add",
        # "http://127.0.0.1:8009/v1/api/add1",
        data=json.dumps(data),
        # data=data,
        headers=headers)
    print r.text

def postupdate():
    data = {"REPORT_PATH":"/update",
            "STATUS":1,
            "ID":"e7869a94-d74c-47e6-ac50-f0330b9442f5"}
    # headers = {'Content-Type': 'application/x-www-form-urlencoded'}
    headers = {'Content-Type': 'application/json'}

    r = requests.post(
        "http://127.0.0.1:8009/v1/api/update",
        # "http://127.0.0.1:8009/v1/api/add1",
        data=json.dumps(data),
        # data=data,
        headers=headers)
    print r.text

if __name__ == '__main__':
    # postadd()
    postupdate()
