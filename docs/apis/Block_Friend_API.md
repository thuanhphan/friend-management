**Add Customer Feedback API**

* **URL** `/block`
* **Method:** `POST`

* **Request Body**

```
{
    "user_email": "example@example.com",
    "friend_email": "anhthu@example.com",
    "status": "blocked"
>>>>>>> 856e22d (Refactor structure)
}
```

* **Success Response:**  
  `HTTP_STATUS: 200`
  ```
{
    "success": true
}
  ```
