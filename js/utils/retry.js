/**
 * Resilient Retry utility with Exponential Backoff and Jitter
 */
async function withRetry(operation, options = {}) {
  const {
    maxAttempts = 3,
    initialDelay = 1000,
    maxDelay = 30000,
    backoffFactor = 2,
    useJitter = true
  } = options;

  let lastError;
  let delay = initialDelay;

  for (let attempt = 1; attempt <= maxAttempts; attempt++) {
    try {
      return await operation();
    } catch (error) {
      lastError = error;
      if (attempt === maxAttempts) break;

      let sleepTime = delay;
      if (useJitter) {
        // Add random jitter: [0, delay / 2]
        const jitter = Math.random() * (delay / 2);
        sleepTime += jitter;
      }

      await new Promise(resolve => setTimeout(resolve, sleepTime));

      delay = Math.min(delay * backoffFactor, maxDelay);
    }
  }

  throw lastError;
}

module.exports = { withRetry };
