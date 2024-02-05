A user wants to count arithmetic expressions.
He enters the string `2 + 2 * 2` and wants to get `6` as the answer.
But our addition and multiplication operations (also division and subtraction) take a **"very, very" long time**.
Therefore, a scenario where a user makes an http request and receives the result as a response is **impossible**.
Moreover: the computation of each such operation in our **"alternative reality "** takes **"gigantic "** computing power.
Accordingly, we must be able to perform each action separately, and we can scale this system by adding computing power to our system in the form of new "**machines**".
Therefore, the user, sending an expression, receives an expression identifier in response and can check with the server at regular intervals "has the expression been calculated"?
If the expression is finally calculated - he will get the result.
Remember that some parts of an arithmetic expression can be evaluated **parallel**.

## Front-end part

## GUI, which can be visualized as 4 pages

1) Arithmetic expression input form. The user enters an arithmetic expression and sends a **POST** http request with that expression to the back-end. Note: Requests must be **idempotent**. A **unique identifier** is appended to requests. If a user sends a request with an identifier that has already been sent and accepted for processing - the response is 200. Possible response options:
    1. _200_ - The expression has been successfully received, parse and accepted for processing
    2. _400_ - The expression is invalid
    3. _500_ - Something is wrong on the back-end. As an answer we need to return the id of the expression accepted for execution.
2) A page with a list of expressions in the form of a list with expressions. Each entry on the page contains status, expression, date of its creation and date of computation completion. The page receives data by GET http-receipt from the back-end.
3) A page with a list of operations in the form of pairs: operation name + time of its execution (editable field). As already specified in the task condition, our operations are executed "as if very long". The page receives data by GET http-receipt from the back-end. The user can customize the operation execution time and save the changes.
4) A page with a list of computational capabilities. The page receives data by GET http-request from the server in the form of pairs: name of computational resource + operation performed on it.

## Requirements:

1) Orchestrator can be restarted without losing state. All expressions must be stored in the DBMS.
2) Orchestrator should keep track of tasks that are running too long (the compute resource can also leave the connection) and make them re-accessible for computation.


## Back-end part.

## Consists of 2 elements:

- A server that takes an arithmetic expression, translates it into a set of sequential tasks, and provides the order in which they are executed. Hereafter we will call it the orchestrator.
- A calculator that can receive a task from the orchestrator, execute it and return the result to the server. Hereinafter we will call it an agent.

### Orchestrator
A server that has the following endpoints:

- Adding an arithmetic expression calculation.
- Obtaining a list of expressions with statuses.
- Getting the value of an expression by its identifier.
- Getting the list of available operations with their execution time.
- Obtaining a task to be executed.
- Receiving the result of data processing.


### Agent
A daemon that receives the expression to be evaluated from the server, evaluates it and sends the result of the expression to the server. At startup, the daemon runs several goroutines, each of which acts as an independent evaluator. The number of goroutines is controlled by an environment variable.