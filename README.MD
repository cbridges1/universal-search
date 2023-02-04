## Universal Search
The purpose of this application is to provide a unified search utility to query various search engines and social platforms.
#### To run enter:
```
go run cmd/app/app.go
or
air (hot reload)
```

#### Query parameters:

- query - Query text for search
- limit - Number of results to return from each platform (the default is 10)

### Supported platforms:

- Google search engine
- Reddit
### Docker configuration


#### To build the image:
```
docker build . --file Dockerfile --tag universal-search
```

#### To run container:
```
docker run -p 8080:8080 universal-search
```

#### To run container with .env file:
```
docker run -p 8080:8080 --env-file .env universal-search
```

### Environment Variables

- `GOOGLE_KEY` API key for Google Programmable Search Engine
- `CX` Search engine ID for Google Programmable Search Engine