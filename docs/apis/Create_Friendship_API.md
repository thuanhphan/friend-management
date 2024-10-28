**Add Customer Feedback API**

* **URL** `/make-friends`
* **Method:** `POST`

* **Request Body**

```
{
    "user_email": "example@example.com",
    "friend_email": "anhthu@example.com",
    "status": "friend"
}
```

* **Success Response:**  
  `HTTP_STATUS: 200`
  ```
{
    "success": true
}
  ```
