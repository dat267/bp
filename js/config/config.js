const fs = require('fs');
const path = require('path');

class Config {
  constructor(configPath = null) {
    this.appEnv = 'development';
    this.port = '3000';
    this.apiKey = '';

    this.loadFromFile(configPath);
    this.loadFromEnv();
  }

  loadFromFile(configPath) {
    const finalPath = configPath || path.join(process.cwd(), 'config.json');
    if (fs.existsSync(finalPath)) {
      try {
        const data = JSON.parse(fs.readFileSync(finalPath, 'utf8'));
        if (data.appEnv) this.appEnv = data.appEnv;
        if (data.port) this.port = data.port;
        if (data.apiKey) this.apiKey = data.apiKey;
      } catch (err) {
        console.warn('Warning: Failed to parse config.json', err.message);
      }
    }
  }

  loadFromEnv() {
    if (process.env.APP_ENV) this.appEnv = process.env.APP_ENV;
    if (process.env.PORT) this.port = process.env.PORT;
    if (process.env.API_KEY) this.apiKey = process.env.API_KEY;
  }
}

module.exports = Config;
