# JavaScript API Client Boilerplate

A modern, asynchronous API client for Node.js.

## Features
- **Promise-based:** Uses `async/await` and `Promise.all` for concurrency.
- **Native Fetch:** Built using the native `fetch` API (Node 18+).
- **Timeouts:** Built-in timeout support via `AbortSignal`.

## Usage

```javascript
const client = new APIClient('https://api.example.com');

const results = await client.doConcurrent([
  { method: 'GET', path: '/resource' }
]);
```

## Running Tests

```bash
cd js
node --test
```
