class APIClient {
  constructor(baseURL, timeout = 10000) {
    this.baseURL = baseURL;
    this.timeout = timeout;
  }

  async do(params) {
    const { method, path, body, headers = {} } = params;
    const url = `${this.baseURL}${path}`;
    
    const options = {
      method,
      headers: { ...headers },
      signal: AbortSignal.timeout(this.timeout)
    };

    if (body) {
      options.body = JSON.stringify(body);
      options.headers['Content-Type'] = 'application/json';
    }

    try {
      const response = await fetch(url, options);
      const data = await response.arrayBuffer();
      
      return {
        method,
        statusCode: response.status,
        data: Buffer.from(data),
        error: null
      };
    } catch (error) {
      return {
        method,
        statusCode: 0,
        data: null,
        error: error.message
      };
    }
  }

  async doConcurrent(requests) {
    return Promise.all(requests.map(req => this.do(req)));
  }
}

async function main() {
  const client = new APIClient('https://httpbin.org');

  const requests = [
    { method: 'GET', path: '/get' },
    { method: 'POST', path: '/post', body: { msg: 'hello' } },
    { method: 'PUT', path: '/put', body: { update: 'true' } },
    { method: 'DELETE', path: '/delete' },
    { method: 'PATCH', path: '/patch', body: { patch: 'true' } },
    { method: 'OPTIONS', path: '/options' },
    { method: 'HEAD', path: '/get' }
  ];

  console.log('Executing requests...');
  const results = await client.doConcurrent(requests);

  results.forEach(res => {
    const status = res.statusCode.toString().padStart(3);
    const method = res.method.padEnd(7);
    console.log(`Method: ${method} | Status: ${status} | Error: ${res.error}`);
  });
}

if (require.main === module) {
  main().catch(console.error);
}

module.exports = APIClient;
