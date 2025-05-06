<!-- Start SDK Example Usage [usage] -->
```python
# Synchronous Example
from openapi import SDK


with SDK(
    admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
) as sdk:

    res = sdk.health.check()

    # Handle response
    print(res)
```

</br>

The same SDK client can also be used to make asychronous requests by importing asyncio.
```python
# Asynchronous Example
import asyncio
from openapi import SDK

async def main():

    async with SDK(
        admin_api_key="<YOUR_BEARER_TOKEN_HERE>",
    ) as sdk:

        res = await sdk.health.check_async()

        # Handle response
        print(res)

asyncio.run(main())
```
<!-- End SDK Example Usage [usage] -->