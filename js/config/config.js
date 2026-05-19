const fs = require('fs');
const path = require('path');

class Config {
  static schema = {
    appEnv: { label: 'App Environment', default: 'development' },
    port: { label: 'Port', default: '3000', validator: Config.validatePort },
    apiKey: { label: 'API Key', default: '', validator: Config.validateNotEmpty }
  };

  static validatePort(val) {
    const port = parseInt(val, 10);
    if (isNaN(port) || port < 1 || port > 65535) return 'Port must be between 1 and 65535';
    return null;
  }

  static validateNotEmpty(val) {
    if (!val || val.trim() === '') return 'Value cannot be empty';
    return null;
  }

  constructor(configPath = null) {
    // Initialize from schema defaults
    Object.keys(Config.schema).forEach(key => {
      this[key] = Config.schema[key].default;
    });

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
        Object.keys(Config.schema).forEach(key => {
          if (data[key]) this[key] = data[key];
        });
      } catch (err) {
        console.warn('Warning: Failed to parse config.json', err.message);
      }
    }
  }

  loadFromEnv() {
    const prefix = 'BP_';
    const getEnv = (key) => process.env[`${prefix}${key.toUpperCase().replace(/-/g, '_')}`];

    Object.keys(Config.schema).forEach(key => {
      // Map camelCase to kebab-case for env check (e.g. appEnv -> app-env)
      const flagName = key.replace(/([A-Z])/g, "-$1").toLowerCase();
      const envVal = getEnv(flagName);
      if (envVal) this[key] = envVal;
    });
  }

  save(configPath = null) {
    const finalPath = configPath || path.join(process.cwd(), 'config.json');
    const data = {};
    Object.keys(Config.schema).forEach(key => {
      data[key] = this[key];
    });
    fs.writeFileSync(finalPath, JSON.stringify(data, null, 2));
  }
}

module.exports = Config;
