// eslint-disable-next-line @typescript-eslint/no-var-requires
const GoLib = require('./lib/golib/index.cjs');
const { contextBridge } = require('electron');

const goLib = new GoLib();

contextBridge.exposeInMainWorld('core', {
  checkProxy: (proxyType, proxies) => {
    return goLib.checkProxy(proxyType, proxies);
  },
  checkTokens: (tokens, proxies, proxyType) => {
    return goLib.checkTokens(tokens, proxies, proxyType);
  },
});
