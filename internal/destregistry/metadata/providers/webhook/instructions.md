# Webhook Configuration Instructions

To receive events from the webhook destination, you need to set up a webhook endpoint. A webhook endpoint is a URL that you provide to a HTTP server. When an event is sent to the webhook destination, the HTTP POST request is sent to the webhook endpoint.

The webhook destination will send the event to the webhook endpoint as a JSON object. The JSON object will contain the event data and the event type.
