# **Instagram Backend API**

This API is solely created as part of the task given by Appointy for internship recruitments. It is created using Go programming language and exploits MongoDB database for data storage, retrieval and modification

The endpoints of this API are mentioned below:

## **/users**
<!-- OL -->
1. Displays/Creates a user
2. Request type: GET, POST
3. GET Parameters: uid (begins with 1)
4. POST Parameters: Name, Email, Password
5. Password isn't stored directly ; it is hashed using sh1 and sha256 algorithms

## **/posts**
<!-- OL -->
1. Displays/Creates a post
2. Request type: GET, POST
3. GET Parameters: pid (begins with 1)
4. POST Parameters: Email, Caption, Image_URL, Posted_Timestamp

## **/posts/users**
<!-- OL -->
1. Displays a set number of posts (5) of a particular user
2. Exploites *pagination* in order to achieve this
2. Request type: GET
3. GET Parameters: uid (begins with 1), page (begins with 0)
4. page denotes the page number that the visitors wants to navigate to
5. A slice denoting containing user's post ids is present in the user schema which makes this query extremely efficient
