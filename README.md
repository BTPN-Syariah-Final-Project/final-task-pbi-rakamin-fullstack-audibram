# M-Banking Application

The M-Banking application is a secure mobile banking platform that allows users to perform various banking operations conveniently.
This file provides comprehensive instructions on setting up, running, and using the application effectively.

**Table of contents**
* [Setup](https://github.com/BTPN-Syariah-Final-Project/final-task-pbi-rakamin-fullstack-audibram#setup)
* [Register User](https://github.com/BTPN-Syariah-Final-Project/final-task-pbi-rakamin-fullstack-audibram#register-user)
* [Login User](https://github.com/BTPN-Syariah-Final-Project/final-task-pbi-rakamin-fullstack-audibram#login-user)
* [Update User Information](https://github.com/BTPN-Syariah-Final-Project/final-task-pbi-rakamin-fullstack-audibram#update-user-information)
* [Upload Profile Photo](https://github.com/BTPN-Syariah-Final-Project/final-task-pbi-rakamin-fullstack-audibram#upload-profile-photo)
* [Delete Profile Photo](https://github.com/BTPN-Syariah-Final-Project/final-task-pbi-rakamin-fullstack-audibram#delete-profile-photo)
* [Get Photo](https://github.com/BTPN-Syariah-Final-Project/final-task-pbi-rakamin-fullstack-audibram#get-photo)
* [Get User Login Information](https://github.com/BTPN-Syariah-Final-Project/final-task-pbi-rakamin-fullstack-audibram#get-user-login-information)

## Setup
To run the M-Banking application locally, follow these steps:
1. __Clone the Repository:__
```
git clone <repository_url>
``` 
2. __Navigate to the Project Directory:__
```
cd m-banking-app
``` 
3. __Build and Run Docker Compose:__
```
docker-compose up -d
``` 
4. __Access the Application__
   Once the containers are up and running, the application will be accessible at `http://localhost:8080`


## Usage
### Register User
To register a new user, send a `POST` request to `/users/register` with the following JSON payload:
```json
{
    "username": "example_user",
    "email": "user@example.com",
    "password": "secure_password"
}

```

### Login User
To log in, send a `POST` request to `/users/login` with the following JSON payload:
```json
{
    "email": "user@example.com",
    "password": "secure_password"
}
```
Upon successful login, the response will include a JWT token, which you should include in the `Authorization` header for subsequent requests.

### Update User Information
To update user information, send a `PUT` request to `/users/:userID` with the following JSON payload:
```json
{
    "username": "new_username",
    "email": "new_email@example.com",
    "password": "new_secure_password"
}
```
Ensure to replace `:userID` with the ID of the user to be updated. This endpoint requires authentication, so include the JWT token in the `Authorization` header.

### Upload Profile Photo
To upload a profile photo, send a `POST` request to `/photos` with the following JSON payload:
```json
{
    "title": "Profile Photo",
    "caption": "This is my profile photo",
    "photo_url": "https://example.com/profile_photo.jpg",
    "userID": 1
}
```
Ensure to replace `userID` with the ID of the user uploading the photo. This endpoint requires authentication, so include the JWT token in the `Authorization` header.

### Delete Profile Photo
To delete a profile photo, send a `DELETE` request to `/photos/:photoID`, where `:photoID` is the ID of the photo to be deleted. Ensure to include the JWT token in the `Authorization` header for authentication.

### Get Photo
To get a photo, send a `GET` request to `/photos/:photoID`, where `:photoID` is the ID of the photo to be retrieved.

### Get User Login Information
To get the logged-in user's information, send a `GET` request to `/users/login`. This endpoint requires authentication, so include the JWT token in the Authorization header.