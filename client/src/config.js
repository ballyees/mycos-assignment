class ConfigSingleton {
  static instance = null
  constructor() {
    if (ConfigSingleton.instance) {
      return ConfigSingleton.instance
    } else {
      ConfigSingleton.instance = this
      this.config = this.configInitial()
      return this
    }
  }
  static getInstance() {
    if(ConfigSingleton.instance == null) {
      new ConfigSingleton() // new instance when the instance is null
    }
    return ConfigSingleton.instance
  }
  static getConfig() {
    return ConfigSingleton.getInstance().config
  }
  configInitial() {
    let config = {}
    if (process.env.NODE_ENV === 'production') {
      config['baseURL'] = process.env.REACT_APP_URL_PROD
    } else {
      config['baseURL'] = process.env.REACT_APP_URL_DEV
    }
    config['pvdURL'] = [config['baseURL'], "v1/pvd"].join("/")
    return config
  }
}
export default ConfigSingleton;