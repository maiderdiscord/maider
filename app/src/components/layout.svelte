<script lang="ts">
  import Header from '@/components/header.svelte';
  import { Container } from 'sveltestrap';
</script>

<Header />
<Container>
  <slot />
</Container>
