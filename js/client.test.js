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

  test('should handle timeouts', async () => {
    const server = http.createServer((req, res) => {
      setTimeout(() => {
        res.writeHead(200);
        res.end();
      }, 100);
    });
    await new Promise(resolve => server.listen(0, resolve));
    const port = server.address().port;
    
    // Short timeout of 20ms
    const client = new APIClient(`http://localhost:${port}`, 20);
    const res = await client.do({ method: 'GET', path: '/timeout' });
    
    assert.ok(res.error, 'Expected a timeout error');
    assert.strictEqual(res.statusCode, 0);
    
    server.close();
  });

  test('should send headers and body correctly', async () => {
    let receivedHeaders = {};
    let receivedBody = '';

    const server = http.createServer((req, res) => {
      receivedHeaders = req.headers;
      req.on('data', chunk => { receivedBody += chunk; });
      req.on('end', () => {
        res.writeHead(200);
        res.end();
      });
    });
    await new Promise(resolve => server.listen(0, resolve));
    const port = server.address().port;
    const client = new APIClient(`http://localhost:${port}`);

    await client.do({
      method: 'POST',
      path: '/post',
      headers: { 'x-test': 'value' },
      body: { foo: 'bar' }
    });

    assert.strictEqual(receivedHeaders['x-test'], 'value');
    assert.strictEqual(receivedHeaders['content-type'], 'application/json');
    assert.strictEqual(receivedBody, JSON.stringify({ foo: 'bar' }));

    server.close();
  });

  test('should handle empty requests list gracefully', async () => {
    const client = new APIClient('http://localhost:0');
    const results = await client.doConcurrent([]);
    assert.strictEqual(results.length, 0);
  });

  test('should respect throttleLimit concurrency constraint', async () => {
    let active = 0;
    let maxActive = 0;
    const server = http.createServer((req, res) => {
      active++;
      if (active > maxActive) {
        maxActive = active;
      }
      setTimeout(() => {
        active--;
        res.writeHead(200);
        res.end();
      }, 20);
    });
    await new Promise(resolve => server.listen(0, resolve));
    const port = server.address().port;
    const client = new APIClient(`http://localhost:${port}`);
    const requests = [
      { method: 'GET', path: '/1' },
      { method: 'GET', path: '/2' },
      { method: 'GET', path: '/3' },
      { method: 'GET', path: '/4' }
    ];
    await client.doConcurrent(requests, 2);
    assert.ok(maxActive <= 2);
    assert.ok(maxActive > 0);
    server.close();
  });
});
