# Contract testing for Golang GRPC with Pact

why do we need contract-testing ?

The main drawbacks of end-to-end testing is that it takes time to execute and sometimes the environment can
be flank so we need to repeat the tests.The main intention of end-to-end test is very the communication between the services.
we can do this with the help of contract testing. 

what is contract testing ?

consumer - who calls
provider - who responds to the requests

We write unit tests on the consumer side with all the test case and it generates the file
And on the provider side, we play the pact files against the API to check whether consumer and provider are in sync. (we can mock
the integration with other services on the provider side)

Pact plugin is used as demo for the contract testing in the example.

pact broker - https://docs.pact.io/pact_broker


TODO