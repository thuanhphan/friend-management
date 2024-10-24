**Add Customer Feedback API**

* **URL** `/receivable-updates`
* **Method:** `POST`

* **Request Body**

```
{
    "email": "example@example.com"
}
```

* **Success Response:**  
  `HTTP_STATUS: 200`
  ```
{
    "success": true,
    "recipients": [
        "anhthu@example.com",
        "alex@example.com",
        "tester@example.com"
    ]
}
  ```
