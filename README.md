# echo-learning

The code examples for Echo web framework. 
The code examples are based on the [tutorial](https://echo.labstack.com/guide/) of the Echo web framework, but they are enhanced on the following aspects:
- Make sure each examples are complete and runnable.
- Create [a Postman collection](postman/echo-learning.postman_collection.json) to test the examples.
- Keep the same stucture between the Postman collection and this repository (It is easy to find the corresponding Postman request to test the example you want).

## Repository structure
- **bind**: The examples for binding.
   - json: The JSON binding example of binding JSON from request body.
- **case**: The example cases (Larger than funtional examples).
   - jwt: The JWT authentication example case.
      - customclaim: The JWT example by using custom claims
      - mapclaim: The JWT example by using map claims.
      - refreshtoken: The JWT example with refresh token functionality.
- **logging**: The examples of logging.
   - request: The examples of request logging.
      - customformat: The example of the custom format on request logging.
      - default: The example of default request logging.
   - system: The examples of system logging.
      - customformat: The example of the custom format on system logging.
      - level: The example of different log levels.
      - outputtofile: The example of outputing logging message to a log file.
- **middleware**: The examples of middlewares.
   - auth: The examples of authentication middlewares.
      - basicauth: The example of basic auth.
- **request**: The examples of request.
   - getdata: The example of getting data from the different places (path parameter, request header, etc.) of the request.
   - route: The example of group route.
- **response**: The examples of response.
   - return: The examples of returning different types of data.
      - html: The example of responding HTML.
      - json: The example of responding JSON.
      - nocontent: The example of responding no content.
      - string: The example of responding string.
      - xml: The example of responding XML.
- **postman**: The postman collection.

