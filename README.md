# echo-learning

The code examples for Echo web framework. 
The code examples are based on the [tutorial](https://echo.labstack.com/guide/) of the Echo web framework, but they are enhanced on the following aspects:
- Make sure each examples are complete and runnable.
- Create [a Postman collection](postman/echo-learning.postman_collection.json) to test the examples.
- Keep the same stucture between the Postman collection and this repository (It is easy to find the corresponding Postman request to test the example you want).

## Repository structure
- **bind**
   - json: The JSON binding example for binding JSON from request body.
- **case**
   - jwt: The JWT authentication example case.
      - customclaim: The JWT example by using custom claims
      - mapclaim: The JWT example by using map claims.
      - refreshtoken: The JWT example with refresh token functionality.
- **logging**
