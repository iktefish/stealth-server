# CURL Requests

-   Register an employee:
    ~~~
    curl -X POST 'localhost:8080/register-employee' -d @dummy/employee-register-form.json
    ~~~

-   Remove an employee:
    ~~~
    curl -X POST 'localhost:8080/remove-employee?id=...'
    ~~~

-   Get an employee (DEBUG):
    ~~~
    curl -X POST 'localhost:8080/get-employee?id=...' | jq
    ~~~
