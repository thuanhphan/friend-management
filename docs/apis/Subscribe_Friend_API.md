**Add Customer Feedback API**

* **URL** `/block`
* **Method:** `POST`

* **Request Body**

```
{
    "user_email": "example@example.com",
    "friend_email": "anhthu@example.com",
<<<<<<< HEAD
    "status": "subscribed"
=======
    "status": "subscribe"
>>>>>>> 7343663 (Update docs, unit test)
}
```

* **Success Response:**  
  `HTTP_STATUS: 200`
  ```
{
    "success": true
}
  ```
