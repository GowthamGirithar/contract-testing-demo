***<ins>Contract testing for Golang GRPC with Pact</ins>***
**Contract testing :**

Contract testing is to verify the contracts between services by mocking the downstream services. It uses two concepts, such as consumer and provider. The consumer is the one who calls the API and waits for the response, whereas the provider is the one who responds to the requests. There are different tools available on the market for contract testing, and some are pact and hypertest.

**Contract testing with Pact:**

****![](https://lh7-us.googleusercontent.com/GXQ8255uXXvGaTudzc9ebMkxK1St_txlsYEU2c6NAayXg4AyM0CbIs_MrO2L0rQ4h4r2aLEM6exChdHymNc07rsEzuGe9wDN8V6u9N03HsiuaV7b80O-Rmb_g_vD6D7gL7vtp6GOqToTVY5HQvU4lH8)****

reference : https\://docs.pact.io

It starts with defining the contract, on which all the parties agree. Once the contract has been agreed upon, the consumer writes the unit tests with a pact, which results in generating the pact file (json) as a result of the run.

To verify the contract, we need to run the pact file against the provider to know whether the contract is the same as what we agreed to or not.

There are two ways in which we can do this: consumer-driven tests, where consumers generate the pact file, which is played against the provider, and provider-driven tests, which are called provider-driver tests.



More details about verification of pacts using pact brokers and can-i-deploy can be found in the links provided below.

- https\://docs.pact.io/pact\_nirvana/step\_6

- https\://docs.pact.io/pact\_broker

- https\://docs.pact.io/pact\_broker/can\_i\_deploy

Next comes a question: do we need contract testing for the RPC frameworks, which have rich tooling and documentation? It really depends on individual use cases, so I thought to add some details for GRPC. (Examples can be seen in this repo.)

**Contract Testing with GRPC:**

- **Loss of visibility into real-world client usage**

With contract testing, the provider knows what all the fields are that the consumer requires, and this information is very helpful if we want to deprecate any field. But if an application uses a field mask to request only required fields in the requests to return only those fields, we won't benefit from contract testing from this point of view.

- **Protocol-level breaking changes**

1. Data Format Changes - Adding the new field should not cause any issues because consumers will process this data only if it is needed for them. This should be covered as part of the design document. Removing the field should not cause any issues, as it will be marked as deprecated.Updating the field is not recommended, so we create a new field and mark the one that needs to be updated as deprecated.

2) Serialisation Changes - GRPC's main advantage is its serialisation technique, so using JSON serialisation is not recommended. Also, we prefer to use the request type message (domain object structure) instead of JSON bytes. If we adhere to best practices, we won't need contract testing for this either. (We prefer GRPC for backend communication and not for UI to backend.)

3) Transport Protocol Changes - GRPC, by default, uses HTTP/2 for better performance; there is no need to make any changes for now. (TODO: research)

4) Security Protocol Changes -  If providers enforce security to use HTTPS, it should be communicated to consumers and tested instead via contract testing.

5) Compression Changes - Contract testing is helpful to detect whether the client and server have enabled the same compression.If the client sends a compression technique that is not supported by the server, then the server returns UNIMPLEMENTED. If the server sends a compression technique that is not supported by the client, the error code is INTERNAL.

6) Performance Optimization Changes - TODO (timeout behaviour)

- **Optional and Required fields**

'required' option is deprecated in the latest version of protobuf and contract testing might help providers to validate the required fields with data format. Same applies to "one of" also.

- **Ensuring narrow type-safety (strict encodings**

With matchers and expressions, we can narrow the allowed sets of values during validation.

References : https://pactflow.io/blog/contract-testing-for-grpc-and-protobufs/
