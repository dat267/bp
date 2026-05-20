const { withRetry } = require('./utils/retry');

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

  async doWithRetry(params, retryOptions = {}) {
    return withRetry(async () => {
      const res = await this.do(params);
      if (res.error) throw new Error(res.error);
      if (res.statusCode >= 500 || res.statusCode === 429) {
        throw new Error(`Transient error: ${res.statusCode}`);
      }
      return res;
    }, retryOptions);
  }

  async doConcurrent(requests, throttleLimit = 5) {
    if (!requests || requests.length === 0) return [];
    const limit = Math.max(1, throttleLimit);
    const results = new Array(requests.length);
    let index = 0;
    const worker = async () => {
      while (index < requests.length) {
        const currentIdx = index++;
        results[currentIdx] = await this.do(requests[currentIdx]);
      }
    };
    const workers = [];
    for (let i = 0; i < Math.min(limit, requests.length); i++) {
      workers.push(worker());
    }
    await Promise.all(workers);
    return results;
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
  const results = await client.doConcurrent(requests, 3);

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
