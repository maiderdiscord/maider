<script lang="ts">
  import { proxies, tokens } from '@/store';
  import { onDestroy } from 'svelte';
  import {
    Nav,
    Collapse,
    Image,
    Navbar,
    NavbarBrand,
    NavItem,
    NavLink,
  } from 'sveltestrap';

  let proxiesCount = 0;
  let tokensCount = 0;

  const unsubscribeProxies = proxies.subscribe((v) => {
    proxiesCount = v.length;
  });

  const unsubscribeTokens = tokens.subscribe((v) => {
    tokensCount = v.length;
  });

  onDestroy(() => {
    unsubscribeProxies();
    unsubscribeTokens();
  });
</script>

<Navbar color="dark" dark expand="md">
  <NavbarBrand>
    <Image
      class="rounded-circle"
      src="https://avatars.githubusercontent.com/u/97165057"
      width="64"
    />
    Maider
  </NavbarBrand>
  <Collapse isOpen={true} navbar expand="md">
    <Nav navbar>
      <NavItem>
        <div class="mx-3">Proxies: {proxiesCount}</div>
      </NavItem>
      <NavItem>
        <div class="mx-3">Tokens: {tokensCount}</div>
      </NavItem>
    </Nav>
  </Collapse>
</Navbar>
