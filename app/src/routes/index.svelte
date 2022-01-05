<script lang="ts">
  import Layout from '@/components/layout.svelte';
  import Proxy from '@/components/proxy.svelte';
  import { Col, Row } from 'sveltestrap';
</script>

<Layout>
  <Row cols="3">
    <Col>
      <h2 class="text-center">Proxy</h2>
      <Proxy />
    </Col>
  </Row>
</Layout>
