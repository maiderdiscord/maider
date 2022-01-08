type CheckProxyResult = {
  success: boolean;
  result: { proxy: string; alive: boolean }[];
};

type CheckTokensResult = {
  success: boolean;
  result: { token: string; alive: boolean }[];
};

interface Window {
  core: {
    getHWID(): Promise<string>;
    checkProxy(proxyType: number, proxies: string[]): Promise<CheckProxyResult>;
    checkTokens(
      tokens: string[],
      proxies: string[],
      proxyType: number,
    ): Promise<CheckTokensResult>;
    joiner(
      tokens: string[],
      proxies: string[],
      proxyType: number,
      code: string,
    ): Promise<null>;
    leaver(
      tokens: string[],
      proxies: string[],
      proxyType: number,
      guildID: string,
    ): Promise<null>;
  };
}
