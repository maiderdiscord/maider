<script lang="ts">
  import { proxies } from '@/store';
  import { onDestroy } from 'svelte';

  import { Alert, Button, Form, FormGroup, Input, Label } from 'sveltestrap';

  let files: FileList;
  let proxyType = 0;
  let proxiesCount = 0;
  let loadedProxiesCount = 0;
  let loading = false;
  let loaded = false;

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
    const res = await window.core.checkProxy(proxyType, loadedProxies);
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
    <Input type="radio" bind:group={proxyType} value={0} label="HTTP" />
    <Input type="radio" bind:group={proxyType} value={1} label="SOCKS5" />
  </FormGroup>

  <div class="text-center">
    <Button
      type="button"
      color="primary"
      on:click={onClick}
      bind:disabled={loading}>Load Proxy</Button
    >
  </div>

  {#if !loading && loaded}
    <Alert class="mt-3" color="primary"
      >{proxiesCount} / {loadedProxiesCount} proxies alive</Alert
    >
  {/if}
</Form>
