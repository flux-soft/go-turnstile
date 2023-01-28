# Go TURNSTILE

## Description
API Implementation of [Cloudflare turnstile](https://www.cloudflare.com/products/turnstile/) for Go

## Installation

```bash
go get github.com/flux-soft/go-turnstile
```

## Usage
Here's an example of usage:

```go
package main

import (
	"log"
	"github.com/flux-soft/go-turnstile"
)

func main() {
    validator := turnstile.New("SECRET_KEY", turnstile.DEFAULT_TIMEOUT)
    // Captcha token receive from header cf-turnstile-response
    // Remote IP isn't required, if you wouldn't pass IP then insert ""
	response, err := validator.Verify("CAPTCHA_TOKEN", "REMOTE_IP")
	if err != nil {
		log.Fatal(err)
		return
	}

	if !response.Success {
		// Failed verify captcha
	}
    // or 
    if !response.IsSuccess()) {
        //
    }

    if response.HasErrors() {
        for k, v := range response.ErrorCodes {
            println(v)
        }
    }
}
```

```html
<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta http-equiv="X-UA-Compatible" content="IE=edge">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
    </head>
    <body>
        <form action="" method="POST">
            <div class="cf-turnstile" data-sitekey="YOUR_PUBLIC_SITE_KEY"></div> 
            <button type="submit">check</button>
        </form>
        <script src="https://challenges.cloudflare.com/turnstile/v0/api.js" async defer></script>
    </body>
</html>
```

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.