const ffi = require('ffi-napi');
const app = require('electron');

module.exports = class GoLib {
  constructor() {
    const libPath = app.isPackaged ? './resources/golib' : './golib';
    this.lib = ffi.Library(libPath, {
      CheckProxy: ['string', ['string']],
    });
  }

  checkProxy(proxyType, proxies) {
    const input = {
      type: proxyType,
      proxies,
    };

    return new Promise((resolve, reject) => {
      this.lib.CheckProxy.async(JSON.stringify(input), (err, res) => {
        if (err != null) {
          reject(err);
          return;
        }
        resolve(JSON.parse(res));
      });
    });
  }
};
