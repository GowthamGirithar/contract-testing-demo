

# Loss of visibility into real-world client usage
With contract testing, the provider knows what all the fields are that the consumer requires, and this information is very helpful if we want to deprecate any field.
But if an application uses a field mask to request only required fields in the requests to return only those fields, we won't benefit from contract testing from this point of view.

# Protocol-level breaking changes

1. Data Format Changes 
 - Adding the new field should not cause any issues because consumer will process this data only if its needed for them. This should be covered as part of the design document.
 - Removing the field should not cause any issues as it will be marked as deprecated 
 - Updating the field is not recommended and so we create new field and mark the one needs to be updated as deprecated.

2. Serialization Changes
 - GRPC main advantages is serialization technique and so using json serialization is not recommended. Also we prefer to use the request of type message (domain object structure) instead json bytes. If we adhere to best practises we wont need contract testing for this too. (we prefer GRPC for backend communication and not for UI to backend)

3. Transport Protocol Changes:
 - GRPC by default uses HTTP/2 for better performance and no need to do changes for now. (TODO: research)
 
4. Security Protocol Changes:
 - If provider enforce security to use HTTPS , it should be communicated to the consumers and test it instead via contract testing

5. Compression Changes:
 - Contract testing is helpful to detect whether client and server enabled the same compression.
 - If client sends the compression technique which is not supported by server, then server return UNIMPLEMENTED
 - If server sends the compression technique which is not supported by client, then error code is INTERNAL
 
6. Performance Optimization Changes:
- TODO

# Optional and Required fields 

'required' option is deprecated in the latest version of protobuf and contract testing might help providers to validate the required fields with data format.
same applicable to "one of" also
