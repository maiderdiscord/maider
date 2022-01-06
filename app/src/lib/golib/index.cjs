const ffi = require('ffi-napi');
const { is } = require('electron');

module.exports = class GoLib {
  constructor() {
    this.lib = ffi.Library('./golib', {
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
