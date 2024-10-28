**Add Customer Feedback API**

* **URL** `/friends`
* **Method:** `POST`

* **Request Body**

```
{
    "email": "example@example.com",
}
```

* **Success Response:**  
  `HTTP_STATUS: 200`
  ```
{
    "success": true,
    "friends": [
        "anhthu@example.com",
        "tester@example.com"
    ],
    "count": 2
}
  ```
