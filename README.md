# sfmlx-test

A proof of concept service meant to assist in injecting ads inside an existing `sfml` file.

# Usage

To run the project simply 
`go run main.go`

To read a local SFML file simply edit `main.go` and change the file you want to read

# TODOs

- [x] Traverse XML and count number of possible ad slots
- [ ] Turn the project into a deployable HTTP service
- [ ] Introduce an endpoint that accepts flyerId, flyerRunId, merchantId, etc... and retrieve our production sfmls, and returns to the caller the modified SFML with ads
- [ ] Introduce a config that outlines business rules on when and where ads are injected.
