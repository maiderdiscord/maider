<script lang="ts">
  import { proxies } from '@/store';
  import { onDestroy } from 'svelte';
  import {
    Nav,
    Collapse,
    Image,
    Navbar,
    NavbarBrand,
    NavItem,
  } from 'sveltestrap';

  let proxiesCount = 0;

  const unsubscribe = proxies.subscribe((v) => {
    proxiesCount = v.length;
  });
  onDestroy(unsubscribe);
</script>

<Navbar color="light" light expand="md">
  <NavbarBrand>
    <Image src="https://avatars.githubusercontent.com/u/97165057" width="64" />
    Maider
  </NavbarBrand>
  <Collapse isOpen={true} navbar expand="md">
    <Nav navbar>
      <NavItem>Proxies: {proxiesCount}</NavItem>
    </Nav>
  </Collapse>
</Navbar>
