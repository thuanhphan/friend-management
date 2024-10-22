### Requirements:
For any application with a need to build its own social network, "Friends Management" is a common requirement which usually starts off simple but can grow in complexity depending on the application's use case. Usually, applications would start with features like "Friend", "Subscribe", "Block", "Receive Updates" etc. 
Develop an API server that does simple "Friend Management" based on the User Stories below:
1 As a user, I need an API to create a friend connection between two email addresses.
2 As a user, I need an API to retrieve the friends list for an email address.
3 As a user, I need an API to retrieve the common friends list between two email addresses.
4 As a user, I need an API to subscribe to updates from an email address.
5 As a user, I need an API to block updates from an email address.
6 As a user, I need an API to retrieve all email addresses that can receive updates from an email address.

### Databases:

Table: users

| Column Name | Data Type                | Constraints                  | Nullable | Default |
|-------------|--------------------------|------------------------------|----------|---------|
| email       | TEXT                     | PRIMARY KEY                  | NO       |         |
| name        | TEXT                     |                              | YES      |         |


Table: friendships

| Column Name | Data Type                | Constraints                  | Nullable | Default |
|-------------|--------------------------|------------------------------|----------|---------|
| id          | SERIAL                   | PRIMARY KEY                  | NO       |         |
| user_email  | TEXT                     | FOREIGN KEY -> users(email)  | NO       |         |
| friend_email| TEXT                     | FOREIGN KEY -> users(email)  | NO       |         |
| status      | TEXT                     |                              | NO       |         |

### Sample table data 

Table: users

        email        |   name   
---------------------+----------
 example@example.com | John Doe 
 anhthu@example.com  | Anh Thu  
 alex@example.com    | Alex     
 tester@example.com  | Tester  

Table: friendships

 id |     user_email      |    friend_email    | status
----+---------------------+--------------------+--------
  6 | alex@example.com    | anhthu@example.com | friend
  7 | tester@example.com  | alex@example.com   | friend
  5 | example@example.com | anhthu@example.com | block

### Flows
1. UI will call this API [POST] `/make-friends` for creating a friend connection between two email addresses
2. UI will call this API [POST] `/friends` for retrieving the friends list for an email address
3. UI will call this API [POST] `/common-friends` for retrieving the common friends list between two email addresses
4. UI will call this API [POST] `/subscribe` for subscribing to updates from an email address
5. UI will call this API [POST] `/block` for blocking updates from an email address