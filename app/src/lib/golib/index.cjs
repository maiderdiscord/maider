const ffi = require('ffi-napi');
const { is } = require('electron');

module.exports = class GoLib {
  constructor() {
    this.lib = ffi.Library('./golib', {
      CheckProxy: ['string', ['string']],
      CheckTokens: ['string', ['string']],
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

  checkTokens(tokens, proxies, proxyType) {
    const input = {
      tokens,
      proxies,
      proxyType,
    };

    return new Promise((resolve, reject) => {
      this.lib.CheckTokens.async(JSON.stringify(input), (err, res) => {
        if (err != null) {
          reject(err);
          return;
        }
        resolve(JSON.parse(res));
      });
    });
  }
};
