type CheckProxyResult = {
  success: boolean;
  result: { proxy: string; alive: boolean }[];
};

interface Window {
  core: {
    checkProxy(proxyType: number, proxies: string[]): Promise<CheckProxyResult>;
  };
}
