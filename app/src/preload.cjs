// eslint-disable-next-line @typescript-eslint/no-var-requires
const GoLib = require('./lib/golib/index.cjs');
const { contextBridge } = require('electron');
const { machineId } = require('node-machine-id');

const goLib = new GoLib();

contextBridge.exposeInMainWorld('core', {
  getHWID: async () => {
    return machineId(true);
  },
  checkProxy: (proxyType, proxies) => {
    return goLib.checkProxy(proxyType, proxies);
  },
  checkTokens: (tokens, proxies, proxyType) => {
    return goLib.checkTokens(tokens, proxies, proxyType);
  },
  joiner: (tokens, proxies, proxyType, code) => {
    return goLib.joiner(tokens, proxies, proxyType, code);
  },
});
