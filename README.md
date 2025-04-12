# Go RSS Aggregator
A pet project created to learn Go.

It is a heavily modified and somewhat improved version of the RSS aggregator server created in the tutorial linked below, by [@bootdotdev](https://www.youtube.com/@bootdotdev):

> [!NOTE]
> **Tutorial**: https://www.youtube.com/watch?v=un6ZyFkqFKo

The API allows you to register RSS feeds and subscribe to them, which then allows you to get a list of all the latest posts from your subscribed feeds.

The server periodically fetches the registered feeds to update the available information about their posts.

## Instructions

I have no idea why you would want to start it, but if for some reason you do, here are the instructions after cloning the repository:

1. You will need to create a database. While I've used PostgreSQL, it should work without issues in other databases. Make sure to set the correct driver and connection URL in the next steps.

2. Run the init command, which will create the config file, download dependencies, and compile the project:
    ```sh
    make init
    ```

3. Edit the `.env` file and modify the values. You will mainly need to modify the values of these two config variables:
    ```sh
    DB_DRIVER=postgres
    DB_URL=${DB_DRIVER}://user:pass@localhost:5432/dbname
    ```

4. You should now be able to start the server using the generated binary:
    ```sh
    ./bin/rssagg-server
    ```

5. If you left the default port in the config file, you should now be able to connect to the server on `localhost:8080`. For example, you can call the health status endpoint *(using `curl` or whatever HTTP client you normally use)*:
    ```
    [GET] http://localhost:8080/v1/healthz
    ```

    You should get a response like this:
    ```json
    {
        "status": "ok"
    }
    ```

6. Voilà! You are now the proud owner and host of a super cool RSS aggregator server.

## API
In the even stranger case that you want to actually *use* the server, here's all the information you'll need: [API documentation](./api/api.md).
