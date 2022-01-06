<script lang="ts">
  import { proxies, proxyType } from '@/store';
  import { onDestroy } from 'svelte';

  import { Alert, Button, Form, FormGroup, Input, Label } from 'sveltestrap';

  let files: FileList;
  let chosenProxyType = 0;
  let proxiesCount = 0;
  let loadedProxiesCount = 0;
  let loading = false;
  let loaded = false;
  let eta = 0;
  let etaMinutes = 0;

  const unsubscribe = proxies.subscribe((v) => {
    proxiesCount = v.length;
  });
  onDestroy(unsubscribe);

  const onClick = async () => {
    proxies.update(() => []);

    if (typeof files === 'undefined' || files.length !== 1) {
      return;
    }

    loading = true;
    const file = files.item(0);
    const text = await file.text();
    const loadedProxies = text
      .split('\n')
      .map((proxy) => proxy.replace('\r', ''));
    loadedProxiesCount = loadedProxies.length;
    eta = (loadedProxiesCount / 50) * 5 + 10;
    etaMinutes = Math.floor(eta / 60);
    const res = await window.core.checkProxy(chosenProxyType, loadedProxies);
    proxyType.update(() => chosenProxyType);
    proxies.update(() => res.result.filter((x) => x.alive).map((x) => x.proxy));
    loading = false;
    loaded = true;
  };
</script>

<Form>
  <FormGroup>
    <Label for="file">Proxy File</Label>
    <Input id="file" type="file" bind:files />
  </FormGroup>

  <FormGroup>
    <Label>Proxy Type</Label>
    <Input type="radio" bind:group={chosenProxyType} value={0} label="HTTP" />
    <Input type="radio" bind:group={chosenProxyType} value={1} label="SOCKS5" />
  </FormGroup>

  <div class="text-center">
    <Button
      type="button"
      color="primary"
      on:click={onClick}
      bind:disabled={loading}>Load Proxy</Button
    >

    {#if loading}
      <p>
        ETA: {etaMinutes}m{Math.floor(eta - etaMinutes * 60)}s
      </p>
    {/if}
  </div>

  {#if !loading && loaded}
    <Alert class="mt-3" color="primary"
      >{proxiesCount} / {loadedProxiesCount} proxies alive</Alert
    >
  {/if}
</Form>
