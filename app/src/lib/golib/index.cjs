const ffi = require('ffi-napi');
const { is } = require('electron');

module.exports = class GoLib {
  constructor() {
    this.lib = ffi.Library('./golib', {
      CheckProxy: ['string', ['string']],
      CheckTokens: ['string', ['string']],
      Joiner: ['void', ['string']],
      Leaver: ['void', ['string']],
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

  joiner(tokens, proxies, proxyType, code) {
    const input = {
      tokens,
      proxies,
      proxyType,
      code,
    };

    return new Promise((resolve, reject) => {
      this.lib.Joiner.async(JSON.stringify(input), (err, _) => {
        if (err != null) {
          reject(err);
          return;
        }
        resolve(null);
      });
    });
  }

  leaver(tokens, proxies, proxyType, guildID) {
    const input = {
      tokens,
      proxies,
      proxyType,
      guildID,
    };

    return new Promise((resolve, reject) => {
      this.lib.Leaver.async(JSON.stringify(input), (err, _) => {
        if (err != null) {
          reject(err);
          return;
        }
        resolve(null);
      });
    });
  }
};
