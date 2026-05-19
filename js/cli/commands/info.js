class InfoCommand {
  constructor(config) {
    this.name = 'info';
    this.description = 'Displays environment information';
    this.config = config;
  }

  execute() {
    console.log(`Environment: ${this.config.appEnv}`);
    console.log(`Port:        ${this.config.port}`);
    console.log(`API Key:     ${this.config.apiKey ? '********' : 'not set'}`);
  }
}

module.exports = InfoCommand;
