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
    checkProxy(proxyType: number, proxies: string[]): Promise<CheckProxyResult>;
    checkTokens(
      tokens: string[],
      proxies: string[],
      proxyType: number,
    ): Promise<CheckTokensResult>;
  };
}
