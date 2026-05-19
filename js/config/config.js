class Config {
  constructor() {
    this.appEnv = process.env.APP_ENV || 'development';
    this.port = process.env.PORT || '3000';
    this.apiKey = process.env.API_KEY || '';
  }
}

module.exports = new Config();
