# Testing the `killCatcher`

In order to test the `killCatcher`, I designed two test cases

- the `killCatcher` will identify the `SIGTERM` but will NOT stop execution of the main logic
- the `killCatcher` will identify the `SIGTERM` but will stop execution of the main logic

Both of these test cases are tested in `minikube` cluster.

## NOT stopping execution of the main logic

If you wish to run this test case, please run

`make testContinue`

This command will run the test and will save the logs of the pod in `test-app.log` file.

## Stopping execution of the main logic

If you wish to run this test case, please run

`make testEnd`

This command will run the test and will save the logs of the pod in `test-app.log` file.