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
      this.save(finalPath);
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
    // rclone style: auto-mapping BP_PORT, BP_APP_ENV, etc.
    const prefix = 'BP_';
    const getEnv = (key) => process.env[`${prefix}${key.toUpperCase().replace(/-/g, '_')}`];

    this.appEnv = getEnv('app-env') || this.appEnv;
    this.port = getEnv('port') || this.port;
    this.apiKey = getEnv('api-key') || this.apiKey;
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
