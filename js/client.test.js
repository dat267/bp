const { test, describe } = require('node:test');
const assert = require('node:assert');
const http = require('node:http');
const APIClient = require('./client');

function createTestServer() {
  return http.createServer((req, res) => {
    res.writeHead(200, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify({ status: 'ok' }));
  });
}

describe('APIClient', () => {
  test('should perform a successful request', async () => {
    const server = createTestServer();
    await new Promise(resolve => server.listen(0, resolve));
    const port = server.address().port;
    const client = new APIClient(`http://localhost:${port}`);

    const res = await client.do({ method: 'GET', path: '/test' });
    
    assert.strictEqual(res.statusCode, 200);
    assert.strictEqual(res.error, null);
    
    server.close();
  });

  test('should perform concurrent requests', async () => {
    const server = createTestServer();
    await new Promise(resolve => server.listen(0, resolve));
    const port = server.address().port;
    const client = new APIClient(`http://localhost:${port}`);

    const requests = [
      { method: 'GET', path: '/1' },
      { method: 'GET', path: '/2' }
    ];

    const results = await client.doConcurrent(requests);
    
    assert.strictEqual(results.length, 2);
    results.forEach(res => {
      assert.strictEqual(res.statusCode, 200);
    });
    
    server.close();
  });
});
