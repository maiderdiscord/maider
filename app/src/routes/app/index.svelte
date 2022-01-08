<script lang="ts">
  import Joiner from '@/components/joiner.svelte';
  import Layout from '@/components/layout.svelte';
  import Leaver from '@/components/leaver.svelte';
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
    <Col>
      <h2 class="text-center">Joiner</h2>
      <Joiner />
    </Col>
  </Row>
  <Row cols="3">
    <Col>
      <h2 class="text-center">Leaver</h2>
      <Leaver />
    </Col>
  </Row>
</Layout>
