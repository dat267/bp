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
    const isDefault = !configPath;
    const finalPath = configPath || path.join(process.cwd(), 'config.json');
    
    if (!fs.existsSync(finalPath) && isDefault) {
      const example = {
        appEnv: this.appEnv,
        port: this.port,
        apiKey: this.apiKey
      };
      try {
        fs.writeFileSync(finalPath, JSON.stringify(example, null, 2));
      } catch (err) {
        console.warn('Warning: Could not auto-generate config.json', err.message);
      }
    }

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

  save(configPath = null) {
    const finalPath = configPath || path.join(process.cwd(), 'config.json');
    const data = {
      appEnv: this.appEnv,
      port: this.port,
      apiKey: this.apiKey
    };
    fs.writeFileSync(finalPath, JSON.stringify(data, null, 2));
  }
}

module.exports = Config;
