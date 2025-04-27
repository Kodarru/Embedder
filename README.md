# 🗃️ Embedder

A simple service to create embeds from a URL and display it in a Discord message.

## Running the Service

> [!NOTE]
> In order to see the embeds, you must host it on a server that is accessible from the internet. You can use services like [Cloudflare Pages](https://pages.cloudflare.com/) or [Railway](https://railway.app/) to host your service for free.

1. Clone the repository.
2. Install the required dependencies:
   ```bash
   go get github.com/gin-gonic/gin
   go get github.com/aymerick/raymond
   ```
3. Build the service:
   ```bash
    go build Main.go
    ```
4. Run the service:
   ```bash
   ./Main
   ```
5. Open your browser and navigate to `http://localhost:8080/` to see the service in action.

## Parameters

- `title`: The title of the embed.
- `description`: The description of the embed.
- `url`: The URL to be embedded.
- `color`: The color of the embed. (default: `#0099ff`)
- `large_image`: The large image to be displayed in the embed. (true or false, default: false)
- `image`: The image to be displayed in the embed.

Example URL:
```
http://localhost:8080/?title=Hello%20World&description=This%20is%20a%20test%20embed&url=https://example.com&color=%230099ff&image=https://example.com/image.png
```