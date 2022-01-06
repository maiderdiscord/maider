<script lang="ts">
  import { proxies, proxyType, tokens } from '@/store';
  import { onDestroy } from 'svelte';

  import { Alert, Button, Form, FormGroup, Input, Label } from 'sveltestrap';

  let files: FileList;
  let tokensCount = 0;
  let loadedTokensCount = 0;
  let loading = false;
  let loaded = false;
  let eta = 0;
  let etaMinutes = 0;
  let loadedProxies = [];
  let loadedProxyType = 0;

  const unsubscribeTokens = tokens.subscribe((v) => {
    tokensCount = v.length;
  });

  const unsubscribeProxies = proxies.subscribe((v) => {
    loadedProxies = v;
  });

  const unsubscribeProxyType = proxyType.subscribe((v) => {
    loadedProxyType = v;
  });

  onDestroy(() => {
    unsubscribeTokens();
    unsubscribeProxies();
    unsubscribeProxyType();
  });

  const onClick = async () => {
    tokens.update(() => []);

    if (typeof files === 'undefined' || files.length !== 1) {
      return;
    }

    loading = true;
    const file = files.item(0);
    const text = await file.text();
    const loadedTokens = text
      .split('\n')
      .map((token) => token.replace('\r', ''))
      .filter((token) => token.length > 0);
    loadedTokensCount = loadedTokens.length;
    eta = (loadedTokensCount / 50) * 5 + 10;
    etaMinutes = Math.floor(eta / 60);
    const res = await window.core.checkTokens(
      loadedTokens,
      loadedProxies,
      loadedProxyType,
    );
    tokens.update(() => res.result.filter((x) => x.alive).map((x) => x.token));
    loading = false;
    loaded = true;
  };
</script>

<Form>
  <FormGroup>
    <Label for="file">Token File</Label>
    <Input id="file" type="file" bind:files />
  </FormGroup>

  <div class="text-center">
    <Button
      type="button"
      color="primary"
      on:click={onClick}
      bind:disabled={loading}>Load Token</Button
    >

    {#if loading}
      <p>
        ETA: {etaMinutes}m{Math.floor(eta - etaMinutes * 60)}s
      </p>
    {/if}
  </div>

  {#if !loading && loaded}
    <Alert class="mt-3" color="primary"
      >{tokensCount} / {loadedTokensCount} tokens alive</Alert
    >
  {/if}
</Form>
