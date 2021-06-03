# sfmlx-test

A proof of concept service meant to assist in injecting ads inside an existing `sfml` file.

# Usage

To run the service
`go run main.go`

Currently we only have one endpoint that returns a static loblaws SFML with ads injected
`curl localhost:8080/api/v1/sfmlx`

To change up the sfml, edit `sfmlParser.go` and point to another sfml file

# TODOs

-   [x] Traverse XML and count number of possible ad slots
-   [x] Turn the project into a deployable HTTP service
-   [x] Introduce an endpoint that returns sfml
-   [ ] Endpoint accepts flyerId, flyerRunId, merchantId, etc... and retrieve our production sfmls, and returns to the caller the modified SFML with ads
-   [ ] Introduce a config that outlines business rules on when and where ads are injected.
