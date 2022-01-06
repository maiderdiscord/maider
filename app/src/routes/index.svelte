<script lang="ts">
  import Layout from '@/components/layout.svelte';
  import Proxy from '@/components/proxy.svelte';
  import Token from '@/components/token.svelte';
  import { Col, Row } from 'sveltestrap';
</script>

<Layout>
  <Row cols="3">
    <Col>
      <h2 class="text-center">Proxy</h2>
      <Proxy />
    </Col>
    <Col>
      <h2 class="text-center">Token</h2>
      <Token />
    </Col>
  </Row>
</Layout>
