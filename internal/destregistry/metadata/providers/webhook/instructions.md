# Webhook Configuration Instructions

To receive events from the webhook destination, you need to set up a webhook endpoint.

A webhook endpoint is a URL that you provide to an HTTP server. When an event is sent to the webhook destination, an HTTP POST request is made to the webhook endpoint with a JSON payload. Information such as the event type will be sent in the HTTP headers.
