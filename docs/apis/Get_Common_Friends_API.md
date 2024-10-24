**Add Customer Feedback API**

* **URL** `/common-friends`
* **Method:** `POST`

* **Request Body**

```
{
    "email1": "alex@example.com",
    "email2": "example@example.com"
}
```

* **Success Response:**  
  `HTTP_STATUS: 200`
  ```
{
    "success": true,
    "friends": [
        "tester@example.com"
    ],
    "count": 1
}
  ```
