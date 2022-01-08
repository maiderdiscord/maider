<script lang="ts">
  import { proxies, proxyType, tokens } from '@/store';
  import { onDestroy } from 'svelte';

  import { Button, Form, FormGroup, Input } from 'sveltestrap';

  let loading = false;
  let loadedProxies = [];
  let loadedProxyType = 0;
  let loadedTokens = [];
  let code = '';

  const unsubscribeTokens = tokens.subscribe((v) => {
    loadedTokens = v;
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
    loading = true;
    await window.core.joiner(
      loadedTokens,
      loadedProxies,
      loadedProxyType,
      code,
    );
    loading = false;
  };
</script>

<Form>
  <FormGroup>
    <Input placeholder="xxxxxxxxxx" bind:value={code} />
  </FormGroup>

  <div class="text-center">
    <Button
      type="button"
      color="primary"
      on:click={onClick}
      bind:disabled={loading}>Join</Button
    >
  </div>
</Form>
